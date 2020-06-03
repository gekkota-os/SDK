// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o sumer.exe sumer.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o sumer.exe sumer.go
//
// * RUN :
// ./sumer -h
//

package main

import (
	"log"
	"strings"
	"strconv"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"time"
	"path/filepath"
	"bytes"
	"flag"
)

// Header parse status (because eurecam header is 2 line : first the header, 2nd must be "fichier de comptage v2" or header is not valid)
type SiteChainHeader int32
const (
	SITE_CHAIN_IDLE   			SiteChainHeader = iota
	SITE_CHAIN_IN_PORGRESS
	SITE_CHAIN_DONE
)

// Write rule from to date
type WriteRules struct {
	Rules 						string 					`json:"rules"`						// rules --> can be : "write_entry_sum" , "write_exit_sum" , "write_crossing_sum"
	Date_from					string					`json:"date_from"`				// from
	Date_to						string					`json:"date_to"`					// to
}

// File info
type FileInfo struct {
	Site							string					`json:"site"`						// site name
	Chain							string					`json:"chain"`						// chain name
	Site_after					string					`json:"site_after"`				// site part after Dst_chain_site_after separator
	Chain_after					string					`json:"chain_after"`				// chain part after Dst_chain_site_after separator
	Access_channel				[]int						`json:"access_channel"`			// Channel id that are access and taking into account for summing Entry and Exits
	Entry_sum					int						`json:"entry_sum"`				// entry sum
	Exit_sum						int						`json:"exit_sum"`					// exit sum
}

// Config struct
type Config struct {
	Src_dir						string					`json:"src_dir"`					// source directory
	Dst_dir						string 					`json:"dst_dir"`					// destination directory
	Nb_file						int						`json:"nb_file"`					// Number of file to try to sum since today
	Dst_file						string					`json:"dst_file"`					// destination file rule containing 'Serial:dst' like '"CPX3-119092","/my/path/to/result/"'
	Write_rule_file			string					`json:"write_rule_file"`		// file containing 'Serial:write_rule:from:to' like '"CPX3-119092","write_exit_sum","20171201","20171214"'
	Write_rules					map[string]*WriteRules `json:"write_rules"`			// write_rules result
	Skip_today					bool						`json:"skip_today"`				// true to skip today files
	File_edit_only_on_diff	bool						`json:"file_edit_only_on_diff"`// true to edit a file only if there is a diff in computed file
	File_no_edit_gzip			bool						`json:"file_no_edit_gzip"`		// true to not edit a gziped file
	File_fix_modtime			bool						`json:"file_fix_modtime"`		// edit modtime file edition to the date of the file
	Verbose						bool						`json:"verbose"`					// true to be verbose
	Serial_dst_dir				map[string]string 	`json:"serial_dst_dir"`			// map for Type-Serial -> Destination directory (processed from Dst_file)
	Must_have_dst				bool						`json:"must_have_dst"`			// true if all serail must be in serial_dst_dir file
	Add_dst_chain_dir			bool						`json:"add_dst_chain_dir"`		// true to add Chain file info to destination directory
	Add_dst_site_dir			bool						`json:"add_dst_site_dir"`		// true to add Site file info to destination directory
	Dst_chain_site_after		string					`json:"dst_chain_site_after"`	// directory id for chain or site (if Add_dst_chain_dir or Add_dst_site_dir is enabled) will be the string after this separator. Ex: if Add_dst_site_dir is enabled with dst_chain_site_after=">>" a site with this chain "My super site >>102" -> the directory destination will be Dst_dir+102
	Write_entry_sum			bool						`json:"write_entry_sum"`		// true to write entry sum in file
	Write_exit_sum				bool						`json:"write_exit_sum"`			// true to write exit sum in file
	Write_crossing_sum		bool						`json:"write_crossing_sum"`	// true to write crossing sum in file
	Gzip_older_than			int						`json:"gzip_older_than"`		// >0 to gzip files that are older than this value (value is number of day from today)
	Gzip_ext_blacklist_src	string					`json:"gzip_ext_blacklist_src"`// list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip,.size')
	Gzip_ext_blacklist		[]string					`json:"gzip_ext_blacklist"`	// list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')
	Periodic_minute			int						`json:"periodic_minute"`		// a value >=0 to restart process every x minute
}

