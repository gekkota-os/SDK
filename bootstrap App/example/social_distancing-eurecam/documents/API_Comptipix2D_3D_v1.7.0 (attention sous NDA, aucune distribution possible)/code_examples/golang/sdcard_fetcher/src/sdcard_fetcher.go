// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build -o sdcard_fetcher
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o sdcard_fetcher.exe sdcard_fetcher.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o sdcard_fetcher.exe sdcard_fetcher.go
//
// * RUN :
// ./sdcard_fetcher -help
//


package main

import (
	"fmt"
	"log"
	"net/http"
	"bytes"
	"bufio"
	"io/ioutil"
	"time"
	"flag"
	"strconv"
	"strings"
	"os"
	"encoding/json"
	"regexp"
	"path/filepath"
)


// Declaration
// --------------------------

// Config struct
type Config struct {																			// DEFAULT:
	Addr 								string 		`json:"addr"`							// "http://192.168.0.100/"
	Verbose							bool 			`json:"verbose"`						// false
	Nb_file 							int	 		`json:"nb_file"`						// 2
	Pass 								string 		`json:"pass"`							// "pass"
	User 								string 		`json:"user"`							// "user"
	Get_counting					bool			`json:"get_counting"`				// false
	Get_occupancy					bool			`json:"get_occupancy"`				// false
	Get_log 							bool			`json:"get_log"`						// false
	Get_jpg							bool			`json:"get_jpg"`						// false
	Get_config						bool			`json:"get_config"`					// false
	Get_line							bool			`json:"get_line"`						// false
	Get_area							bool			`json:"get_area"`						// false
	Get_track						bool			`json:"get_track"`					// false
	Get_heat							bool			`json:"get_heat"`						// false
	Get_depth						bool			`json:"get_depth"`					// false
	Get_all							bool			`json:"get_all"`						// true
	Save_to							string		`json:"Save_to"`						// ""
	Date_start						string		`json:"Date_start"`					// ""
	Date_stop						string		`json:"Date_stop"`					// ""
	Stop_size						bool			`json:"stop_size"`					// false
	Log_suffix						bool			`json:"log_suffix"`					// false
	Yyyymmdd_to_yymmdd			bool			`json:"yyyymmdd_to_yymmdd"`		// false
	Yyyymmdd_to_yyddmm			bool			`json:"yyyymmdd_to_yyddmm"`		// false
	Yyyymmdd_to_mmyydd			bool			`json:"yyyymmdd_to_mmyydd"`		// false
	Yyyymmdd_to_mmddyy			bool			`json:"yyyymmdd_to_mmddyy"`		// false
	Yyyymmdd_to_ddmmyy			bool			`json:"yyyymmdd_to_ddmmyy"`		// false
	Yyyymmdd_to_ddyymm			bool			`json:"yyyymmdd_to_ddyymm"`		// false
	Yyyymmdd_to_ddmmyyyy			bool			`json:"yyyymmdd_to_ddmmyyyy"`		// false
	Yyyymmdd_to_ddyyyymm			bool			`json:"yyyymmdd_to_ddyyyymm"`		// false
	Yyyymmdd_to_mmddyyyy			bool			`json:"yyyymmdd_to_mmddyyyy"`		// false
	Yyyymmdd_to_mmyyyydd			bool			`json:"yyyymmdd_to_mmyyyydd"`		// false
	Yyyymmdd_to_yyyyddmm			bool			`json:"yyyymmdd_to_yyyyddmm"`		// false
	Yyyymmdd_separator			string 		`json:"yyyymmdd_separator"`		// ""
	File_save_CR					bool 			`json:"file_save_CR"`				// false
	File_save_CRLF					bool 			`json:"file_save_CRLF"`				// false
	Replace_csv_by					string		`json:"replace_csv_by"`				// ""
	Replace_txt_by					string		`json:"replace_txt_by"`				// ""
	Replace_comma_by				string		`json:"replace_comma_by"`			// ""
	No_replace_comma_replace	bool			`json:"no_replace_comma_replace"`// false
	Default_timeout				int			`json:"default_timeout"`			// 40
	Wait_timer						int			`json:"wait_timer"`					// 1
	Serial_prefix					bool			`json:"serial_prefix"`				// false
	Serial_suffix					bool			`json:"serial_suffix"`				// false
	Serial_suffix_final			bool 			`json:"serial_suffix_final"`		// false
	Save_dir_yyyy					bool			`json:"save_dir_yyyy"`				// false
	Save_dir_yy						bool			`json:"save_dir_yy"`					// false
	Save_dir_mm						bool			`json:"save_dir_mm"`					// false
	Save_dir_dd						bool			`json:"save_dir_dd"`					// false
	Save_dir_sep					string		`json:"save_dir_sep"`				// "/"
	No_serial_dir					bool			`json:"no_serial_dir"`				// false
	No_size_check					bool			`json:"no_size_check"`				// false
	No_sanitize						bool			`json:"no_sanitize"`					// false
	No_log     						bool			`json:"no_log"`						// false
	No_logout						bool			`json:"no_logout"`					// false
	Periodic_minute				int			`json:"periodic_minute"`			// a value >=0 to restart process every x minute
}

// All infos we get from a comptipix file name
type FileNameInfo struct {
	Yyyy								string		`json:"yyyy"`							// for file "20171107_presence.csv" it's "2017"
	Yy									string 		`json:"yy"`								// for file "20171107_presence.csv" it's "17"
	Mm									string		`json:"mm"`								// for file "20171107_presence.csv" it's "11"
	Dd 								string 		`json:"dd"`								// for file "20171107_presence.csv" it's "07"
	Prefix 							string		`json:"prefix"`						// for file "20171107_presence.csv" it's "", and for file "log_20171107.csv" it's "log_"
	Suffix 							string		`json:"suffix"`						// for file "20171107_presence.csv" it's "_presence"
	Extension 						string 		`json:"extension"`					// for file "20171107_presence.csv" it's ".csv"
}

// Structure of file to get
type FileType struct {
	Prefix		string		// file prefix (ex: 'log_')
	Suffix		string		// file suffix (ex: '.csv')
}


// Functions
// --------------------------

// Display config
func displayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("address:                  %s\n", conf.Addr)
	fmt.Printf("nb_file:                  %d\n", conf.Nb_file)
	fmt.Printf("date_start:               %s\n", conf.Date_start)
	fmt.Printf("date_stop:                %s\n", conf.Date_stop)
	fmt.Printf("stop_size:                %t\n", conf.Stop_size)
	fmt.Printf("user:                     %s\n", conf.User)
	fmt.Printf("pass:                     %s\n", conf.Pass)
	fmt.Printf("get_counting:             %t\n", conf.Get_counting)
	fmt.Printf("get_occupancy:            %t\n", conf.Get_occupancy)
	fmt.Printf("get_log:                  %t\n", conf.Get_log)
	fmt.Printf("get_jpg:                  %t\n", conf.Get_jpg)
	fmt.Printf("get_config:               %t\n", conf.Get_config)
	fmt.Printf("get_line:                 %t\n", conf.Get_line)
	fmt.Printf("get_area:                 %t\n", conf.Get_area)
	fmt.Printf("get_track:                %t\n", conf.Get_track)
	fmt.Printf("get_heat:                 %t\n", conf.Get_heat)
	fmt.Printf("get_depth:                %t\n", conf.Get_depth)
	fmt.Printf("get_all:                  %t\n", conf.Get_all)
	fmt.Printf("Save_to:                  %s\n", conf.Save_to)
	fmt.Printf("date_start:               %s\n", conf.Date_start)
	fmt.Printf("date_stop:                %s\n", conf.Date_stop)
	fmt.Printf("stop_size:                %t\n", conf.Stop_size)
	fmt.Printf("log_suffix:               %t\n", conf.Log_suffix)
	fmt.Printf("yyyymmdd_to_yymmdd:       %t\n", conf.Yyyymmdd_to_yymmdd)
	fmt.Printf("yyyymmdd_to_yyddmm:       %t\n", conf.Yyyymmdd_to_yyddmm)
	fmt.Printf("yyyymmdd_to_mmyydd:       %t\n", conf.Yyyymmdd_to_mmyydd)
	fmt.Printf("yyyymmdd_to_mmddyy:       %t\n", conf.Yyyymmdd_to_mmddyy)
	fmt.Printf("yyyymmdd_to_ddmmyy:       %t\n", conf.Yyyymmdd_to_ddmmyy)
	fmt.Printf("yyyymmdd_to_ddyymm:       %t\n", conf.Yyyymmdd_to_ddyymm)
	fmt.Printf("yyyymmdd_to_ddmmyyyy:     %t\n", conf.Yyyymmdd_to_ddmmyyyy)
	fmt.Printf("yyyymmdd_to_ddyyyymm:     %t\n", conf.Yyyymmdd_to_ddyyyymm)
	fmt.Printf("yyyymmdd_to_mmddyyyy:     %t\n", conf.Yyyymmdd_to_mmddyyyy)
	fmt.Printf("yyyymmdd_to_mmyyyydd:     %t\n", conf.Yyyymmdd_to_mmyyyydd)
	fmt.Printf("yyyymmdd_to_yyyyddmm:     %t\n", conf.Yyyymmdd_to_yyyyddmm)
	fmt.Printf("yyyymmdd_separator:       %s\n", conf.Yyyymmdd_separator)
	fmt.Printf("file_save_CR:             %t\n", conf.File_save_CR)
	fmt.Printf("file_save_CRLF:           %t\n", conf.File_save_CRLF)
	fmt.Printf("replace_txt_by:           %s\n", conf.Replace_txt_by)
	fmt.Printf("replace_comma_by:         %s\n", conf.Replace_comma_by)
	fmt.Printf("no_replace_comma_replace: %t\n", conf.No_replace_comma_replace)
	fmt.Printf("serial_prefix:            %t\n", conf.Serial_prefix)
	fmt.Printf("serial_suffix:            %t\n", conf.Serial_suffix)
	fmt.Printf("serial_suffix_final:      %t\n", conf.Serial_suffix_final)
	fmt.Printf("save_dir_yyyy:            %t\n", conf.Save_dir_yyyy)
	fmt.Printf("save_dir_yy:              %t\n", conf.Save_dir_yy)
	fmt.Printf("save_dir_mm:              %t\n", conf.Save_dir_mm)
	fmt.Printf("save_dir_dd:              %t\n", conf.Save_dir_dd)
	fmt.Printf("save_dir_sep:             %s\n", conf.Save_dir_sep)
	fmt.Printf("no_serial_dir:            %t\n", conf.No_serial_dir)
	fmt.Printf("no_size_check:            %t\n", conf.No_size_check)
	fmt.Printf("no_sanitize:              %t\n", conf.No_sanitize)
	fmt.Printf("no_log:                   %t\n", conf.No_log)
	fmt.Printf(" ----\n")
}

