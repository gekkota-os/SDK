package putil

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"pconf"
	"pformat"
	"strconv"
	"strings"
	"time"
)

// Private Functions
// --------------------------

// Create a ".gz" file from original
// and ".size" file that contains only original size info
// and delete original file
func createGzAndDeleteOriginal(original_path string) {
	fi_b, fi_err := ioutil.ReadFile(original_path)
	if nil != fi_err {
		log.Printf("createGzAndDeleteOriginal -> Error opening (%s) : %s\n", original_path, fi_err.Error()) // Print error
		return
	}

	// Create gz file
	gzip_path := original_path + ".gz"
	fgz, fgz_err := os.Create(gzip_path)
	if nil != fgz_err {
		log.Printf("createGzAndDeleteOriginal -> Error creating GZ (%s) : %s\n", gzip_path, fgz_err.Error()) // Print error
		return
	}
	// Write compressed data.
	wgz := gzip.NewWriter(fgz)
	wgz.Write(fi_b)
	wgz.Close()

	// Write original file size
	fi_size_int := GetFileSize(original_path)
	fi_size_path := original_path + ".size"
	fi_size, fi_size_err := os.Create(fi_size_path)
	if nil != fi_size_err {
		log.Printf("createGzAndDeleteOriginal -> Error creating Size (%s) : %s\n", fi_size_path, fi_size_err.Error()) // Print error
		return
	}
	fi_size.Write([]byte(strconv.FormatInt(fi_size_int, 10)))
	fi_size.Close()

	// Delete original file
	del_err := os.Remove(original_path)
	if nil != del_err {
		log.Printf("createGzAndDeleteOriginal -> Error Deleting (%s) : %s\n", original_path, del_err.Error()) // Print error
	}
}

// Public Functions
// --------------------------

// Gzip or not file_path according to config
func GzipOldFile(f_path string, conf *pconf.Config, check_file_date bool) {
	if conf.Gzip_older_than <= 0 {
		return
	}

	// Get file info
	_, f_name := filepath.Split(f_path)
	ext_blacklisted, _ := SliceStringContains(conf.Gzip_ext_blacklist, filepath.Ext(f_name))
	if ext_blacklisted {
		// PrintAndLog(fmt.Sprintf("-> gzipOldFile no need to gzip %s !\n", f_path), conf, false)// Verbose only
		return
	}

	if check_file_date {
		file_date := GetDateFromComptipixFileName(strings.TrimSuffix(f_name, filepath.Ext(f_name)))
		now := GetNowMidnight()
		after_date := now.Add(time.Duration(-1*conf.Gzip_older_than*24) * time.Hour)
		if file_date.Before(after_date) {
			// PrintAndLog(fmt.Sprintf("-> gzipOldFile GZIPING is after %s -%ddays = %s -> %s \n", now.Format("2006-01-02 15:04:05"), conf.Gzip_older_than, after_date.Format("2006-01-02 15:04:05"), f_path), conf, false)// Verbose only
			createGzAndDeleteOriginal(f_path)
		}
	} else {
		// PrintAndLog(fmt.Sprintf("-> gzipOldFile GZIPING ALWAYS = %s !\n", f_path), conf, false)// Verbose only
		createGzAndDeleteOriginal(f_path)
	}
}

// Get a date from a string date Comptipix format (like "20171122") from a date type
func GetDateFromComptipixFileName(date_str string) time.Time {
	now := GetNowMidnight() // use now by default, or if user date_str is invalid

	if date_str != "" {
		time_format_start := "20060102"
		d_date, d_date_err := time.Parse(time_format_start, date_str[0:8])
		if nil == d_date_err {
			now = d_date
		} else {
			log.Printf("Invalid date_str (%s)\n", date_str) // Print error day
		}
	}

	return now
}

// Get a date from a string inside Comptipix format (like "11/01/2018,01:42:16")
func GetDateFromComptipixFileContent(date_str string) (d time.Time, d_type pformat.DateTimeFormat, d_err error) {
	// Select ISO format or OLD format
	time_format := "2006-01-02,15:04:05" // ISO date by default like "2018-01-11,01:42:16"
	d_type = pformat.DateTimeFormat{
		Date_format_type: pformat.DATE_YYYYMMDD,
		Date_format_sep:  "-",
		Date_time_sep:    ",",
		Date_time_order:  pformat.DT_ORDER_DATE_TIME,
		Time_format_type: pformat.TIME_HHMMSS,
		Time_format_sep:  ":",
	}
	if 2 == strings.Count(date_str, "/") { // OLD format like "11/01/2018,01:42:16"
		time_format = "02/01/2006,15:04:05"
		d_type.Date_format_type = pformat.DATE_DDMMYYYY
		d_type.Date_format_sep = "/"
	}

	// Parse date
	d, d_err = time.Parse(time_format, date_str)
	if nil == d_err {
		return d, d_type, nil // Date read OK
	}

	log.Printf("GetDateFromComptipixFileContent -> Error read date %s -> %s\n", date_str, d_err.Error())
	return GetNowMidnight(), d_type, d_err // Date read error
}