// Functions
// --------------------------

// Display config
func displayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("src_dir:                  %s\n", conf.Src_dir)
	fmt.Printf("dst_dir:                  %s\n", conf.Dst_dir)
	fmt.Printf("nb_file:                  %d\n", conf.Nb_file)
	fmt.Printf("dst_file:                 %s\n", conf.Dst_file)
	fmt.Printf("write_rule_file:          %s\n", conf.Write_rule_file)
	fmt.Printf("write_rules:              %+v\n", conf.Write_rules)
	fmt.Printf("skip_today:               %t\n", conf.Skip_today)
	fmt.Printf("verbose:                  %t\n", conf.Verbose)
	fmt.Printf("serial_dst_dir:           %+v\n", conf.Serial_dst_dir)
	fmt.Printf("add_dst_chain_dir:        %t\n", conf.Add_dst_chain_dir)
	fmt.Printf("add_dst_site_dir:         %t\n", conf.Add_dst_site_dir)
	fmt.Printf("write_entry_sum:          %t\n", conf.Write_entry_sum)
	fmt.Printf("write_exit_sum:           %t\n", conf.Write_exit_sum)
	fmt.Printf("write_crossing_sum:       %t\n", conf.Write_crossing_sum)
	fmt.Printf("gzip_older_than:          %d\n", conf.Gzip_older_than)
	fmt.Printf("gzip_ext_blacklist_src:   %s\n", conf.Gzip_ext_blacklist_src)
	fmt.Printf("gzip_ext_blacklist:       %+v\n", conf.Gzip_ext_blacklist)
	fmt.Printf("periodic_minute:          %d\n", conf.Periodic_minute)
	fmt.Printf(" ----\n")
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

// Get string after first after occurance
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
func creteGzAndDeleteOriginal(original_path string) {
	fi_b, fi_err := ioutil.ReadFile(original_path)
	if nil != fi_err {
		log.Printf("creteGzAndDeleteOriginal -> Error opening (%s) : %s\n", original_path, fi_err.Error()) // Print error
		return
	}

	// Create gz file
	gzip_path := original_path + ".gz"
	fgz, fgz_err := os.Create(gzip_path)
	if nil != fgz_err {
		log.Printf("creteGzAndDeleteOriginal -> Error creating GZ (%s) : %s\n", gzip_path, fgz_err.Error()) // Print error
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
		log.Printf("creteGzAndDeleteOriginal -> Error creating Size (%s) : %s\n", fi_size_path, fi_size_err.Error()) // Print error
		return
	}
	fi_size.Write([]byte(strconv.FormatInt(fi_size_int, 10)))
	fi_size.Close()

	// Delete original file
	del_err := os.Remove(original_path)
	if nil != del_err {
		log.Printf("creteGzAndDeleteOriginal -> Error Deleting (%s) : %s\n", original_path, del_err.Error()) // Print error
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
			creteGzAndDeleteOriginal(f_path)
		}
	} else {
		// PrintAndLog(fmt.Sprintf("-> gzipOldFile GZIPING ALWAYS = %s !\n", f_path), conf, false)// Verbose only
		creteGzAndDeleteOriginal(f_path)
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

// Parse an Eurecam file an return the entries + exits sum
func getSumFile(fdata io.Reader) (file_info FileInfo) {
	file_info = FileInfo{Site:"", Chain:"", Entry_sum:0, Exit_sum:0}

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
				continue
			} else {
				site_chain_header_state = SITE_CHAIN_IDLE
				file_info.Site = ""
				file_info.Chain = ""
			}
		}

		// Check site/chain header
		if (1 == strings.Count(line, ",")) && (SITE_CHAIN_DONE != site_chain_header_state) { //TODO: Should we update Site/Chain header if there is another ?
			if len(line_split) > 0 {
				file_info.Site = line_split[0]
			}
			if len(line_split) > 1 {
				file_info.Chain = line_split[1]
			}
			site_chain_header_state = SITE_CHAIN_IN_PORGRESS
			continue
		}

		// Check Channel type
		if (2 == strings.Count(line, ",")) && (3 == len(line_split)) {
			c_i, c_err := strconv.Atoi(line_split[0])
			if (nil == c_err) && (c_i > 0) {
				contains_already, contains_pos := sliceIntContains(file_info.Access_channel, c_i)
				if ("acces" == line_split[2]) && (!contains_already) {
					file_info.Access_channel = append(file_info.Access_channel, c_i)	// Add this channel to acces channel
				}
				if ("acces" != line_split[2]) && contains_already {
					file_info.Access_channel = append(file_info.Access_channel[:contains_pos], file_info.Access_channel[contains_pos+1:]...) // Remove the channel preserving the slice order see https://github.com/golang/go/wiki/SliceTricks
				}
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

			// loop through Access_channel
			for i := range(file_info.Access_channel) {
				channel_col := file_info.Access_channel[i]*2
				if len(line_split) >= (channel_col+1) {
					// Add channel Entry
					entry_add, entry_add_err := strconv.Atoi(line_split[channel_col])
					if (nil == entry_add_err) {
						file_info.Entry_sum += entry_add
					}
					// Add channel Exits
					exits_add, exits_add_err := strconv.Atoi(line_split[channel_col+1])
					if (nil == exits_add_err) {
						file_info.Exit_sum += exits_add
					}
				}
			}
		}
	}

	return file_info
}