// Export current config
func exportConfigJSON(conf *Config, file_name string) (error) {
	export_conf_b, export_conf_b_err := json.Marshal(conf)
	if nil == export_conf_b_err {
		export_conf_f, export_conf_f_err := os.Create(file_name)
		if nil == export_conf_f_err {
			export_conf_f.Write(export_conf_b)
			export_conf_f.Close()
			if (conf.Verbose) {
				fmt.Printf("Export current config as sdcard_fetcher.json.new OK\n")
			}
		} else {
			if (conf.Verbose) {
				fmt.Printf("Export current config as sdcard_fetcher.json.new Error writing file !\n")
			}
			return export_conf_f_err
		}
	} else {
		if (conf.Verbose) {
			fmt.Printf("Export current config as sdcard_fetcher.json.new Error generating JSON !\n")
		}
	}

	return export_conf_b_err
}

// Sanitize directory (Add the last '/' if not present)
func SanitizeDir(dir string) (string) {
	// Check dir ends with '/'
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	return dir
}

// Sanitize URL
func SanitizeURL(url string) (string) {
	// Check url has 'http://' or 'HTTP://' or 'https://' or 'HTTPS://'
	// NOTE: https is not supported by Comptipix for now (but may be client has a PC that manage https and forward http:// to Comptipix)
	if (false == strings.HasPrefix(strings.ToLower(url), "http://")) && (false == strings.HasPrefix(strings.ToLower(url), "https://")) {
		url = "http://" + url
	}

	// Finaly fix url to ends with a final '/'
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}

	return url
}

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

	return name_info, nil
}

