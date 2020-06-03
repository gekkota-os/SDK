// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o zuller.exe zuller.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o zuller.exe zuller.go
//
// * RUN :
// ./zuller -h
//

package main

import (
	"log"
	"strings"
	"strconv"
	"compress/gzip"
	// "encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"time"
	"path/filepath"
	"flag"
)

// Header parse status (because eurecam header is 2 line : first the header, 2nd must be "fichier de comptage v2" or header is not valid)
type SiteChainHeader int32
const (
	SITE_CHAIN_IDLE   			SiteChainHeader = iota
	SITE_CHAIN_IN_PORGRESS
	SITE_CHAIN_DONE
)

// FileEndLine to use
type FileEndLine int
const (
	END_LINE_LF			FileEndLine = iota 			// 0 : LF
	END_LINE_CR												// 1 : CR
	END_LINE_CRLF											// 2 : CRLF
	END_LINE_LFCR											// 3 : CRLF ... because we never know what client will ask
)

// File info
type FileInfo struct {
	Site							string					`json:"site"`						// site name
	Chain							string					`json:"chain"`						// chain name
	Site_after					string					`json:"site_after"`				// site part after Dst_chain_site_after separator
	Chain_after					string					`json:"chain_after"`				// chain part after Dst_chain_site_after separator
	Channel_name				map[int]string			`json:"channel_name"`			// channel name
	Enabled_channel			[]int 					`json:"enabled_channel"`		// channel id (id is position) that has name != ""
	Access_channel				[]int						`json:"access_channel"`			// channel id (id is position) that are access and taking into account for summing Entry and Exits
	All_channel					[]int						`json:"all_channel"`				// channel id (id is position), all
	Entry_sum					int						`json:"entry_sum"`				// entry sum
	Exit_sum						int						`json:"exit_sum"`					// exit sum
}

// Counting data : info of 1 line of 1 channel
type CountingData struct {
	Entry							int						`json:"Entry"`						// entry
	Exit							int						`json:"Exit"`						// exit
}

// A parsed file
type FileParsed struct {
	FileInfo
	FileData						map[int64][]CountingData									// map of slice counting info -> index = unix timestamp -> position in slice = channel id
	FileTimestampSteps		[]int64															// All timestamp step list
	FileTimestamp				int64																// File timestamp (starting at 00:00:00)
}

// Timestamped counting data for all channel
type TimestampedCountingData struct {
	Timestamp					int64																// unix timestamp
	ChannelData					[]CountingData													// Data for all used channel
}

// Config struct
type Config struct {
	Src_dir						string					`json:"src_dir"`					// source directory
	Dst_dir						string 					`json:"dst_dir"`					// destination directory
	Nb_file						int						`json:"nb_file"`					// Number of file to try to sum since today
	Data_resolution			int64						`json:"data_resolution"`		// data resolution wanted in seconds
	Skip_today					bool						`json:"skip_today"`				// true to skip today files
	Verbose						bool						`json:"verbose"`					// true to be verbose
	Use_enabled_channel		bool						`json:"use_enabled_channel"`	// true to generate zul file with only enabled channels from Eurecam file (an enabled channel is a channel that has his name different from empty string "")
	Use_access_channel		bool						`json:"use_access_channel"`	// true to generate zul file with only access channels from Eurecam file
	Use_all_channel			bool						`json:"use_all_channel"`		// true to generate zul file with all channels from Eurecam file
	Gzip_older_than			int						`json:"gzip_older_than"`		// >0 to gzip files that are older than this value (value is number of day from today)
	Gzip_ext_blacklist_src	string					`json:"gzip_ext_blacklist_src"`// list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip,.size')
	Gzip_ext_blacklist		[]string					`json:"gzip_ext_blacklist"`	// list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')
	Periodic_minute			int						`json:"periodic_minute"`		// a value >=0 to restart process every x minute
	Out_file_end_line			int						`json:"out_file_end_line"`		// new line charactere see \ref FileEndLine
	Out_strip_header			bool						`json:"Out_strip_header"`		// true to strip file header
	Out_file_ext				string					`json:"Out_file_ext"`			// file extension output
}