// Sum a file
func sumFile(conf *Config) (filepath.WalkFunc) {
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

			// Check we want this sensor
			dst_path_rule, is_in_dst_rules := conf.Serial_dst_dir[last_dir]
			if (is_in_dst_rules && conf.Must_have_dst) || (!conf.Must_have_dst) {
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

				// Open file to sum
				f, err := os.Open(fpath)
				if err != nil {
					PrintAndLog(fmt.Sprintf("--> ERROR Opening : %s\n", err.Error()), conf, true)// Verbose + Do Log
					return nil
				}
				defer f.Close()

				// Get sum + write result  (from a gziped file or a normale one)
				file_info := FileInfo{}
				if gziped {
					gz, gz_err := gzip.NewReader(f)
					if gz_err != nil {
						PrintAndLog(fmt.Sprintf("--> ERROR unZIPing : %s\n", gz_err.Error()), conf, true)// Verbose + Do log
						return nil
					}
					file_info = getSumFile(gz) // Get sum + file info from gziped file
					gz.Close()
				} else {
					file_info = getSumFile(f) // Get sum + file info from vanilla file
				}
				f.Close()

				PrintAndLog(fmt.Sprintf("Site-Chain: %s-%s -> Sum = %d / %d", file_info.Site, file_info.Chain, file_info.Entry_sum, file_info.Exit_sum), conf, false)//Verbose only

				// Set destination directory
				dst_path := conf.Dst_dir
				if is_in_dst_rules && "" != dst_path_rule {
					dst_path = filepath.ToSlash(dst_path_rule) // If file is in in dst rules we use the content of the file
					PrintAndLog(fmt.Sprintf(" - Rules dst_path -> %s", dst_path), conf, false)//Verbose only
				} else {
					// Get site_after and chain_after
					chain_path := file_info.Chain
					site_path := file_info.Site
					if "" != conf.Dst_chain_site_after {
						chain_path = getStringAfter(file_info.Chain, conf.Dst_chain_site_after)
						site_path = getStringAfter(file_info.Site, conf.Dst_chain_site_after)
					}

					// Write result file
					if conf.Add_dst_chain_dir && ("" != chain_path) {
						dst_path += chain_path + "/"
					}
					if conf.Add_dst_site_dir && ("" != site_path) {
						dst_path += site_path + "/"
					}
					PrintAndLog(fmt.Sprintf(" - Site-Chain dst_path -> %s", dst_path), conf, false)//Verbose only
				}
				dst_path = SanitizeDir(dst_path) // Add last '/' if missing

				// Check for a valid rules in write_rules
				things_to_write := ""
				write_rule, is_in_write_rule := conf.Write_rules[last_dir]
				write_rule_used := false
				if is_in_write_rule {
					write_rule_from := getDateFromEurecam(write_rule.Date_from)
					write_rule_to := getDateFromEurecam(write_rule.Date_to)
					if inTimeSpan(write_rule_from.Add(-12 * time.Hour), write_rule_to.Add(12 * time.Hour), date_file) {
						write_rule_used = true

						// Write_rule can be "write_entry_sum" , "write_exit_sum" , "write_crossing_sum"
						PrintAndLog(fmt.Sprintf(" - Use write_rules -> %s", write_rule.Rules), conf, false)//Verbose only
						if "write_entry_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Entry_sum)
						} else if "write_exit_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Exit_sum)
						} else if "write_crossing_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Entry_sum + file_info.Exit_sum)
						} else if "write_entry_sum,write_exit_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Entry_sum) + "," + strconv.Itoa(file_info.Exit_sum)
						} else if "write_entry_sum,write_crossing_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Entry_sum) + "," + strconv.Itoa(file_info.Entry_sum + file_info.Exit_sum)
						} else if "write_entry_sum,write_exit_sum,write_crossing_sum" == write_rule.Rules {
							things_to_write = strconv.Itoa(file_info.Entry_sum) + "," + strconv.Itoa(file_info.Exit_sum) + "," + strconv.Itoa(file_info.Entry_sum + file_info.Exit_sum)
						}  else {
							write_rule_used = false // Rules is wrong
							PrintAndLog(fmt.Sprintf("--> ERROR Write Rule unknow : %s -> Use config rules instead \n", write_rule.Rules), conf, true)// Verbose + do LOG
						}
					}
				}

				// Use config rules
				if !write_rule_used {
					if conf.Write_entry_sum {
						things_to_write = strconv.Itoa(file_info.Entry_sum)
					}
					if conf.Write_exit_sum {
						if conf.Write_entry_sum {
							things_to_write += ","
						}
						things_to_write += strconv.Itoa(file_info.Exit_sum)
					}
					if conf.Write_crossing_sum {
						if conf.Write_entry_sum || conf.Write_exit_sum {
							things_to_write += ","
						}
						things_to_write += strconv.Itoa(file_info.Entry_sum + file_info.Exit_sum)
					}
				}
				// Final things to write
				b_things_to_write := []byte(things_to_write)

				// Write result
				os.MkdirAll(dst_path, 0777) // always check and create directory before write
				f_path_to_write := dst_path + strings.TrimSuffix(file_name, ".gz") // Original can be a .gz file : fix it for destination

				// Check if new file is different from the one on disk
				need_write := true
				if conf.File_edit_only_on_diff {
					// Check file already exist
					if GetFileSize(f_path_to_write) > 0 {
						// Get file content
						bf, bf_err := ioutil.ReadFile(f_path_to_write)
						if nil != bf_err {
							PrintAndLog(fmt.Sprintf("--> ERROR reading existing file : %s -> %s \n", f_path_to_write, bf_err.Error()), conf, true)// Verbose + do LOG
						} else {
							// Check for a diff in file content
							if bytes.Equal(b_things_to_write, bf) {
								need_write = false
							} else {
								PrintAndLog(fmt.Sprintf("Need to update file : %s -> %s(%s) vs %s\n", f_path_to_write, string(b_things_to_write[:]), things_to_write, string(bf[:])), conf, false)// Verbose only
							}
						}
					}
				}

				if need_write {
					// Write result file
					f_result, f_result_err := os.Create(f_path_to_write)
					if nil != f_result_err {
						PrintAndLog(fmt.Sprintf("--> ERROR Wrinting result : to : %s -> %s \n", f_path_to_write, f_result_err.Error()), conf, true)// Verbose + do LOG
						return nil
					}
					f_result.Write(b_things_to_write)
					f_result.Close()
				}
				if !conf.File_no_edit_gzip || !strings.HasSuffix(file_name, ".gz") {
					gzipOldFile(fpath, conf, true) // Gzip file only if is after n day
				} else {
					PrintAndLog(fmt.Sprintf("--> No gzip edit for : %s\n", f_path_to_write), conf, false)// Verbose only
				}

				// Check modtime is correct
				if conf.File_fix_modtime {
					// Get mod file time
					f_info, f_info_err := os.Open(f_path_to_write)
					if nil != f_info_err {
						PrintAndLog(fmt.Sprintf("--> ERROR Opening file : %s -> %s \n", f_path_to_write, f_info_err.Error()), conf, true)// Verbose + do LOG
						return nil
					}
					// Get modtime
					f_info_stat, f_info_stat_err := f_info.Stat()
					if nil != f_info_stat_err {
						PrintAndLog(fmt.Sprintf("--> ERROR Getting stat file : %s -> %s \n", f_info_stat, f_info_stat_err.Error()), conf, true)// Verbose + do LOG
						return nil
					}

					if (f_info_stat.ModTime().Year() != date_file.Year()) || (f_info_stat.ModTime().Month() != date_file.Month()) || (f_info_stat.ModTime().Day() != date_file.Day()) || (f_info_stat.ModTime().Hour() != 23) || (f_info_stat.ModTime().Second() != 59) {
						var chtime_err error

						// Is it today file ?
						currenttime := time.Now()
						if (currenttime.Year() == date_file.Year()) && (currenttime.Month() == date_file.Month()) && (currenttime.Day() == date_file.Day()) {
							chtime_err = os.Chtimes(f_path_to_write, currenttime, currenttime) // For a today file fix date to today
							PrintAndLog(fmt.Sprintf("--> fix (today) chtime : %s -> from %s to %s\n", f_path_to_write, f_info_stat.ModTime().Format("2006-01-02 15:04:05"), currenttime.Format("2006-01-02 15:04:05")), conf, false)// Verbose only
						} else {
							last_minute_day := time.Date(date_file.Year(), date_file.Month(), date_file.Day(), 23, 59, 59, 0, currenttime.Location())
							chtime_err = os.Chtimes(f_path_to_write, last_minute_day, last_minute_day) // For date in the past set like if file were written the last minute of the day :)
							PrintAndLog(fmt.Sprintf("--> fix chtime : %s -> from %s to %s\n", f_path_to_write, f_info_stat.ModTime().Format("2006-01-02 15:04:05"), last_minute_day.Format("2006-01-02 15:04:05")), conf, false)// Verbose only
						}
						if nil != chtime_err {
							PrintAndLog(fmt.Sprintf("--> ERROR Setting stat file : %s -> %s \n", f_path_to_write, chtime_err.Error()), conf, true)// Verbose + do LOG
						}
					}
				}

				PrintAndLog(fmt.Sprintf("----> OK\n"), conf, false)//Verbose only
			} else {
				// Skip file, because not in our config
				gzipOldFile(fpath, conf, false) // Always gzip this files
				PrintAndLog(fmt.Sprintf("--> SKIP file %s -> %s not wanted!!\n", fpath, last_dir), conf, true)// Verbose + do LOG
				return nil
			}
		}

		return nil
	}
}

