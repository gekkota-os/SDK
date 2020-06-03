package putil

import (
	"fmt"
	"pformat"
	"pconf"
	"strings"
	"os"
	"time"
)

// Declaration
// --------------------------

// All infos we get from a comptipix file name
type FileNameInfo struct {
	Yyyy								string		`json:"yyyy"`							// for file "20171107_presence.csv" it's "2017"
	Yy									string 		`json:"yy"`								// for file "20171107_presence.csv" it's "17"
	Mm									string		`json:"mm"`								// for file "20171107_presence.csv" it's "11"
	Dd 								string 		`json:"dd"`								// for file "20171107_presence.csv" it's "07"
	Time								*time.Time	`json:"-"`								// The go time defined Yyyy Mm Dd
	Suffix 							string		`json:"suffix"`						// for file "20171107_presence.csv" it's "_presence"
	Prefix 							string		`json:"prefix"`						// for file "20171107_presence.csv" it's "", and for file "log_20171107.csv" it's "log_"
	Extension 						string 		`json:"extension"`					// for file "20171107_presence.csv" it's ".csv"
}

// Private Functions
// --------------------------

// Create a ".gz" file from original
// and ".size" file that contains only original size info
// and delete original file
// func createGzAndDeleteOriginal(original_path string) {
// 	fi_b, fi_err := ioutil.ReadFile(original_path)
// 	if nil != fi_err {
// 		log.Printf("createGzAndDeleteOriginal -> Error opening (%s) : %s\n", original_path, fi_err.Error()) // Print error
// 		return
// 	}

// 	// Create gz file
// 	gzip_path := original_path + ".gz"
// 	fgz, fgz_err := os.Create(gzip_path)
// 	if nil != fgz_err {
// 		log.Printf("createGzAndDeleteOriginal -> Error creating GZ (%s) : %s\n", gzip_path, fgz_err.Error()) // Print error
// 		return
// 	}
// 	// Write compressed data.
// 	wgz := gzip.NewWriter(fgz)
// 	wgz.Write(fi_b)
// 	wgz.Close()

// 	// Write original file size
// 	fi_size_int := GetFileSize(original_path)
// 	fi_size_path := original_path + ".size"
// 	fi_size, fi_size_err := os.Create(fi_size_path)
// 	if nil != fi_size_err {
// 		log.Printf("createGzAndDeleteOriginal -> Error creating Size (%s) : %s\n", fi_size_path, fi_size_err.Error()) // Print error
// 		return
// 	}
// 	fi_size.Write([]byte(strconv.FormatInt(fi_size_int, 10)))
// 	fi_size.Close()

// 	// Delete original file
// 	del_err := os.Remove(original_path)
// 	if nil != del_err {
// 		log.Printf("createGzAndDeleteOriginal -> Error Deleting (%s) : %s\n", original_path, del_err.Error()) // Print error
// 	}
// }

// Public Functions
// --------------------------

// Get a FileNameInfo struct from a fileName
// return FileNameInfo struct, and error, nil if no error
func GetFileNameInfo(file_name string) (FileNameInfo, error) {
	var name_info FileNameInfo

	// Check file_name size
	if len(file_name) < 12 {
		return name_info, fmt.Errorf("GetFileNameInfo(%s) : File name too short", file_name)
	}

	// Get file yyyy, yy, mm, dd and prefix
	suffix_pos := strings.LastIndex(file_name, ".")
	if strings.HasPrefix(file_name, "log_") {
		name_info.Yyyy = file_name[4:8]
		name_info.Yy = file_name[6:8]
		name_info.Mm = file_name[8:10]
		name_info.Dd = file_name[10:12]
		name_info.Prefix = "log_"
		name_info.Suffix = file_name[12:suffix_pos]
	} else {
		name_info.Yyyy = file_name[0:4]
		name_info.Yy = file_name[2:4]
		name_info.Mm = file_name[4:6]
		name_info.Dd = file_name[6:8]
		name_info.Prefix = ""
		name_info.Suffix = file_name[8:suffix_pos]
	}
	name_info.Extension = file_name[suffix_pos:len(file_name)]

	// Set time
	time_format := "2006-01-02,15:04:05"
	d, d_err := time.Parse(time_format, name_info.Yyyy + "-" + name_info.Mm + "-" + name_info.Dd + ",00:00:00")
	if nil != d_err {
		return name_info, d_err		// Date read error
	}
	name_info.Time = &d

	return name_info, nil
}