// Functions
// --------------------------

// Display config
func displayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("src_dir:                     %s\n", conf.Src_dir)
	fmt.Printf("dst_dir:                     %s\n", conf.Dst_dir)
	fmt.Printf("nb_file:                     %d\n", conf.Nb_file)
	fmt.Printf("skip_today:                  %t\n", conf.Skip_today)
	fmt.Printf("verbose:                     %t\n", conf.Verbose)
	fmt.Printf("use_enabled_channel:         %t\n", conf.Use_enabled_channel)
	fmt.Printf("use_access_channelose:       %t\n", conf.Use_access_channel)
	fmt.Printf("use_all_channel:             %t\n", conf.Use_all_channel)
	fmt.Printf("gzip_older_than:             %d\n", conf.Gzip_older_than)
	fmt.Printf("gzip_ext_blacklist_src:      %s\n", conf.Gzip_ext_blacklist_src)
	fmt.Printf("gzip_ext_blacklist:          %+v\n", conf.Gzip_ext_blacklist)
	fmt.Printf("periodic_minute:             %d\n", conf.Periodic_minute)
	fmt.Printf(" ----\n")
}

// Get configured end of line char
func getConfEndOfLine(conf *Config) (string) {
	line_end := "\n"
	switch FileEndLine(conf.Out_file_end_line) {
		case END_LINE_CR: 			// 1
		line_end = "\r"
		case END_LINE_CRLF: 			// 2
		line_end = "\r\n"
		case END_LINE_LFCR: 			// 3
		line_end = "\n\r"
		default:
		line_end = "\n"
	}
	return line_end
}

// Do a verbose print (according to config)
// and produce log (if log set to true)
func PrintAndLog(msg string, conf *Config, do_log bool) {
	if conf.Verbose {
		fmt.Printf(msg)
	}
	if do_log {
		log.Printf(msg)
	}
}

// Sanitize directory (Add the last '/' if not present)
func SanitizeDir(dir string) (string) {
	// Check dir ends with '/'
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	return dir
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

// Return true + index position if found, false + anything otherwise
func sliceIntContains(s []int, e int) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}
// Return true + index position if found, false + anything otherwise
func sliceStringContains(s []string, e string) (bool, int) {
	for index, a := range s {
		if a == e {
			return true, index
		}
	}
	return false, -1
}