// Read file rule
func getFileRules(conf *Config) (error, map[string]string) {
	map_rules := make(map[string]string)
	f_data, f_err := ioutil.ReadFile(conf.Dst_file)
	if f_err != nil {
		if "" != conf.Dst_file {
			PrintAndLog(fmt.Sprintf("Error : getFileRules can't open file: %s -> %s \n", conf.Dst_file, f_err.Error()), conf, true)// Verbose + do LOG
		}
		return f_err, map_rules
	}
	r := csv.NewReader(strings.NewReader(string(f_data)))

	// Parse csv data
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			PrintAndLog(fmt.Sprintf("Error : getFileRules can't read rule -> %s \n", err.Error()), conf, true)// Verbose + do LOG
			continue
		}

		if len(record) > 1 {
			map_rules[record[0]] = record[1]
		} else {
			PrintAndLog(fmt.Sprintf("Rule file must have 2 record : found only one -> %s \n", record[0]), conf, true)// Verbose + do LOG
			continue
		}

		// fmt.Println(record)
		// fmt.Printf("%s - %s - %d\n", record[0], record[1], len(record))
	}

	return nil, map_rules
}

// Read file rule
func getFileWriteRules(conf *Config) (error, map[string]*WriteRules) {
	map_write_rules := make(map[string]*WriteRules)
	f_data, f_err := ioutil.ReadFile(conf.Write_rule_file)
	if f_err != nil {
		if "" != conf.Dst_file {
			PrintAndLog(fmt.Sprintf("Error : getFileWriteRules can't open file: %s -> %s \n", conf.Write_rule_file, f_err.Error()), conf, true)// Verbose + do LOG
		}
		return f_err, map_write_rules
	}
	r := csv.NewReader(strings.NewReader(string(f_data)))

	// Parse csv data
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			PrintAndLog(fmt.Sprintf("Error : getFileWriteRules can't read rule -> %s \n", err.Error()), conf, true)// Verbose + do LOG
			continue
		}

		if len(record) > 3 {
			map_write_rules[record[0]] = &WriteRules{record[1], record[2], record[3]}
		} else {
			PrintAndLog(fmt.Sprintf("WriteRule file must have 4 record : found only %d -> %s \n", len(record), record[0]), conf, true)// Verbose + do LOG
			continue
		}

		// fmt.Println(record)
		// fmt.Printf("%s - %s - %d\n", record[0], record[1], len(record))
	}

	return nil, map_write_rules
}