// Get File to save name according to option
// return file_name to use
func GetFileName(file string, sensor_type string, sensor_serial string, conf *pconf.Config) (string) {
	// init file_to_save
	file_to_save := file
	name_info, _ := GetFileNameInfo(file) //TODO manage error
	extension := name_info.Extension
	suffix := name_info.Suffix

	// Date transform
	if 0 != conf.Out_date_format_type {
		file_to_save = name_info.Prefix + pformat.GetTimeFormated(name_info.Time, &conf.Date_time_format, &conf.Date_time_format) + name_info.Extension
	}

	// Set log as prefix
	if conf.Log_suffix && ("log_" == name_info.Prefix) {
		file_to_save = strings.TrimPrefix(file_to_save, "log_")
		file_to_save = strings.TrimSuffix(file_to_save, extension)
		file_to_save += "_log" + extension
		suffix = "_log"
	}

	// Replace cvs suffix
	if (conf.Replace_csv_by != "") && ( strings.HasSuffix(file, ".csv") ) { // Replace file ".csv" extension by user choice
		file_to_save = strings.TrimSuffix(file_to_save, ".csv")
		file_to_save += conf.Replace_csv_by
		extension = conf.Replace_csv_by
	}
	// Replace txt suffix
	if (conf.Replace_txt_by != "") && ( strings.HasSuffix(file, ".txt") ) { // Replace file ".txt" extension by user choice
		file_to_save = strings.TrimSuffix(file_to_save, ".txt")
		file_to_save += conf.Replace_txt_by
		extension = conf.Replace_txt_by
	}

	// Add type-serial prefix
	if conf.Serial_prefix {
		file_to_save = sensor_type + "-" + sensor_serial + "_" + file_to_save // add "type-serial" + "_" prefix
	}

	// Add type-serial suffix
	if conf.Serial_suffix {
		file_to_save = strings.TrimSuffix(file_to_save, extension) // Remove current extension
		file_to_save = strings.TrimSuffix(file_to_save, suffix) // Remove current suffix too
		file_to_save += "_" + sensor_type + "-" + sensor_serial + suffix + extension // Add "_type-serial" + file suffix + file extension 		-> 20171219_CPX3-121010_presence.csv
	} else if conf.Serial_suffix_final { 	// OR Add type-serial suffix final
		file_to_save = strings.TrimSuffix(file_to_save, extension) // Remove current extension
		file_to_save += "_" + sensor_type + "-" + sensor_serial + extension // Add "_type-serial" + file extension 													-> 20171219_presence_CPX3-121010.csv
	}

	return file_to_save
}

// Get File path to save name according to config
// return file_path to use
func GetFilePath(file string, sensor_type string, sensor_serial string, url_path string, conf *pconf.Config) (string, error) {
	// Default save directory
	file_path := SanitizeDir(conf.Out_dir)
	if conf.No_sanitize {
		file_path = conf.Out_dir
	}

	// Check if url path need to be added to file_path
	if conf.Out_dir_add_path {
		if conf.No_sanitize {
			file_path += url_path
		} else if !strings.HasPrefix(url_path, "/") {
			file_path += SanitizeDir(url_path)
		} else {
			file_path += SanitizeDir(url_path[1:len(url_path)]) // because file_path should already end with a "/"
		}
	}

	// Add "type-Serial" directory according to config
	if !conf.No_serial_dir {
		file_path += sensor_type + "-" + sensor_serial + "/"
	}

	// Add yyyy/mm/dd according to config
	if conf.Save_dir_yyyy || conf.Save_dir_yy || conf.Save_dir_mm || conf.Save_dir_dd {
		name_info, err := GetFileNameInfo(file)
		if nil != err {
			return "", err
		}

		// Year directory
		if conf.Save_dir_yyyy {
			file_path += name_info.Yyyy + conf.Save_dir_sep
		} else if conf.Save_dir_yy {
			file_path += name_info.Yy + conf.Save_dir_sep
		}
		// Month directory
		if conf.Save_dir_mm {
			file_path += name_info.Mm + conf.Save_dir_sep
		}
		// Day directory
		if conf.Save_dir_dd {
			file_path += name_info.Dd + conf.Save_dir_sep
		}

		// Fix path if separator not "/"
		if ("/" != conf.Save_dir_sep) && (strings.HasSuffix(file_path, conf.Save_dir_sep)) {
			file_path = strings.TrimSuffix(file_path, conf.Save_dir_sep)
		}
	}

	return SanitizeDir(file_path), nil
}

// Get File content modified according to config
// take file_content as a string and transform it according to config
func GetFileContent(file_name string, file_content string, conf *pconf.Config) (string) {
	// Check replace comma before write
	if (conf.Replace_comma_by != "") && (!strings.HasSuffix(file_name, ".jpg")) && (!strings.HasSuffix(file_name, ".bmp") ) { // Do not replace anything if file is .jpg type
		// Check if we have already the replace char in string
		// In this case we proceed to a double replace (to not erase a replace char)
		tmp_replace_char := ""
		if strings.Contains(file_content, conf.Replace_comma_by) && (false == conf.No_replace_comma_replace) {
			tmp_replace_char = "&&&&&&&&" // "&" is a forbidden char in Comptipix protocol so "&&&&&&&&" is very unlikely to append
			if conf.Replace_comma_by == tmp_replace_char { // Seriously ?
				tmp_replace_char += "&" // ok, one more ... and no don't want to check if one more still exist
			}
			file_content = strings.Replace(file_content, conf.Replace_comma_by, tmp_replace_char, -1) // Replace the replace comma char by "&&&&&&&&"
		}
		file_content = strings.Replace(file_content, ",", conf.Replace_comma_by, -1) // Replace comma by user replace char

		// End double replace
		if tmp_replace_char != "" {
			file_content = strings.Replace(file_content, tmp_replace_char, ",", -1) // Replace replaced separator ... by comma
		}
	}

	// User want a CR or CRLF change (thoug it's a bad idea)
	if (0 != conf.Out_file_end_line) && (!strings.HasSuffix(file_name, ".jpg")) && (!strings.HasSuffix(file_name, ".bmp") ) { // Avoid messing jog and bmp files
		// Replace LF by CR or CRLF or nothing
		file_content = strings.Replace(file_content, "\n", pformat.GetConfEndOfLine(conf.Out_file_end_line), -1)
	}

	return file_content
}

// Get file existence + size
// Return a value < 0 for a not exist file or directory: -1 for inexistence file, -2 for permission error
// Return the size otherwise
func GetFileSize(file_full_path string) (int64) {
	// Check for directory existence
	if fi, fi_err := os.Stat(file_full_path); fi_err != nil {
		if os.IsNotExist(fi_err) {
			return -1		// file does not exist
		} else {
			return -2		// other error (like permission error)
		}
	} else {
		return fi.Size() // Return file size
	}
}