// Return true if date check is between start and end
func inTimeSpan(start time.Time, end time.Time, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// Get string after first occurance
func getStringAfter(str string, str_after string) (str_result string) {
	index := strings.LastIndex(str, str_after)
	if -1 == index {
		return str
	}
	runes := []rune(str)	// Take substring of first word with runes. This handles any kind of rune in the string.
	safe_substring := string(runes[index+len(str_after):len(str)])	// Convert back into a string from rune slice.

	return safe_substring
}

// Get now (today set at midnight)
func getNowMidnight() (time.Time) {
	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return rounded
}

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

// Gzip or not file_path according to config
func gzipOldFile(f_path string, conf *Config, check_file_date bool) {
	if conf.Gzip_older_than <= 0 {
		return
	}

	// Get file info
	_, f_name := filepath.Split(f_path)
	ext_blacklisted, _ := sliceStringContains(conf.Gzip_ext_blacklist, filepath.Ext(f_name))
	if ext_blacklisted {
		// PrintAndLog(fmt.Sprintf("-> gzipOldFile no need to gzip %s !\n", f_path), conf, false)// Verbose only
		return
	}

	if check_file_date {
		file_date := getDateFromEurecam(strings.TrimSuffix(f_name, filepath.Ext(f_name)))
		now := getNowMidnight()
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

// Get a date from a string date Eurecam format (like "20171122") from a date type
func getDateFromEurecam(date_str string) (time.Time) {
	now := getNowMidnight() // use now by default, or if user date_str is invalid

	if date_str != "" {
		time_format_start := "20060102"
		d_date, d_date_err := time.Parse(time_format_start, date_str)
		if nil == d_date_err {
			now = d_date
		} else {
			log.Printf("Invalid date_str (%s)\n", date_str) // Print error day
		}
	}

	return now
}

// Get a date from a string inside Eurecam format (like "11/01/2018,01:42:16")
func parseDateEurecamFile(date_str string) (time.Time, error) {
	// Select ISO format or OLD format
	time_format := "2006-01-02,15:04:05" 	// ISO date by default like "2018-01-11,01:42:16"
	if 2 == strings.Count(date_str, "/") {	// OLD format like "11/01/2018,01:42:16"
		time_format = "02/01/2006,15:04:05"
	}

	// Parse date
	d, d_err := time.Parse(time_format, date_str)
	if nil == d_err {
		return d, nil 		// Date read OK
	}

	return getNowMidnight(), d_err	// Date read error
}

// Parse an Eurecam file an return the entries + exits sum
func walkParseFile(conf *Config) (filepath.WalkFunc)  {
	// Use a closure to pass argument to walk function (pass the config)
	return func(fpath string, info os.FileInfo, err error) (error) {
		if err != nil {
			log.Printf("Error : sumFile dir -> %s \n", err.Error())
			return nil
		}
		if info.IsDir() {
			return nil // Don't care of directory
		} else {
			// We have a file
			dir, file_name := filepath.Split(fpath)
			last_dir := filepath.Base(dir)

			// Check we want this file (check for an Eurecam counting file)
			gziped := false
			PrintAndLog(fmt.Sprintf("File: %s (last_dir:%s, name:%s)", fpath, last_dir, file_name), conf, false)// Verbose only
			if (12 != len(file_name)) || (".csv" != filepath.Ext(file_name)) { // Skip not counting file -> like "20171207.csv"
				if (15 == len(file_name)) && (strings.HasSuffix(file_name, ".csv.gz")) {
					gziped = true
				} else {
					gzipOldFile(fpath, conf, false) // Always gzip this files
					PrintAndLog(fmt.Sprintf("--> SKIP file -> not a counting file!!\n"), conf, false)// Verbose only
					return nil
				}
			}

			// Get date from file name
			file_suffix := ".csv"
			if gziped {
				file_suffix = ".csv.gz"
			}
			date_file := getDateFromEurecam(strings.TrimSuffix(file_name, file_suffix))

			// Check file is in TimeSpan we want to test
			if conf.Nb_file > 0 {
				date_end := getNowMidnight()
				date_start := date_end.Add(time.Duration(-1*conf.Nb_file*24) * time.Hour)

				// Check skip today
				if conf.Skip_today && (date_file.Year() == date_end.Year()) && (date_file.Month() == date_end.Month()) && (date_file.Day() == date_end.Day()) {
					PrintAndLog(fmt.Sprintf("--> SKIP file date(%s) -> because skip_today(%s)!!\n", date_file.Format("2006-01-02 15:04:05"), date_end.Format("2006-01-02 15:04:05")), conf, false)//Verbose only
					return nil
				}

				// Check date is in day wanted
				if !inTimeSpan(date_start, date_end.Add(12 * time.Hour), date_file) { // Add 12hours to date_end which is today at midnight (to be sure today file will be processed)
					gzipOldFile(fpath, conf, true) // Gzip file only if is after n day
					PrintAndLog(fmt.Sprintf("--> SKIP file date(%s) -> not in nb file wanted (from:%s to %s)!!\n", date_file.Format("2006-01-02 15:04:05"), date_start.Format("2006-01-02 15:04:05"), date_end.Format("2006-01-02 15:04:05")), conf, false)//Verbose only
					return nil
				}
			}

			// Open file to parse
			f, err := os.Open(fpath)
			if err != nil {
				PrintAndLog(fmt.Sprintf("--> ERROR Opening : %s\n", err.Error()), conf, true)// Verbose + Do Log
				return nil
			}
			defer f.Close()

			// fmt.Printf("Here %s \n",file_name)//DEBUG

			// Get parsed result  (from a gziped file or a normale one)
			file_parsed := FileParsed{}
			if gziped {
				gz, gz_err := gzip.NewReader(f)
				if gz_err != nil {
					PrintAndLog(fmt.Sprintf("--> ERROR unZIPing : %s\n", gz_err.Error()), conf, true)// Verbose + Do log
					return nil
				}
				file_parsed = getParsedFile(gz, &date_file, conf.Data_resolution) // Get sum + file info from gziped file
				gz.Close()
			} else {
				file_parsed = getParsedFile(f, &date_file, conf.Data_resolution) // Get sum + file info from vanilla file
			}
			f.Close()

			// Write new file in destination
			zul_name, zul_content := getZul(&file_parsed, conf)
			dst_path := SanitizeDir(conf.Dst_dir)
			os.MkdirAll(dst_path, 0777) // always check and create directory before write
			f_path_to_write := dst_path + zul_name
			f_result, f_result_err := os.Create(f_path_to_write)
			if f_result_err != nil {
				PrintAndLog(fmt.Sprintf("--> ERROR Wrinting result : to : %s -> %s \n", f_path_to_write, f_result_err.Error()), conf, true)// Verbose + do LOG
				return nil
			}
			f_result.Write([]byte(zul_content))
			f_result.Close()
			gzipOldFile(fpath, conf, true) // Gzip file only if is after n day

			PrintAndLog(fmt.Sprintf(" -> Zulled to %s \n", zul_name), conf, true)// Verbose + NO LOG
			return nil
		}
	}
}

// Get parsed data from a file content
// Parsed data are stored at the wanted resolution
func getParsedFile(fdata io.Reader, t_date_file *time.Time, resolution int64) (file_parsed FileParsed) {
	file_info := FileInfo{Site:"", Chain:"", Channel_name:make(map[int]string), Entry_sum:0, Exit_sum:0}
	file_parsed = FileParsed{FileInfo:file_info, FileData:make(map[int64][]CountingData)}

	// Round the date_file to day at 00:00:00 (it should be the case but it's better to be sure)
	date_file := time.Date(t_date_file.Year(), t_date_file.Month(), t_date_file.Day(), 0, 0, 0, 0, time.UTC)
	file_parsed.FileTimestamp = date_file.UTC().Unix()

	// Temporary store of Site chain and Site name,
	// because a Site Chain/Name line must be validated by a line "fichier de comptage v2"
	// so we store temporary read site and chain, waiting for validation
	tmp_site := ""
	tmp_chain := ""

	// Build resolution array
	steps := []int64{}
	var step_current int64 = 0
	for step_current<86400 {
		steps = append(steps, step_current)
		step_current += resolution
	}
	file_parsed.FileTimestampSteps = steps
	// fmt.Printf("\nfile date: %s = %d, resolution: %d\n", date_file.String(), file_parsed.FileTimestamp, resolution)//DEBUG
	// fmt.Printf("%+v\n", steps)//DEBUG

	// Use scanner to read file line by line
	scanner := bufio.NewScanner(fdata)
	site_chain_header_state := SITE_CHAIN_IDLE
	for scanner.Scan() {
		// Read one line
		line := scanner.Text()
		line_split := strings.Split(line, ",")

		// Check site chain header preceed a "fichier de comptage v2" line
		if SITE_CHAIN_IN_PORGRESS == site_chain_header_state {
			if "fichier de comptage v2" == line {
				site_chain_header_state = SITE_CHAIN_DONE
				file_parsed.Site = tmp_site
				file_parsed.Chain = tmp_chain
				continue
			} else {
				site_chain_header_state = SITE_CHAIN_IDLE
			}
		}

		// Check site/chain header
		if (1 == strings.Count(line, ",")) {
			if SITE_CHAIN_DONE == site_chain_header_state {
				fmt.Printf("Double header ! (NOTE: it may be normal)\n")	// We update Site/Chain header if there is another header
			}
			if len(line_split) > 0 {
				tmp_site = line_split[0]
			}
			if len(line_split) > 1 {
				tmp_chain = line_split[1]
			}
			site_chain_header_state = SITE_CHAIN_IN_PORGRESS
			continue
		}

		// Check Channel type
		if (2 == strings.Count(line, ",")) && (3 == len(line_split)) {
			c_i, c_err := strconv.Atoi(line_split[0])
			if (nil == c_err) && (c_i > 0) {
				// Add/Remove channel from access channel list
				access_contains_already, access_contains_pos := sliceIntContains(file_parsed.FileInfo.Access_channel, c_i)
				if ("acces" == line_split[2]) && (!access_contains_already) {	// Add
					file_parsed.FileInfo.Access_channel = append(file_parsed.FileInfo.Access_channel, c_i)	// Is access
				}
				if ("acces" != line_split[2]) && access_contains_already {		// Remove
					file_parsed.FileInfo.Access_channel = append(file_parsed.FileInfo.Access_channel[:access_contains_pos], file_parsed.FileInfo.Access_channel[access_contains_pos+1:]...) // Remove the channel preserving the slice order see https://github.com/golang/go/wiki/SliceTricks
				}

				// Add/Remove channel from enabled channel list
				enabled_contains_already, enabled_contains_pos := sliceIntContains(file_parsed.FileInfo.Enabled_channel, c_i)
				if ("" != line_split[1]) && (!enabled_contains_already) { 	// Add
					file_parsed.FileInfo.Enabled_channel = append(file_parsed.FileInfo.Enabled_channel, c_i)	// Is enabled
				}
				if ("" == line_split[1]) && enabled_contains_already {		// Remove
					file_parsed.FileInfo.Enabled_channel = append(file_parsed.FileInfo.Enabled_channel[:enabled_contains_pos], file_parsed.FileInfo.Enabled_channel[enabled_contains_pos+1:]...) // Remove the channel preserving the slice order see https://github.com/golang/go/wiki/SliceTricks
				}

				// Add/Update Channel name + all channel
				file_parsed.FileInfo.Channel_name[c_i] = line_split[1]
				file_parsed.FileInfo.All_channel = append(file_parsed.FileInfo.All_channel, c_i)	// Add to all channel

				// Nothing more to do with this line
				continue
			}
		}

		// Chech for line "Date,Heure,E,S"
		if strings.HasPrefix(line, "Date,Heure,") {
			continue // Nothing to do
		}

		// Check for a counting -> sum only access channel
		if (strings.Count(line, ",") >= 3) {
			// Line format is like :
			// Date,Heure,E,S
			// 07/12/2017,00:00:00,0,0

			// Add channel counting to data map at good resolution
			// ---
			// Read date-time
			line_time, line_time_err := parseDateEurecamFile(line_split[0]+","+line_split[1])
			if nil == line_time_err {
				// Check date is the same as date_file
				if (line_time.Year() != date_file.Year()) || (line_time.Month() != date_file.Month()) || (line_time.Day() != date_file.Day()) {
					//TODO : we should log this : this appen on SDcard FAT corruption -> the content of another file in 1 file
					fmt.Printf("line: date  %s NOT in file %s\n", line_time.Format("2006-01-02"), date_file.Format("2006-01-02"))
					continue
				}

				// Get resolution position
				line_time_sec := line_time.Unix() - date_file.Unix()	// Keep only second for today
				res_pos := steps[line_time_sec / resolution]
				// fmt.Printf("line: %s  ->  %d = %d (%s)\n", line, line_time_sec, res_pos, line_time.Format("2006-01-02 15:04:05"))//DEBUG

				// Parse all counting column
				counting_data := []CountingData{}
				for j:=2; j<len(line_split)-1; j+=2 {
					tmp_e, tmp_e_err := strconv.Atoi(line_split[j])
					tmp_x, tmp_x_err := strconv.Atoi(line_split[j+1])
					if (nil == tmp_e_err) && (nil == tmp_x_err) {
						counting_data = append(counting_data, CountingData{Entry:tmp_e, Exit:tmp_x})
					} else {
						// TODO: there is error inside file -> Log it!
						fmt.Printf("There error in file %s !!!\n", date_file.Format("2006-01-02"))
						if (nil != tmp_e_err) && (nil != tmp_x_err) {
							counting_data = append(counting_data, CountingData{Entry:0, Exit:0})		// Error on both entry + exit
						} else if nil != tmp_e_err {
							counting_data = append(counting_data, CountingData{Entry:0, Exit:tmp_x})// Error on entry
						} else {
							counting_data = append(counting_data, CountingData{Entry:tmp_e, Exit:0})// Error on exit
						}
					}
				}

				// Set it to map or add it to an existing map entrie
				map_val, map_ok := file_parsed.FileData[res_pos]
				if map_ok {
					// Add slice
					new_counting_data := []CountingData{}
					for k := range(file_parsed.FileData[res_pos]) {
						new_e := map_val[k].Entry + counting_data[k].Entry
						new_x := map_val[k].Exit + counting_data[k].Exit
						// fmt.Printf("---> Add-Add %d(%d) : E:%d+%d=%d , X:%d+%d=%d\n", res_pos, k, map_val[k].Entry, counting_data[k].Entry, new_e, map_val[k].Exit, counting_data[k].Exit, new_x)//DEBUG
						new_counting_data = append(new_counting_data, CountingData{Entry:new_e, Exit:new_x})
					}
					// fmt.Printf("---> Add Slice at res %d =\n", res_pos)
					// fmt.Printf("%+v\n", new_counting_data)
					// fmt.Printf("+\n")
					// fmt.Printf("%+v\n", file_parsed.FileData[res_pos])

					file_parsed.FileData[res_pos] = new_counting_data

					// fmt.Printf("=\n")
					// fmt.Printf("%+v\n", file_parsed.FileData[res_pos])
				} else {
					// Set slice
					file_parsed.FileData[res_pos] = counting_data

					// fmt.Printf("--> Set Slice at res %d =\n", res_pos)
					// fmt.Printf("%+v\n", file_parsed.FileData[res_pos])
				}
			}


			// Do sum (sum only access channel)
			// ---
			for i := range(file_parsed.FileInfo.Access_channel) {	// loop through Access_channel
				channel_col := file_parsed.FileInfo.Access_channel[i]*2
				if len(line_split) >= (channel_col+1) {
					// Add channel Entry
					entry_add, entry_add_err := strconv.Atoi(line_split[channel_col])
					if (nil == entry_add_err) {
						file_parsed.Entry_sum += entry_add
					}
					// Add channel Exit
					exits_add, exits_add_err := strconv.Atoi(line_split[channel_col+1])
					if (nil == exits_add_err) {
						file_parsed.Exit_sum += exits_add
					}
				}
			}
		}
	}

	return file_parsed
}

// Get Timestamped data from a file parsed
// Parsed data are stored at the wanted resolution
func getTimestampedData(file_parsed *FileParsed) (time_data []TimestampedCountingData) {
	// fmt.Printf("FILE timestamp = %d -> %s\n", file_parsed.FileTimestamp, time.Unix(file_parsed.FileTimestamp, 0).UTC().Format("2006-01-02 15:04:05"))//DEBUG

	for i := range(file_parsed.FileTimestampSteps) {
		// Get map data
		map_val, map_ok := file_parsed.FileData[file_parsed.FileTimestampSteps[i]]
		if map_ok {
			// fmt.Printf("getTimestampedData %d -> %s\n",file_parsed.FileTimestamp, time.Unix(file_parsed.FileTimestamp + file_parsed.FileTimestampSteps[i], 0).UTC().Format("2006-01-02 15:04:05"))//DEBUG
			counting_data := TimestampedCountingData{Timestamp:file_parsed.FileTimestamp + file_parsed.FileTimestampSteps[i], ChannelData:map_val}
			time_data = append(time_data, counting_data)
		}
	}

	return time_data
}

// Get Zul data name file + zul content
func getZul(file_parsed *FileParsed, conf *Config) (zul_name string, zul_content string) {
	time_data := getTimestampedData(file_parsed)

	// use all channel or just enabled or just access
	channel_to_check := make([]int, len(file_parsed.FileInfo.All_channel), cap(file_parsed.FileInfo.All_channel))
	if conf.Use_all_channel {
		copy(channel_to_check, file_parsed.FileInfo.All_channel)
	} else if conf.Use_access_channel {
		channel_to_check = make([]int, len(file_parsed.FileInfo.Access_channel), cap(file_parsed.FileInfo.Access_channel))
		copy(channel_to_check, file_parsed.FileInfo.Access_channel)
	} else {
		channel_to_check = make([]int, len(file_parsed.FileInfo.Enabled_channel), cap(file_parsed.FileInfo.Enabled_channel))
		copy(channel_to_check, file_parsed.FileInfo.Enabled_channel)
	}

	// Content
	zul_end_line := getConfEndOfLine(conf)
	if !conf.Out_strip_header {
		zul_content = fmt.Sprintf("%d doors%s", len(channel_to_check), zul_end_line)
	}

	// fmt.Printf("Channel to check: %d , %d\n", len(channel_to_check), cap(channel_to_check))
	// fmt.Printf("%+v\n", channel_to_check)

	last_date_line := time.Unix(file_parsed.FileTimestamp, 0).UTC()
	for j := range(time_data) {
		date_line := time.Unix(time_data[j].Timestamp, 0).UTC()
		zul_door := 1
		for k := range(time_data[j].ChannelData) {
			is_wanted, _ := sliceIntContains(channel_to_check, k+1) // Channel start at 1, so use k+1 to look for a channel
			if is_wanted {
				// fmt.Printf("%d -> %d:%d = %d / %d (%d)\n",time_data[j].Timestamp, j,k, time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, channel_to_check[k])//DEBUG
				zul_content += fmt.Sprintf("%d\t%s\t%s\t%04d\t%04d\t1%s", zul_door, date_line.Format("02/01/2006"), date_line.Format("15:04"), time_data[j].ChannelData[k].Entry, time_data[j].ChannelData[k].Exit, zul_end_line)
				zul_door++
			}
		}
		last_date_line = date_line
	}

	// File name
	// '-> zul name is MMDDHHMM.zul
	zul_name =  fmt.Sprintf("%s.%s", last_date_line.Format("01021504"), conf.Out_file_ext)

	return zul_name, zul_content
}

// Loop to zul all files
func zulAllFiles(conf *Config) (err error) {
	err = filepath.Walk(conf.Src_dir, walkParseFile(conf))
	if err != nil {
		PrintAndLog(fmt.Sprintf("Error : zulAllFiles: %s \n", err.Error()), conf, true)// Verbose + do LOG
	}

	return err
}

// Read config from conf file
func readConfigFile(conf *Config, json_path string) (read_ok bool) {
	_, fi_err := os.Stat(json_path)
	if nil == fi_err {
		// Read file
		f, f_err := ioutil.ReadFile(json_path)
		if nil == f_err {
			// Decode JSON
			json_err := json.Unmarshal(f, conf)
			if nil == json_err {
				conf.Src_dir = filepath.ToSlash(conf.Src_dir) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
				conf.Dst_dir = filepath.ToSlash(conf.Dst_dir)
				conf.Gzip_ext_blacklist = strings.Split(conf.Gzip_ext_blacklist_src, ",")
				return true
			} else {
				fmt.Printf("BAD JSON config file: %s\n", json_err.Error())
			}
		}
	}
	return false
}

// Main
// Parse option and start to read files
func main() {
	// Read arg
	// ----------------
	// help_demo := "This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE"
	default_json_config_path := "./zuller.json"

	// Flags
	Src_dir := flag.String("src_dir", "./Data/", "Source directory that store files to parse")
	Dst_dir := flag.String("dst_dir", "./Dst/", "Destination directory for result sum files")
	Nb_file := flag.Int("nb_file", 60, "Number of file to consider from today, -1 for all days (Note: file not present in destination will still be written)")
	Data_resolution := flag.Int64("data_resolution", 1800, "Data resolution wanted in seconds")
	Skip_today := flag.Bool("skip_today", false, "Add this flag to skip today file (file in destination will arrive only when the day is over")
	Verbose := flag.Bool("verbose", false, "Set it to be verbose")
	Use_enabled_channel := flag.Bool("use_enabled_channel", false, "Set to generate zul file with only enabled channels from Eurecam file -> DEFAULT if no choice made, (an enabled channel is a channel that has his name different from empty string)")
	Use_access_channel := flag.Bool("use_access_channel", false, "Set to generate zul file with only access channels from Eurecam file")
	Use_all_channel := flag.Bool("use_all_channel", false, "Set to generate zul file with all channels from Eurecam file")
	Gzip_older_than := flag.Int("gzip_older_than", 35, "Set a value greater than 0 to gzip files older than this value (value is number of day from today)")
	Gzip_ext_blacklist_src := flag.String("gzip_ext_blacklist_src", ".jpg,.gz,.zip,.gzip,.size", "List of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')")
	Periodic_minute := flag.Int("periodic_minute", -1, "Minute periodicity to restart parsing process (a value negative or 0 will not restart process")
	Out_strip_header := flag.Bool("out_strip_header", false, "Set this flag to strip header from output file (Default: disabled)")
	Out_file_end_line := flag.Int("out_file_end_line", 0, "File en line char:\n\t0:LF (Default)\n\t1:CR\n\t2:CRLF\n\t3:LFCR")
	Log_file := flag.String("log", "./log.log", "Destination log file")
	Config_json := flag.Bool("config_json", false, "Set this flag to read config from zuller.json file (Default no read from zuller.json)")
	Config_json_path := flag.String("config_json_path", default_json_config_path, "full path to file config including his name (ex:'my/super/path/to/my_config.json')")
	Out_file_ext := flag.String("out_file_ext", "zul", "File output extension (Default: 'zul'")
	// help := flag.Bool("help", false, "Print help usage and demo")
	flag.Parse()

	// read config from JSON
	var conf Config
	if (true == *Config_json) || (default_json_config_path != *Config_json_path) {
		// Read config file
		if !readConfigFile(&conf, *Config_json_path) {
			fmt.Printf("Error reading json file ABORT\n")
			return
		} else {
			fmt.Printf("Read json config (%s) OK\n", *Config_json_path)
		}
	} else {
		// Because we have to do it after flag.Parse()
		conf.Src_dir = filepath.ToSlash(*Src_dir)// Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
		conf.Dst_dir = filepath.ToSlash(*Dst_dir)
		conf.Nb_file = *Nb_file
		conf.Data_resolution = *Data_resolution
		conf.Skip_today = *Skip_today
		conf.Verbose = *Verbose
		conf.Use_enabled_channel = *Use_enabled_channel
		conf.Use_access_channel = *Use_access_channel
		conf.Use_all_channel = *Use_all_channel
		// conf.Log_file
		conf.Gzip_older_than = *Gzip_older_than
		conf.Gzip_ext_blacklist_src = *Gzip_ext_blacklist_src
		conf.Gzip_ext_blacklist = strings.Split(*Gzip_ext_blacklist_src, ",")
		conf.Periodic_minute = *Periodic_minute
		conf.Out_file_end_line = *Out_file_end_line
		conf.Out_strip_header = *Out_strip_header
		conf.Out_file_ext = *Out_file_ext
	}
	displayConfig(&conf)

	// Set log file
	if *Log_file != "" {
		f, f_err := os.OpenFile(*Log_file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if f_err != nil {
			fmt.Printf("Error opening log file(%s): %v !!!\n", *Log_file, f_err)
		} else {
			defer f.Close()
			log.SetOutput(f)
			fmt.Printf("Zuller started at %s → src: %s , dst: %s , nb_file: %d\n ---\n", time.Now().Format("2006-01-02 15:04:05"), conf.Src_dir, conf.Dst_dir, conf.Nb_file)
		}
	}

	// Zul all file according to config
	if conf.Periodic_minute > 0 {
		// Deamon mode : run each periodic_minute
		PrintAndLog(fmt.Sprintf("Started with periodicity = %d minutes\n", conf.Periodic_minute), &conf, true)// Verbose + do LOG
		for {
			if (true == *Config_json) || (default_json_config_path != *Config_json_path) {	// Try to re-read json config
				if !readConfigFile(&conf, *Config_json_path) {
					fmt.Printf("Error reading json file : config not changed\n")
				}
			}

			// Re-zul file
			PrintAndLog(fmt.Sprintf("Re-Started because periodicity = %d minutes\n", conf.Periodic_minute), &conf, false)// Verbose only
			zulAllFiles(&conf)
			time.Sleep(time.Duration(conf.Periodic_minute) * time.Minute)
		}
	}

	// One shoot mode : run just once
	zulAllFiles(&conf)
}