// Get File to save name according to option
// return file_name to use
func GetFileName(file string, sensor_type string, sensor_serial string, conf *Config) (string) {
	// init file_to_save
	file_to_save := file
	name_info, _ := GetFileNameInfo(file) //TODO manage error
	extension := name_info.Extension
	suffix := name_info.Suffix

	// Date transform
	if conf.Yyyymmdd_to_yymmdd || conf.Yyyymmdd_to_yyddmm || conf.Yyyymmdd_to_mmddyy || conf.Yyyymmdd_to_mmyydd || conf.Yyyymmdd_to_ddmmyy || conf.Yyyymmdd_to_ddyymm || conf.Yyyymmdd_to_ddmmyyyy || conf.Yyyymmdd_to_ddyyyymm || conf.Yyyymmdd_to_mmddyyyy || conf.Yyyymmdd_to_mmyyyydd || conf.Yyyymmdd_to_yyyyddmm || ("" != conf.Yyyymmdd_separator) {
		// Transform date
		s := conf.Yyyymmdd_separator
		if conf.Yyyymmdd_to_yymmdd {
			file_to_save = name_info.Prefix + name_info.Yy + s + name_info.Mm + s + name_info.Dd + name_info.Suffix + name_info.Extension		// yyyymmdd → yymmdd
		} else if conf.Yyyymmdd_to_yyddmm {
			file_to_save = name_info.Prefix + name_info.Yy + s + name_info.Dd + s + name_info.Mm + name_info.Suffix + name_info.Extension		// yyyymmdd → yyddmm
		} else if conf.Yyyymmdd_to_mmddyy {
			file_to_save = name_info.Prefix + name_info.Mm + s + name_info.Dd + s + name_info.Yy + name_info.Suffix + name_info.Extension		// yyyymmdd → mmddyy
		} else if conf.Yyyymmdd_to_mmyydd {
			file_to_save = name_info.Prefix + name_info.Mm + s + name_info.Yy + s + name_info.Dd + name_info.Suffix + name_info.Extension		// yyyymmdd → mmyydd
		} else if conf.Yyyymmdd_to_ddmmyy {
			file_to_save = name_info.Prefix + name_info.Dd + s + name_info.Mm + s + name_info.Yy + name_info.Suffix + name_info.Extension		// yyyymmdd → ddmmyy
		} else if conf.Yyyymmdd_to_ddyymm {
			file_to_save = name_info.Prefix + name_info.Dd + s + name_info.Yy + s + name_info.Mm + name_info.Suffix + name_info.Extension		// yyyymmdd → ddyymm
		} else if conf.Yyyymmdd_to_ddmmyyyy {
			file_to_save = name_info.Prefix + name_info.Dd + s + name_info.Mm + s + name_info.Yyyy + name_info.Suffix + name_info.Extension	// yyyymmdd → ddmmyyyy
		} else if conf.Yyyymmdd_to_ddyyyymm {
			file_to_save = name_info.Prefix + name_info.Dd + s + name_info.Yyyy + s + name_info.Mm + name_info.Suffix + name_info.Extension	// yyyymmdd → ddyyyymm
		} else if conf.Yyyymmdd_to_mmddyyyy {
			file_to_save = name_info.Prefix + name_info.Mm + s + name_info.Dd + s + name_info.Yyyy + name_info.Suffix + name_info.Extension	// yyyymmdd → mmddyyyy
		} else if conf.Yyyymmdd_to_mmyyyydd {
			file_to_save = name_info.Prefix + name_info.Mm + s + name_info.Yyyy + s + name_info.Dd + name_info.Suffix + name_info.Extension	// yyyymmdd → mmyyyydd
		} else if conf.Yyyymmdd_to_yyyyddmm {
			file_to_save = name_info.Prefix + name_info.Yyyy + s + name_info.Dd + s + name_info.Mm + name_info.Suffix + name_info.Extension	// yyyymmdd → yyyyddmm
		}
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
func GetFilePath(file string, sensor_type string, sensor_serial string, conf *Config) (string, error) {
	// Default save directory
	file_path := conf.Save_to
	if !conf.No_sanitize {
		file_path = SanitizeDir(conf.Save_to)
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

// Get File to save name according to option
// return file_name to use
func GetFileFullName(file string, sensor_type string, sensor_serial string, conf *Config) (string, string, error) {
	file_name := GetFileName(file, sensor_type, sensor_serial, conf)
	file_path, file_path_err := GetFilePath(file, sensor_type, sensor_serial, conf)
	if nil != file_path_err {
		return "", "", file_path_err
	}

	return file_path + file_name, file_path, file_path_err
}

// Get File content modified according to config
// take file_content as a string and transform it according to config
func GetFileContent(file_name string, file_content string, conf *Config) (string) {
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
	if (conf.File_save_CR || conf.File_save_CRLF) && (!strings.HasSuffix(file_name, ".jpg")) && (!strings.HasSuffix(file_name, ".bmp") ) {
		// Replace LF by CR or CRLF
		if conf.File_save_CRLF {
			file_content = strings.Replace(file_content, "\n", "\r\n", -1) // Happy Windows user
		} else if conf.File_save_CR {
			file_content = strings.Replace(file_content, "\n", "\r", -1) 	// Happy old OSX user
		}
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

// Get version int "1.3.X" or 1.3.0-rc (1515)" becomes 13
// Return a default version = 14 if there is error during string conversion
func GetSensorVersionInt(version_str string) (int) {
	version_int := 14 // Set a default version to 14
	version_str = strings.Replace(version_str, "X", "0", -1) // "X" = 0  -> version "1.3.X" = "1.3.0"
	version_str_splited := strings.Split(version_str, ".") // Version start always with 2 digits seperated by '.' "1.3.0-rc (~7327)"
	if len(version_str_splited) >= 3 { // Eurecam version always look like X.Y.Z, so it's not normal
		version_tobeint := "14"
		version_tobeint_tmp := version_str_splited[0] + "" + version_str_splited[1] //+ "" + version_str_splited[2]
		reg, reg_err := regexp.Compile("[^0-9]+")
		if nil == reg_err {
			version_tobeint = reg.ReplaceAllString(version_tobeint_tmp, "")
		}
		version_int_tmp, version_int_err := strconv.Atoi(version_tobeint)
		if nil == version_int_err {
			version_int = version_int_tmp
		}
	}

	return version_int
}

// Set fileType we are downloading according to config
func GetFileToGet(sensor_type string, sensor_version string, conf *Config) ([]FileType) {
	version_int := GetSensorVersionInt(sensor_version)
	files_to_get := []FileType{}

	// Get file common to Comptipix-2D and 3D + Concentrix and iComptipix
	if conf.Get_counting || conf.Get_all {
		counting_type := FileType{Prefix:"", Suffix:".csv"} // add counting file, ex: 20140525.csv
		files_to_get = append(files_to_get, counting_type)
	}
	if conf.Get_occupancy || conf.Get_all {
		occupancy_type := FileType{Prefix:"", Suffix:"_presence.csv"} // add occupancy file, ex: 20140525_presence.csv
		files_to_get = append(files_to_get, occupancy_type)
	}
	if conf.Get_log || conf.Get_all {
		log_type := FileType{Prefix:"log_", Suffix:".csv"} // add log file, ex: log_20140525.csv
		files_to_get = append(files_to_get, log_type)
	}

	if ("CPX3" == sensor_type) || ("SX3D" == sensor_type) {
		// Get file common to Comptipix-2D and 3D
		if conf.Get_config || conf.Get_all {
			config_type := FileType{Prefix:"", Suffix:".txt"} // add config file, ex: 20140525.txt
			files_to_get = append(files_to_get, config_type)
		}
		if conf.Get_jpg || conf.Get_all {
			jpg_type := FileType{Prefix:"", Suffix:".jpg"} // add jpg file, ex: 20140525.jpg
			files_to_get = append(files_to_get, jpg_type)
		}

		// Get file specific to Comptipix-2D
		if "CPX3" == sensor_type {
			if (conf.Get_all && (version_int >= 14)) || conf.Get_line { // only for version 1.4.0 or if user explicitly add get_line option
				line_type := FileType{Prefix:"", Suffix:"_line.csv"} // add line file, ex: 20140525_line.csv
				files_to_get = append(files_to_get, line_type)
			}
		}

		// Get file specific to Comptipix-3D
		if "SX3D" == sensor_type {
			if conf.Get_area || conf.Get_all {
				area_type := FileType{Prefix:"", Suffix:"_area.csv"} // add area file, ex: 20140525_area.csv
				files_to_get = append(files_to_get, area_type)
			}
			if conf.Get_track || conf.Get_all {
				track_type := FileType{Prefix:"", Suffix:"_track.csv"} // add track file, ex: 20140525_track.csv
				files_to_get = append(files_to_get, track_type)
			}
			if conf.Get_heat || conf.Get_all {
				heatmap_type := FileType{Prefix:"", Suffix:"_heat.csv"} // add heatmap file, ex: 20140525_heat.csv
				files_to_get = append(files_to_get, heatmap_type)
			}
			if conf.Get_depth || conf.Get_all {
				depthmap_type := FileType{Prefix:"", Suffix:"_depth.bmp"} // add depthmap file, ex: 20140525_depth.bmp
				files_to_get = append(files_to_get, depthmap_type)
			}
		}
	}

	return files_to_get
}

// Do a POST request (used for login only)
// return http status code, answer content and possible error
func doPOST(http_client *http.Client, addr string, post_data []byte, param string) (int, string, error) {
	http_code := 500
	http_answ := ""

	// prepare POST in octet-stream to get token
	req, req_err := http.NewRequest("POST", addr + param, bytes.NewBuffer(post_data))
	if req_err != nil {
		return http_code, http_answ, req_err
	}
	req.Header.Set("Content-Type", "application/octet-stream") // explicit content-type (octet-stream has avantage to not be re-arranged by any "intelligent" router/switch)

	// Send POST request
	resp, http_err := http_client.Do(req)
	if http_err != nil {
		return http_code, http_answ, http_err
	}
	defer resp.Body.Close()

	// read code
	http_code = resp.StatusCode

	// Read answer
	body, _ := ioutil.ReadAll(resp.Body)
	http_answ = string(body)
	resp.Body.Close()

	// Return answer
	return http_code, http_answ, http_err
}

// Do login
// Comptipix-V3 mode (get a token)
// return http status code, answer content and possible error
func DoLogin(http_client *http.Client, addr string, user string, pass string) (int, string, error) {
	post_data := "user="+ user +"&pass="+ pass
	return doPOST(http_client, addr, []byte(post_data), "CONFIG?get_tkn")
}

// Do logout
// return possible error
func DoLogout(http_client *http.Client, addr string, token string, old_mode bool) (error) {
	var logout_err error
	if !old_mode {
		_, logout_err = http_client.Get(addr + "logout.html?tkn=" + token) // always use token for Comptipix-V3 API
	} else {
		_, logout_err = http_client.Get(addr + "logout.html") // no token for old API
	}
	return logout_err
}

// Do an old login (no token)
// Compatibility mode for Concentrix and iComptipix
// return http status code, answer content and possible error
func DoLoginOld(http_client *http.Client, addr string, user string, pass string) (int, string, error) {
	// adapt user/pass if user provided reader/reader
	if "reader" == user {
		user = "user"
		if "reader" == pass {
			pass = "user"
		}
	}
	post_data := "user="+ user +"&pass="+ pass
	return doPOST(http_client, addr, []byte(post_data), "")
}

// Get type, serial and version
// return http status code (int), answer type (string), answer serial (string), answer version (string) and possible error (error)
func GetTypeSerial(http_client *http.Client, addr string, tkn string) (int, string, string, string, error) {
	http_code := 500
	answ_type := ""
	answ_serial := ""
	answ_version := ""

	// Send GET request
	resp, http_err := http_client.Get(addr + "CONFIG?tkn=" + tkn + "&type&serial&version")
	if http_err != nil {
		return http_code, answ_type, answ_serial, answ_version, http_err
	}
	defer resp.Body.Close()

	// Read http code
	http_code = resp.StatusCode

	// Extract Type + Serial
	if 200 == resp.StatusCode {
		// The answer should be like :
		// "type=CPX3
		// serial=119093
		// version=1.3.1"

		// Read line by line
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			// Read one line
			line := scanner.Text()

			// Extract type and serial
			if strings.HasPrefix(line, "type=") {
				answ_type = strings.TrimPrefix(line, "type=")
				answ_type = strings.TrimSuffix(answ_type, "\n")
			} else if strings.HasPrefix(line, "serial=") {
				answ_serial = strings.TrimPrefix(line, "serial=")
				answ_serial = strings.TrimSuffix(answ_serial, "\n")
			} else if strings.HasPrefix(line, "version=") {
				answ_version = strings.TrimPrefix(line, "version=")
				answ_version = strings.TrimSuffix(answ_version, "\n")
			}

			if ("" != answ_type) && ("" != answ_serial) && ("" != answ_version) {
				break // we got all wanted parameters, no need to read more
			}
		}
	}

	// Return answer
	resp.Body.Close()
	return http_code, answ_type, answ_serial, answ_version, http_err
}

// Get type, serial and version
// Compatibility mode for Concentrix and iComptipix
// return http status code (int), answer type (string), answer serial (string), answer version (string) and possible error (error)
func GetTypeSerialOld(http_client *http.Client, addr string) (int, string, string, string, error) {
	http_code := 500
	answ_type := ""
	answ_serial := ""
	answ_version := ""

	// Send GET request
	resp, http_err := http.Get(addr + "CONFIG?capteur_nb&serie&version")
	if http_err != nil {
		return http_code, answ_type, answ_serial, answ_version, http_err
	}
	defer resp.Body.Close()

	// Read http code
	http_code = resp.StatusCode

	// Extract Type + Serial
	if 200 == resp.StatusCode {
		// The answer should be like :
		// "capteur_nb=8
		// serie=119093
		// version=1.4.4"

		// Read line by line
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			// Read one line
			line := scanner.Text()

			// Extract type and serial
			if strings.HasPrefix(line, "capteur_nb=") {
				nb_sensor_str := strings.TrimPrefix(line, "capteur_nb=")
				nb_sensor_str = strings.TrimSuffix(nb_sensor_str, "\n")
				// Get sensor number
				nb_sensor, err_nb_sensor := strconv.Atoi(nb_sensor_str)
				if nil == err_nb_sensor {
					if nb_sensor > 1 {
						answ_type = "CTX" // more than one sensor is a concentrix (should be 8)
					} else {
						answ_type = "ICPX" // 1 sensor is an iComptipix
					}
				} else {
					http_err = err_nb_sensor // Set error returned
				}
			} else if strings.HasPrefix(line, "serie=") {
				answ_serial = strings.TrimPrefix(line, "serie=")
				answ_serial = strings.TrimSuffix(answ_serial, "\n")
			} else if strings.HasPrefix(line, "version=") {
				answ_version = strings.TrimPrefix(line, "version=")
				answ_version = strings.TrimSuffix(answ_version, "\n")
			}

			if ("" != answ_type) && ("" != answ_serial) && ("" != answ_version) {
				break // we got all wanted parameters, no need to read more
			}
		}
	}

	// Return answer
	resp.Body.Close()
	return http_code, answ_type, answ_serial, answ_version, http_err
}

// Get type and serial of sensor
// return : http answer code (int), token (string) and possible error (error)
func SensorGetTypeSerial(http_client *http.Client, addr string, token string, old_mode bool) (int, string, string, string, error) {
	sensor_type := ""
	sensor_serial := ""
	sensor_version := ""
	http_code := 0
	var http_err error

	if !old_mode {
		http_code, sensor_type, sensor_serial, sensor_version, http_err = GetTypeSerial(http_client, addr, token) // Comptipix-2D or Comptipix-3D API
	} else {
		http_code, sensor_type, sensor_serial, sensor_version, http_err = GetTypeSerialOld(http_client, addr)		// Old API : Concentrix or iComptpipix
	}
	return http_code, sensor_type, sensor_serial, sensor_version, http_err
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

// Login to a sensor
// return : http answer code (int), token (string), old_mode true if this is a old API (bool) and possible error (error)
func SensorLogin(http_client *http.Client, addr string, conf *Config) (int, string, bool, error) {
	old_mode := false // Flag to use Concentrix/iComptipix old protocol, not used by default → will be used if 'get_tkn' command fail

	// Start process to connect to sensor
	PrintAndLog(fmt.Sprintf("Connecting to : %s\n", addr), conf, false) // Verbose only

	// Try to get a token
	http_code, token, http_err := DoLogin(http_client, addr, conf.User, conf.Pass)
	if nil != http_err { // Check we have something
		PrintAndLog(fmt.Sprintf("Error no answer from this addr (%s) : ABORT !!!\n", addr), conf, true) // Verbose + log
		return http_code, token, old_mode, http_err
	}

	// Check answer
	if 200 != http_code {
		// Authentication error
		if (401 == http_code) || (403 == http_code) || (418 == http_code) {
			PrintAndLog(fmt.Sprintf("Error Getting token (%d) : login/pass (%s/%s) wrong\n", http_code, conf.User, conf.Pass), conf, false) // Verbose only
			if (418 == http_code) {
				PrintAndLog(fmt.Sprintf("ABORT !!\n"), conf, false) // 418 is login error in new API
				return http_code, token, old_mode, http_err
			}
		} else {
			PrintAndLog(fmt.Sprintf("Error Getting token (%d)\n", http_code), conf, true) // Verbose + log
		}

		// Re-Test with old mode
		PrintAndLog(fmt.Sprintf("try old auth ... \n"), conf, false)
		http_code, token, http_err = DoLoginOld(http_client, addr, conf.User, conf.Pass)
		if (200 == http_code)&&("erreur" != token) {
			old_mode = true
		} else {
			// If token is wrong no need to continue
			PrintAndLog(fmt.Sprintf("Error authentication (%d) : ABORT !!!\n", http_code), conf, true)
			return http_code, token, old_mode, http_err
		}
	} else if token == "erreur" { // Concentrix/iComptipix send status=200 + resp="erreur"
		// Re-Test with old mode
		http_code, token, http_err = DoLoginOld(http_client, addr, conf.User, conf.Pass)
		if (200 == http_code) && ("erreur" != token) {
			old_mode = true
			PrintAndLog(fmt.Sprintf("Authentication with old API OK (%d) %s\n", http_code, token), conf, false)
		} else {
			// If token is wrong no need to continue
			PrintAndLog(fmt.Sprintf("Error authentication with old API (%d / token:%s) : ABORT !!!\n", http_code, token), conf, true)
			return http_code, token, old_mode, http_err
		}
	}
	PrintAndLog(fmt.Sprintf("Token ok (%d) : %s \n", http_code, token), conf, false)

	return http_code, token, old_mode, http_err
}

// Read user date start and date stop and nb_file
// and compute date_start + nulber of file to get
// return date_start and number of file to get
func ReadDateStartStop(conf *Config) (time.Time, int) {
	now := time.Now() // use now by default, or if user date_start is invalid
	nb_file := conf.Nb_file

	// Read user date_start
	if conf.Date_start != "" {
		time_format_start := "2006-01-02"
		d_start, d_start_err := time.Parse(time_format_start, conf.Date_start)
		if nil == d_start_err {
			now = d_start
		} else {
			PrintAndLog(fmt.Sprintf("Invalid date_start (%s), Ex valid date : '2014-05-25' err: %v  → date_start ignored, use now (%s) instead \n\n", conf.Date_start, d_start_err, now.Local()), conf, true) // Print error day
		}
	}

	// Read user date_stop and recompute nb_file
	if conf.Date_stop != "" {
		time_format_stop := "2006-01-02"
		d_stop, d_stop_err := time.Parse(time_format_stop, conf.Date_stop)
		if nil == d_stop_err {
			tmp_nb := int(now.Sub(d_stop).Hours()/24)
			if tmp_nb >= 0 {
				nb_file = tmp_nb
			} else {
				// Invert date_stop and now (or date_start), so user always get files
				nb_file = int(d_stop.Sub(now).Hours()/24)
				now = d_stop
			}
		} else {
			PrintAndLog(fmt.Sprintf("Invalid date_stop (%s), Ex valid date : '2014-05-25' err: %v  → date_stop ignored \n\n", conf.Date_stop, d_stop_err), conf, true) // Print error day
		}
	}

	return now, nb_file
}

// Get a string date format Eurecam (like "20171122") from a date type
func GetDateStringEurecam(date_current time.Time) (string) {
	t := date_current.Format("20060102-150405") // Format date so we can easily remove hour min seconds
	tt := strings.Split(t, "-") // get a string like "20140525"

	return tt[0]
}

// Parse and get sensor answer file size
// return :
// skip_file : true to skip this file,
// check_size : true to check next file size, false otherwise
// second_wait : number of second to wait before continue
// file_size : size of file reported by info command
func GetSensorFileInfo(file_info string, file_name_sensor string, old_mode bool, conf *Config) (bool, bool, int, int64) {
	skip_file := false
	check_size := true
	second_wait := 0
	var file_size int64 = 0
	if !old_mode {	// Comptipix-V3 answer
		if strings.HasPrefix(file_info, "sdcard_info=-1,,") {
			PrintAndLog(fmt.Sprintf("File %s ERROR : There is no file with this name, skip file\n\n", file_name_sensor), conf, false)
			skip_file = true // try next file
		} else if strings.HasPrefix(file_info, "sdcard_info=busy,,") || strings.HasPrefix(file_info, "error") {
			PrintAndLog(fmt.Sprintf("File %s : SD card is busy ... wait 5 seconds\n", file_name_sensor), conf, false)
			check_size = false // bypass check size in this case
			second_wait = 5
		} else {
			PrintAndLog(fmt.Sprintf("File %s, info OK : %s", file_name_sensor, strings.TrimPrefix(file_info, "sdcard_info=")), conf, false)
		}
	} else { // Old Concentrix/iComptipix answer
		if strings.HasPrefix(file_info, file_name_sensor+"=") {
			tmp_old_infos := strings.Split(file_info, "=")
			tmp_old_size, _ := strconv.Atoi(tmp_old_infos[1])
			if -1 == tmp_old_size {
				PrintAndLog(fmt.Sprintf("File %s ERROR : There is no file with this name\n", file_name_sensor), conf, false)
				skip_file = true // try next file
			} else {
				PrintAndLog(fmt.Sprintf("File %s, info OK : %s\n", file_name_sensor, strings.TrimPrefix(file_info, file_name_sensor+"=")), conf, false)
			}
		} else {
			PrintAndLog(fmt.Sprintf("File %s : SD card is in error ... wait 5 seconds\n", file_name_sensor), conf, false)
			check_size = false // bypass check size in this case
			second_wait = 5
		}
	}

	// Get SDcard file size
	if !skip_file && check_size {
		tmp_infos := strings.TrimPrefix(file_info, "sdcard_info=")
		fi_infos := strings.Split(tmp_infos, ",")
		fi_info_index := 0
		if old_mode {
			fi_infos = strings.Split(file_info, "=") // Old mode answer is like '20140525.csv=65456' and if file not exist '20140525.csv=-1'
			fi_info_index = 1
		}
		fi_size_sd, fi_size_size_err := strconv.ParseInt(fi_infos[fi_info_index], 10, 64)
		if nil == fi_size_size_err {
			file_size = fi_size_sd
		} else {
			PrintAndLog(fmt.Sprintf("ERROR %s : converting file_size from file_info = %s\n", fi_size_size_err.Error(), file_info), conf, false)
			check_size = false // bypass check size in this case
		}
	}

	return skip_file, check_size, second_wait, file_size
}

// Get file from sensor according to config
// return number of file succefully fetched (int), number of file error (int)
func GetAllFiles(conf *Config) (int, int) {
	// Set net client
	var http_client = &http.Client{
		Timeout: time.Second * time.Duration(conf.Default_timeout), //NOTE: It's very important to set a default timeout that fit your application need (Golang default is no Timeout, we discourage to keep the default in production)
	}

	// Process sensor info and user config
	// ---

	// Sanitize user entry
	addr := conf.Addr
	if !conf.No_sanitize {
		addr = SanitizeURL(conf.Addr)
	}

	// Login to sensor
	http_login_code, token, old_mode, http_login_err := SensorLogin(http_client, addr, conf)
	if (200 != http_login_code) || (nil != http_login_err) {
		if (nil != http_login_err) {
			PrintAndLog(fmt.Sprintf("Login ERROR: %s - code: %d \n", http_login_err.Error(), http_login_code), conf, false)
		}
		return 0, 0 // Login error
	}

	// Get sensor we are connecting to type and serial info
	http_serial_code, sensor_type, sensor_serial, sensor_version, http_serial_err := SensorGetTypeSerial(http_client, addr, token, old_mode)
	if (200 != http_serial_code) || (nil != http_serial_err) {
		if (nil != http_serial_err) {
			PrintAndLog(fmt.Sprintf("Basic info ERROR: %s - code: %d \n", http_serial_err.Error(), http_login_code), conf, false)
		}
		return 0, 0 // Can't get basic info from sensor
	}

	// Set request to get file (according to old or new API)
	req_get_info := addr+"CONFIG?tkn="+token+"&sdcard_info="
	req_get_file := addr+"CONFIG?tkn="+token+"&sdcard_read="
	if old_mode {
		req_get_info = addr+"FICHIER?info="		// old request for Concentrix/iComptipix
		req_get_file = addr+"FICHIER?lecture="
	}

	// Read user file to get + start / stop date
	files_to_get := GetFileToGet(sensor_type, sensor_version, conf)	// Files Type to get
	date_start, nb_file := ReadDateStartStop(conf)							// Read user date_start and date_stop and nb_file


	// Loop through Date to get files
	// ---

	// Get files, var to store infos
	var resp *http.Response
	var err error
	var body []byte
	var body_err error
	c := 1 		// files counter
	c_err := 0 	// files error counter
	c_ok := 0	// files ok counter
	flag_stop_on_same_size := false // flag to stop downloading files when there is alredy same size present

	PrintAndLog(fmt.Sprintf("Will get %d day(s) from : %d-%02d-%02d \n", nb_file, date_start.Year(), date_start.Month(), date_start.Day()), conf, false) // Print how many files to get
	for c <= nb_file {
		c++ // increment file counter

		// Check if we have to stop
		if flag_stop_on_same_size {
			PrintAndLog(fmt.Sprintf("↳ Stopping on size match because 'stop_size' is used\n"), conf, false)
			break
		}

		// Loop through files types
		for i := range(files_to_get) {
			// Get file_name_sensor + file_name_disk to save
			file_name_sensor := files_to_get[i].Prefix + GetDateStringEurecam(date_start) + files_to_get[i].Suffix	// File name to get from sensor
			file_name_disk, file_path_disk, file_name_disk_err := GetFileFullName(file_name_sensor, sensor_type, sensor_serial, conf)// File name full path to save to disk according to config
			if nil != file_name_disk_err {
				PrintAndLog(fmt.Sprintf("ERROR file name: %s (%s)  → Stop here\n", file_name_sensor, file_name_disk_err.Error()), conf, false)
				return c_ok, c_err
			}

			// Get file Infos
			// ---

			// Check file exist
			resp, err = http_client.Get(req_get_info + file_name_sensor)
			if err != nil {
				PrintAndLog(fmt.Sprintf("ERROR getting file: %s  → Skip file\n", file_name_sensor), conf, false)
				continue // Skip file not exist
			}
			defer resp.Body.Close()

			// Init check size
			check_size := true
			if conf.No_size_check { // Set a temp variable as we may bypass check size if SDcard has error
				check_size = false
			}

			// Read file_info answer code
			if 200 != resp.StatusCode {
				PrintAndLog(fmt.Sprintf("Error reading file info %s, error code: %s", file_name_sensor, resp.Status), conf, false)
				if 401 == resp.StatusCode {	// 401 = we are not loged anymore
					PrintAndLog(fmt.Sprintf("File %s → Error code: 401 ... need to relogin or login/pass is now wrong ABORT \n", file_name_sensor), conf, false)
					c_err++
					return c, c_err //TODO: relogin
				} else if 503 == resp.StatusCode { 	// 503 means server is busy, stop during 5 seconds
					PrintAndLog(fmt.Sprintf("File %s → Server is busy (%s) ... wait 5 seconds and bypass check size\n", file_name_sensor, resp.Status), conf, false) // NOTE: we don't retry here as file info is not realy important, we will retry if there is a 503 on getting file
					check_size = false // Bypass check size in this case
					time.Sleep(5 * time.Second)
				} else {										// If file not ok try the next one
					PrintAndLog(fmt.Sprintf("File %s → ERROR (%s), skip file", file_name_sensor, resp.Status), conf, false)
					continue
				}
			}

			// Read file info
			body, body_err = ioutil.ReadAll(resp.Body)
			if nil != body_err {
				PrintAndLog(fmt.Sprintf("File %s → read info ERROR (%s), skip file\n", file_name_sensor, body_err.Error()), conf, false)
				continue
			}
			file_info := string(body)
			resp.Body.Close()

			// Check sensor's file info
			skip_file, check_size, second_wait, file_size := GetSensorFileInfo(file_info, file_name_sensor, old_mode, conf)
			if skip_file {
				continue 	// File doesn't exist
			} else if second_wait > 0 {
				time.Sleep(time.Duration(second_wait) * time.Second)	// Wait a little because SDcard is busy
			}

			// Check if files allready present with same size
			if !conf.No_size_check && check_size {
				c_ok++
				if file_size == GetFileSize(file_name_disk) { // File on SDcard and Disk has same size
					PrintAndLog(fmt.Sprintf("File %s → same size on disk (%d)\n", file_name_sensor, file_size), conf, false)
					if conf.Stop_size {
						flag_stop_on_same_size = true
					}
					continue // Skip and go to next file
				}
			}

			// Get file
			// ---

			// Get file
			resp, err = http_client.Get(req_get_file + file_name_sensor)
			if nil != err {
				PrintAndLog(fmt.Sprintf("ERROR getting file: %s  → Skip file\n", file_name_sensor), conf, false)
				continue
			}
			defer resp.Body.Close()

			// Read get file answer code
			if 200 != resp.StatusCode {
				PrintAndLog(fmt.Sprintf("Error reading file %s, error code: %s", file_name_sensor, resp.Status), conf, false)
				if 404 == resp.StatusCode {			// 404 = no file
					continue // try next file
				} else if 503 == resp.StatusCode {	// 503 means server is busy, stop during 5 seconds and retry
					PrintAndLog(fmt.Sprintf("Server is busy ... wait 5 seconds\n"), conf, false)
					time.Sleep(5 * time.Second)
					i = -1 // Start again
					continue
				} else {
					PrintAndLog(fmt.Sprintf("ERROR getting file %d !!! \n", resp.StatusCode), conf, false)
					continue // try next file
				}
			}

			// Read file
			body, body_err = ioutil.ReadAll(resp.Body)
			if nil != body_err {
				PrintAndLog(fmt.Sprintf("File %s → read file ERROR (%s), skip file\n", file_name_sensor, body_err.Error()), conf, false)
				continue
			}
			resp.Body.Close()
			PrintAndLog(fmt.Sprintf("↳ Save File %s as %s\n", file_name_sensor, file_name_disk), conf, false)

			// Save file
			os.MkdirAll(file_path_disk, 0777) // always check and create directory before write
			f, f_err := os.Create(file_name_disk)
			if nil != f_err {
				PrintAndLog(fmt.Sprintf("ERROR saving file: %s ERROR: %s  → Skip file\n", file_name_disk, f_err.Error()), conf, false)
				continue // try next file
			}
			defer f.Close()

			// Check replace comma before write or CRLF replacement
			if (conf.File_save_CR || conf.File_save_CRLF || conf.Replace_comma_by != "") && (!strings.HasSuffix(file_name_sensor, ".jpg")) && (!strings.HasSuffix(file_name_sensor, ".bmp")) { // Check file must be transformed and file is txt file type
				file_content := GetFileContent(file_name_sensor, string(body[:]), conf)
				f.Write([]byte(file_content))	// Write transformed file
			} else {
				f.Write(body)						// Write Vanilla file
			}
			f.Close()								// Close file
			c_ok++

			// Wait some seconds (we don't want SDcard to produce error)
			time.Sleep(time.Duration(conf.Wait_timer) * time.Second)
			PrintAndLog(fmt.Sprintf("---\n"), conf, false)
		}

		// Compute new date to get
		date_start = date_start.AddDate(0,0,-1) // date_start -1 day

		// Wait more second every 2 group of file (because we don't want SDcard to produce error)
		if c%2 == 0 {
			time.Sleep(time.Duration(conf.Wait_timer) * time.Second)
		}
	}

	return c_ok, c_err
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
				conf.Save_to = filepath.ToSlash(conf.Save_to) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
				return true
			} else {
				fmt.Printf("BAD JSON config file: %s\n", json_err.Error())
			}
		}
	}
	return false
}


// Main program
// --------------------------

// Read cmd line
// Get all files
func main() {

	default_json_config_path := "./sdcard_fetcher.json"

	// Help arg (to keep an old compatibility argument wrapper)
	// ----------------

	help_addr := "Comptipix-V3 address (Default: 'http://192.168.0.100/')"
	help_verbose := "To be verbose (Default: no verbose)"
	help_nb_file := "Number of file to get since today (Default: 2)"
	help_pass := "User password (Default: 'reader')"
	help_user := "User type (Default: 'reader')"
	help_save_to := "Directory to save files (Default: inside execution path)"
	help_get_counting := "Get counting file (ex: '20170822.csv')"
	help_get_occupancy := "Get occupancy file (ex: '20170822_presence.csv')"
	help_get_log := "Get log file (ex: 'log_20170822.csv')"
	help_get_jpg := "Get image file (ex: '20170822.jpg')"
	help_get_config := "Get config file (ex: '20170822.txt')"
	help_get_line := "Get line file (ex: '20170822_line.csv') (Comptipix-2D only)"
	help_get_area := "Get area	file (ex: '20170822_area.csv') (Comptipix-3D only)"
	help_get_track := "Get track file (ex: '20170822_track.csv') (Comptipix-3D only)"
	help_get_heat := "Get heatmap file (ex: '20170822_heat.csv') (Comptipix-3D only)"
	help_get_depth := "Get depthmap file (ex: '20170822_depth.bmp') (Comptipix-3D only)"
	help_get_all := "Get all files types (Default: Get all files, according to sensor type)"

	help_demo := "This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE"


	// Read arg
	// ----------------

	Addr := flag.String("addr", "http://192.168.0.100/", help_addr)
	Verbose := flag.Bool("verbose", false, help_verbose)
	Save_to := flag.String("save_to", "./Data/", help_save_to)
	Nb_file := flag.Int("nb_file", 2, help_nb_file)
	Pass := flag.String("pass", "reader", help_pass)
	User := flag.String("user", "reader", help_user)
	Get_counting := flag.Bool("get_counting", false, help_get_counting)
	Get_occupancy := flag.Bool("get_occupancy", false, help_get_occupancy)
	Get_log := flag.Bool("get_log", false, help_get_log)
	Get_jpg := flag.Bool("get_jpg", false, help_get_jpg)
	Get_config := flag.Bool("get_config", false, help_get_config)
	Get_line := flag.Bool("get_line", false, help_get_line)
	Get_area := flag.Bool("get_area", false, help_get_area)
	Get_track := flag.Bool("get_track", false, help_get_track)
	Get_heat := flag.Bool("get_heat", false, help_get_heat)
	Get_depth := flag.Bool("get_depth", false, help_get_depth)
	Get_all := flag.Bool("get_all", false, help_get_all)
	Date_start := flag.String("date_start", "", "Date to start getting files (Default: now)")
	Date_stop := flag.String("date_stop", "", "Date to stop getting files (Default: not used)")
	Stop_size := flag.Bool("stop_size", false, "Stop downloading other files when there is same size files (Default: don't stop)")
	Log_suffix := flag.Bool("log_suffix", false, "Transform log file to have 'log_' as suffix and no more as prefix (Default: don't transform log name)")
	Yyyymmdd_to_yymmdd := flag.Bool("yyyymmdd_to_yymmdd", false, "Replace yyyymmdd by yymmdd : '20171203' becomes '171203' (Default: don't replace)")
	Yyyymmdd_to_yyddmm := flag.Bool("yyyymmdd_to_yyddmm", false, "Replace yyyymmdd by yyddmm : '20171203' becomes '170312' (Default: don't replace)")
	Yyyymmdd_to_mmddyy := flag.Bool("yyyymmdd_to_mmddyy", false, "Replace yyyymmdd by mmddyy : '20171203' becomes '120317' (Default: don't replace)")
	Yyyymmdd_to_mmyydd := flag.Bool("yyyymmdd_to_mmyydd", false, "Replace yyyymmdd by mmyydd : '20171203' becomes '121703' (Default: don't replace)")
	Yyyymmdd_to_ddmmyy := flag.Bool("yyyymmdd_to_ddmmyy", false, "Replace yyyymmdd by ddmmyy : '20171203' becomes '031217' (Default: don't replace)")
	Yyyymmdd_to_ddyymm := flag.Bool("yyyymmdd_to_ddyymm", false, "Replace yyyymmdd by ddyymm : '20171203' becomes '031712' (Default: don't replace)")
	Yyyymmdd_to_ddmmyyyy := flag.Bool("yyyymmdd_to_ddmmyyyy", false, "Replace yyyymmdd by ddyyyymm : '20171203' becomes '03201712' (Default: don't replace)")
	Yyyymmdd_to_ddyyyymm := flag.Bool("yyyymmdd_to_ddyyyymm", false, "Replace yyyymmdd by ddyyyymm : '20171203' becomes '03201712' (Default: don't replace)")
	Yyyymmdd_to_mmddyyyy := flag.Bool("yyyymmdd_to_mmddyyyy", false, "Replace yyyymmdd by mmddyyyy : '20171203' becomes '12032017' (Default: don't replace)")
	Yyyymmdd_to_mmyyyydd := flag.Bool("yyyymmdd_to_mmyyyydd", false, "Replace yyyymmdd by mmyyyydd : '20171203' becomes '12201703' (Default: don't replace)")
	Yyyymmdd_to_yyyyddmm := flag.Bool("yyyymmdd_to_yyyyddmm", false, "Replace yyyymmdd by yyyyddmm : '20171203' becomes '20170312' (Default: don't replace)")
	Yyyymmdd_separator := flag.String("yyyymmdd_separator", "", "Add a separator to date, if separator is '-' : '20171203.csv' becomes '2017-12-03.csv' (Default: no separator)")
	File_save_CR := flag.Bool("file_save_CR", false, "Transform default Comptipix 'LF' carriage return to 'CR', maybe usefull for old OSX program (Default: don't transform)")
	File_save_CRLF := flag.Bool("file_save_CRLF", false, "Transform default Comptipix 'LF' carriage return to 'CRLF', maybe usefull for old Windows program (Default: don't transform) - WARNING: in this mode all file will be redownloaded, because it's no longer same size (CRLF > LF)")
	Replace_csv_by := flag.String("replace_csv_by", "", "Replace csv extension by specified extension (Default: don't replace csv extension)")
	Replace_txt_by := flag.String("replace_txt_by", "", "Replace txt extension by specified extension (Default: don't replace txt extension)")
	Replace_comma_by := flag.String("replace_comma_by", "", "Replace comma inside file by specified char (Default: don't replace)")
	No_replace_comma_replace := flag.Bool("no_replace_comma_replace", false, "Don't replace the replacer char by comma if there is already a replacer char in file (Default: if replace_comma_by enabled → replace the replacer by comma)")
	Serial_prefix := flag.Bool("serial_prefix", false, "Add type + serial number as file name prefix for each files saved : '20170712.csv' becomes 'CPX3-117007_20170712.csv' for Comptipix-V3 SN 117007 (Default: no prefix)")
	Serial_suffix := flag.Bool("serial_suffix", false, "Add type + serial number as file name suffix for each files saved : '20170712_presence.csv' becomes '20170712_CPX3-117007_presence.csv' for Comptipix-V3 SN 117007 (Default: no suffix)")
	Serial_suffix_final := flag.Bool("serial_suffix_final", false, "Add type + serial number as file name suffix for each files saved : '20170712_presence.csv' becomes '20170712_presence_CPX3-117007.csv' for Comptipix-V3 SN 117007 (Default: no suffix final)")
	Save_dir_yyyy := flag.Bool("save_dir_yyyy", false, "Save file in Year (4 digits) of file directory like '2017/' for file '20171207.csv', can be used with save_dir_mm and save_dir_dd (Default: not used)")
	Save_dir_yy := flag.Bool("save_dir_yy", false, "Save file in Year (2 digits) of file directory like '17/' for file '20171207.csv', can be used with save_dir_mm and save_dir_dd (Default: not used)")
	Save_dir_mm := flag.Bool("save_dir_mm", false, "Save file in month of file directory like '12/' for file '20171207.csv', can be used with save_dir_dd and save_dir_yyyy (Default: not used)")
	Save_dir_dd := flag.Bool("save_dir_dd", false, "Save file in day of file directory like '07/' for file '20171207.csv', can be used with save_dir_mm and save_dir_yyyy (Default: not used)")
	Save_dir_sep := flag.String("save_dir_sep", "/", "Save Directory separator : '/' to create sub directory like yyyy/mm/dd (Default:'/')")
	No_serial_dir := flag.Bool("no_serial_dir", false, "Don't save file in type+serial sub-directory like 'CPX3-119063' for Comptipix-V3 with serial number 119063 (Default:  sub-directory enabled)")
	No_size_check := flag.Bool("no_size_check", false, "Don't check if there is already file with same name and size present (Default: size check enabled)")
	No_sanitize := flag.Bool("no_sanitize", false, "Don't Sanitize user entry (Default: sanitize enabled) : '-a 192.168.0.100' → '-a http://192.168.0.100/' and '-s my_dir' → '-s my_dir/'")
	No_logout := flag.Bool("no_logout", false, "Don't send logout when download finish (Default: logout enabled)")
	No_log := flag.Bool("no_log", false, "Don't log error in a file log.log (Default: log enabled)")
	Default_timeout := flag.Int("default_timeout", 40, "Connection timeout in seconds (Default: 40)")
	Wait_timer := flag.Int("wait_timer", 1, "Time in second to wait between 2 files downloading : needed to not bother sdcard too much (Default: 1s)")
	// Read flags
	Export_this_config_json := flag.Bool("export_this_config_json", false, "Export the current config as sdcard_fetcher.json.new (Default don't export)")
	Config_json := flag.Bool("config_json", false, "Set this flag to read config from sdcard_fetcher.json file (Default no read from sdcard_fetcher.json)")
	Config_json_path := flag.String("config_json_path", default_json_config_path, "full path to file config (ex:'my/super/path/to/my_config.json')")
	flag_help := flag.Bool("help", false, "Print demo usage") // Override default help to add tips and usage

	// Old compatibility flags
	Old_addr := flag.String("a", "http://192.168.0.100/", help_addr)
	Old_verbose := flag.Bool("v", false, help_verbose)
	Old_nb_file := flag.Int("n", 2, help_nb_file)
	Old_pass := flag.String("p", "reader", help_pass)
	Old_user := flag.String("u", "reader", help_user)
	Old_get_counting := flag.Bool("fcounting", false, help_get_counting)
	Old_get_occupancy := flag.Bool("foccupancy", false, help_get_occupancy)
	Old_get_log := flag.Bool("flog", false, help_get_log)
	Old_get_jpg := flag.Bool("fjpg", false, help_get_jpg)
	Old_get_config := flag.Bool("fconfig", false, help_get_config)
	Old_get_line := flag.Bool("fline", false, help_get_line)
	Old_get_area := flag.Bool("farea", false, help_get_area)
	Old_get_track := flag.Bool("ftrack", false, help_get_track)
	Old_get_heat := flag.Bool("fheat", false, help_get_heat)
	Old_get_depth := flag.Bool("fdepth", false, help_get_depth)
	Old_get_all := flag.Bool("fall", false, help_get_all)
	Old_save_to := flag.String("s", "", help_save_to)

	// Parse all flags
	flag.Parse()

	// Print help if user asked for it
	if *flag_help != false {
		fmt.Printf("\nOPTIONS:\n")
		fmt.Printf("----------\n")
		fmt.Printf("see -h\n\n")
		fmt.Printf("\nEXAMPLES:\n")
		fmt.Printf("----------\n")
		fmt.Printf("Get 42 last counting files from sensor in 192.168.0.141 :\n")
		fmt.Printf(" ./sdcard_fetcher -n=42 -a=http://192.168.0.141/ -get_counting -v\n\n")
		fmt.Printf("Get 2 counting + log + jpg files from sensor in 192.168.0.141, saving data in another directory and using password for reader level 'my_pass' :\n")
		fmt.Printf(" ./sdcard_fetcher -a=http://192.168.0.141/ -p=my_pass -get_counting -get_log -get_jpg -save_to=directory/to/save/data/ -v\n\n")
		fmt.Printf("Get all files (counting, occupancy, log, jpg and config) from 2017-01-22 to 2014-05-24 (=974 days), from sensor in 192.168.0.141 :\n")
		fmt.Printf(" ./sdcard_fetcher -n=974 -a=http://192.168.0.141/ -date_start=2017-01-22 -v\n\n")
		fmt.Printf("Same as preceding using date_start and date_stop :\n")
		fmt.Printf(" ./sdcard_fetcher -a=http://192.168.0.141/ -date_start=2017-01-22 -date_stop=2014-05-24 -v\n\n")
		fmt.Printf("Get 2 counting from sensor in 192.168.0.141, replace comma in file by ';' and convert files in format 'mm-dd-yy' :\n")
		fmt.Printf(" ./sdcard_fetcher -a=http://192.168.0.141/ -fcounting -replace_comma_by ';' -yyyymmdd_to_mmddyy -yyyymmdd_separator '-' -v\n\n")
		fmt.Printf("Get 42 last counting files from sensor in 192.168.0.141, and save it in 'my/data/yyyy/mm' directory without Type-Serial sub_dir and adding Type-Serial as file name suffix + set prefix 'log_' to becomes a file suffix '_log' :\n")
		fmt.Printf(" ./sdcard_fetcher -n=42 -a=http://192.168.0.141/ -get_counting -no_serial_dir -serial_suffix -save_to 'my/data/dir' -save_dir_yyyy -save_dir_mm -log_suffix -v\n\n")
		fmt.Printf("\nLIMITATIONS:\n")
		fmt.Printf("----------\n")
		fmt.Printf("* %s\n", help_demo)
		fmt.Printf("* This program get files saved on SD card (from Comptipix-V3(2D) or Comptipix-3D or iComptipix or Concentrix), by default check files sizes differences before downloading new file \n")
		fmt.Printf("  ↳ So it can NOT work with a TinyCount or an Affix ! Because they haven't any SD card nor memory saving files ! \n")
		return // if we print help, we do nothing else
	}
	fmt.Printf("Note: %s\n", help_demo) // Print demo notice at each run

	// if nothing set by user --> enable -get_all flag
	if !(*Get_counting) && !(*Get_occupancy) && !(*Get_log) && !(*Get_jpg) && !(*Get_config) && !(*Get_line) && !(*Get_area) && !(*Get_track) && !(*Get_heat) && !(*Get_depth) {
		if !(*Old_get_counting) && !(*Old_get_occupancy) && !(*Old_get_log) && !(*Old_get_jpg) && !(*Old_get_config) && !(*Old_get_line) && !(*Old_get_area) && !(*Old_get_track) && !(*Old_get_heat) && !(*Old_get_depth) {
			*Get_all = true
		}
	}


	// read config from JSON
	var conf Config
	if (true == *Config_json) || (*Config_json_path != default_json_config_path) {
		// Read config file
		read_json_ok := true
		_, fi_err := os.Stat(*Config_json_path)
		fmt.Printf("\n* Read json file %s :\n", *Config_json_path)
		if nil == fi_err {
			// Read file
			f, f_err := ioutil.ReadFile(*Config_json_path)
			if nil == f_err {
				// Decode JSON
				json_err := json.Unmarshal(f, &conf)
				if nil != json_err {
					fmt.Printf("BAD JSON config file: %s,  STOP execution \n", json_err.Error())
					read_json_ok = false
				}
			}
		}
		if !read_json_ok {
			fmt.Printf("Error reading json file ABORT\n")
			return
		}
	} else {
		// Because we have to do it after flag.Parse()

		// Parse old argument compatibility style
		if "http://192.168.0.100/" != *Old_addr {
			conf.Addr = *Old_addr
		} else {
			conf.Addr = *Addr
		}
		if false != *Old_verbose {
			conf.Verbose = *Old_verbose
		} else {
			conf.Verbose = *Verbose
		}
		if 2 != *Old_nb_file {
			conf.Nb_file = *Old_nb_file
		} else {
			conf.Nb_file = *Nb_file
		}
		if "reader" != *Old_pass {
			conf.Pass = *Old_pass
		} else {
			conf.Pass = *Pass
		}
		if "reader" != *Old_user {
			conf.User = *Old_user
		} else {
			conf.User = *User
		}
		if false != *Old_get_counting {
			conf.Get_counting = *Old_get_counting
		} else {
			conf.Get_counting = *Get_counting
		}
		if false != *Old_get_occupancy {
			conf.Get_occupancy = *Old_get_occupancy
		} else {
			conf.Get_occupancy = *Get_occupancy
		}
		if false != *Old_get_log {
			conf.Get_log = *Old_get_log
		} else {
			conf.Get_log = *Get_log
		}
		if false != *Old_get_jpg {
			conf.Get_jpg = *Old_get_jpg
		} else {
			conf.Get_jpg = *Get_jpg
		}
		if false != *Old_get_config {
			conf.Get_config = *Old_get_config
		} else {
			conf.Get_config = *Get_config
		}
		if false != *Old_get_line {
			conf.Get_line = *Old_get_line
		} else {
			conf.Get_line = *Get_line
		}
		if false != *Old_get_area {
			conf.Get_area = *Old_get_area
		} else {
			conf.Get_area = *Get_area
		}
		if false != *Old_get_track {
			conf.Get_track = *Old_get_track
		} else {
			conf.Get_track = *Get_track
		}
		if false != *Old_get_heat {
			conf.Get_heat = *Old_get_heat
		} else {
			conf.Get_heat = *Get_heat
		}
		if false != *Old_get_depth {
			conf.Get_depth = *Old_get_depth
		} else {
			conf.Get_depth = *Get_depth
		}
		if false != *Old_get_all {
			conf.Get_all = *Old_get_all
		} else {
			conf.Get_all = *Get_all
		}
		if "" != *Old_save_to { // The old "s" option was empty
			conf.Save_to = *Old_save_to
		} else {
			conf.Save_to = *Save_to
		}
		conf.Stop_size = *Stop_size
		conf.Log_suffix = *Log_suffix
		conf.Date_start = *Date_start
		conf.Date_stop = *Date_stop
		conf.Yyyymmdd_to_yymmdd = *Yyyymmdd_to_yymmdd
		conf.Yyyymmdd_to_yyddmm = *Yyyymmdd_to_yyddmm
		conf.Yyyymmdd_to_mmddyy = *Yyyymmdd_to_mmddyy
		conf.Yyyymmdd_to_mmyydd = *Yyyymmdd_to_mmyydd
		conf.Yyyymmdd_to_ddmmyy = *Yyyymmdd_to_ddmmyy
		conf.Yyyymmdd_to_ddyymm = *Yyyymmdd_to_ddyymm
		conf.Yyyymmdd_to_ddmmyyyy = *Yyyymmdd_to_ddmmyyyy
		conf.Yyyymmdd_to_ddyyyymm = *Yyyymmdd_to_ddyyyymm
		conf.Yyyymmdd_to_mmddyyyy = *Yyyymmdd_to_mmddyyyy
		conf.Yyyymmdd_to_mmyyyydd = *Yyyymmdd_to_mmyyyydd
		conf.Yyyymmdd_to_yyyyddmm = *Yyyymmdd_to_yyyyddmm
		conf.Yyyymmdd_separator = *Yyyymmdd_separator
		conf.File_save_CR = *File_save_CR
		conf.File_save_CRLF = *File_save_CRLF
		conf.Replace_csv_by = *Replace_csv_by
		conf.Replace_txt_by = *Replace_txt_by
		conf.Replace_comma_by = *Replace_comma_by
		conf.No_replace_comma_replace = *No_replace_comma_replace
		conf.Serial_prefix = *Serial_prefix
		conf.Serial_suffix = *Serial_suffix
		conf.Serial_suffix_final = *Serial_suffix_final
		conf.Save_dir_yyyy = *Save_dir_yyyy
		conf.Save_dir_yy = *Save_dir_yy
		conf.Save_dir_mm = *Save_dir_mm
		conf.Save_dir_dd = *Save_dir_dd
		conf.Save_dir_sep = *Save_dir_sep
		conf.No_serial_dir = *No_serial_dir
		conf.No_size_check = *No_size_check
		conf.No_sanitize = *No_sanitize
		conf.No_logout = *No_logout
		conf.No_log = *No_log
		conf.Default_timeout = *Default_timeout
		conf.Wait_timer = *Wait_timer
	}
	// Display config
	displayConfig(&conf)

	// Export current config
	if *Export_this_config_json {
		exportConfigJSON(&conf, "sdcard_fetcher.json.new")
	}

	// Set log file
	if !conf.No_log {
		f, f_err := os.OpenFile("log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if f_err != nil {
			fmt.Printf("Error opening log file: %v !!!\n", f_err)
		} else {
			defer f.Close()
			log.SetOutput(f)
			PrintAndLog(fmt.Sprintf("SDcard_fetcher started at %s → on sensor %s\n---\n", time.Now().Format("2006-01-02 15:04:05"), conf.Addr), &conf, false)
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

			// Re-get file
			PrintAndLog(fmt.Sprintf("Re-Started because periodicity = %d minutes\n", conf.Periodic_minute), &conf, false)// Verbose only
			mfiles_ok, mfiles_err := GetAllFiles(&conf)
			PrintAndLog(fmt.Sprintf("\nDone at %s : Sensor %s → File OK: %d / error: %d \n", time.Now().Format("2006-01-02 15:04:05"), conf.Addr, mfiles_ok, mfiles_err), &conf, true)
			time.Sleep(time.Duration(conf.Periodic_minute) * time.Minute)
		}
	}

	// One shoot mode : Get files
	files_ok, files_err := GetAllFiles(&conf)
	PrintAndLog(fmt.Sprintf("\nDone: Sensor %s → File OK: %d / error: %d \n", conf.Addr, files_ok, files_err), &conf, true)
}