// Loop to sum all files
func sumAllFiles(conf *Config) (err error) {
	err = filepath.Walk(conf.Src_dir, sumFile(conf))
	if err != nil {
		PrintAndLog(fmt.Sprintf("Error : sumAllFiles: %s \n", err.Error()), conf, true)// Verbose + do LOG
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
				_, conf.Serial_dst_dir = getFileRules(conf)
				_, conf.Write_rules = getFileWriteRules(conf)
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
	default_json_config_path := "./sumer.json"

	// Flags
	Src_dir := flag.String("src_dir", "./Data/", "Source directory that store files to parse")
	Dst_dir := flag.String("dst_dir", "./Dst/", "Destination directory for result sum files")
	Nb_file := flag.Int("nb_file", 3, "Number of file to consider from today, -1 for all days (Note: file not present in destination will still be written)")
	Dst_file := flag.String("dst_file", "./dst_rules.csv", "destination file rule containing 'Serial:dst' like 'CPX3-119092:/my/path/to/result/'")
	Write_rule_file := flag.String("write_rule_file", "./write_rules.csv", "file containing Rules for overriding config write type : Entry/Exit from date to date : 'Serial,write_rules,from,to' like '\"CPX3-119092\",\"write_exit_sum\",\"20171201\",\"20171214\"'")
	Skip_today := flag.Bool("skip_today", false, "Add this flag to skip today file (file in destination will arrive only when the day is over")
	File_edit_only_on_diff := flag.Bool("file_edit_only_on_diff", false, "Edit a file only if there is a diff in computed file")
	File_no_edit_gzip := flag.Bool("file_no_edit_gzip", false, "Do not edit a gziped file")
	File_fix_modtime := flag.Bool("file_fix_modtime", false, "Edit access time and modification time of the file to match file name")
	Verbose := flag.Bool("verbose", false, "Set it to be verbose")
	Must_have_dst := flag.Bool("must_have_dst", false, "Set it if all serial must be in Dst_file file")
	Add_dst_chain_dir := flag.Bool("add_dst_chain_dir", false, "Set it to add Chain file info to destination directory")
	Add_dst_site_dir := flag.Bool("add_dst_site_dir", false, "Set it to add Site file info to destination directory")
	Dst_chain_site_after := flag.String("dst_chain_site_after", "", "Set it to add a separator to added chain-site destination ex: '-' will add 'Chain-Site' to destination directory")
	Write_entry_sum := flag.Bool("write_entry_sum", false, "Set it to write entry sum in result file")
	Write_exit_sum := flag.Bool("write_exit_sum", false, "Set it to write exits sum in result file")
	Write_crossing_sum := flag.Bool("write_crossing_sum", false, "Set it to write crossing sum in result file")
	Gzip_older_than := flag.Int("gzip_older_than", 35, "Set a value greater than 0 to gzip files older than this value (value is number of day from today)")
	Gzip_ext_blacklist_src := flag.String("gzip_ext_blacklist_src", ".jpg,.gz,.zip,.gzip,.size", "List of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')")
	Periodic_minute := flag.Int("periodic_minute", -1, "Minute periodicity to restart parsing process (a value negative or 0 will not restart process")
	Log_file := flag.String("log", "./log.log", "Destination log file")
	Config_json := flag.Bool("config_json", false, "Set this flag to read config from sumer.json file (Default no read from sumer.json)")
	Config_json_path := flag.String("config_json_path", default_json_config_path, "full path to file config including his name (ex:'my/super/path/to/my_config.json')")
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
		conf.Dst_file = *Dst_file
		conf.Write_rule_file = *Write_rule_file
		conf.Skip_today = *Skip_today
		conf.File_edit_only_on_diff = *File_edit_only_on_diff
		conf.File_no_edit_gzip = *File_no_edit_gzip
		conf.File_fix_modtime = *File_fix_modtime
		conf.Verbose = *Verbose
		conf.Must_have_dst = *Must_have_dst
		conf.Add_dst_chain_dir = *Add_dst_chain_dir
		conf.Add_dst_site_dir = *Add_dst_site_dir
		conf.Dst_chain_site_after = *Dst_chain_site_after
		conf.Write_entry_sum = *Write_entry_sum
		conf.Write_exit_sum = *Write_exit_sum
		conf.Write_crossing_sum = *Write_crossing_sum
		// conf.Log_file
		conf.Gzip_older_than = *Gzip_older_than
		conf.Gzip_ext_blacklist_src = *Gzip_ext_blacklist_src
		conf.Gzip_ext_blacklist = strings.Split(*Gzip_ext_blacklist_src, ",")
		conf.Periodic_minute = *Periodic_minute
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
			fmt.Printf("sumer started at %s → src: %s , dst: %s , nb_file: %d\n---\n", time.Now().Format("2006-01-02 15:04:05"), conf.Src_dir, conf.Dst_dir, conf.Nb_file)
		}
	}

	// Sum all file according to config
	if conf.Periodic_minute > 0 {
		// Deamon mode : run each periodic_minute
		PrintAndLog(fmt.Sprintf("Started with periodicity = %d minutes\n", conf.Periodic_minute), &conf, true)// Verbose + do LOG
		for {
			if (true == *Config_json) || (default_json_config_path != *Config_json_path) {	// Try to re-read json config
				if !readConfigFile(&conf, *Config_json_path) {
					fmt.Printf("Error reading json file : config not changed\n")
				}
			}

			// Re-sum file
			PrintAndLog(fmt.Sprintf("Re-Started because periodicity = %d minutes\n", conf.Periodic_minute), &conf, false)// Verbose only
			sumAllFiles(&conf)
			time.Sleep(time.Duration(conf.Periodic_minute) * time.Minute)
		}
	}

	// One shoot mode : run just once
	sumAllFiles(&conf)
}
