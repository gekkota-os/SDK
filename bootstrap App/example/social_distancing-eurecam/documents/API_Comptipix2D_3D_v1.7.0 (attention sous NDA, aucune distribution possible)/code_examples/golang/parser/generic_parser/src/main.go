// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o parser.exe parser.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o parser.exe parser.go
//
// * RUN :
// ./parser -h
//

package main

import (
	"fmt"
	"log"
	"pconf"
	"pformat"
	"pparse"
	"putil"
	"sort"
	"strings"
	"sync"

	// "encoding/csv"
	"compress/gzip"
	"flag"
	"os"
	"path/filepath"
	"time"
)

// SumedFile file already sumed (serial name sorted as key)
type SumedFile struct {
	files map[string]bool
	mux   sync.RWMutex
}

// Functions
// --------------------------

// Parse an Comptipix file and save the output file
// in out directory (keeping last directory hierachy)
func walkParseFile(conf *pconf.Config, sumed SumedFile) filepath.WalkFunc {
	// Use a closure to pass argument to walk function (pass the config)
	return func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			putil.PrintAndLog(fmt.Sprintf("Error : walkParseFile dir -> %s \n", err.Error()), conf.Verbose, false) // Verbose only
			return nil
		}
		if info.IsDir() {
			return nil // Don't care of directory
		} else {
			// We have a file
			dir, file_name := filepath.Split(fpath)
			last_dir := filepath.Base(dir)

			// fmt.Printf("dir-in: %s\n", dir)

			// Check we want this file (check for a Comptipix counting or occupancy file)
			lenFile := len("20180720.csv")
			if conf.In_format == 1 {
				lenFile = len("20180720_presence.csv")
			}
			gziped := false
			putil.PrintAndLog(fmt.Sprintf("File: %s (last_dir:%s, name:%s)", fpath, last_dir, file_name), conf.Verbose, false) // Verbose only
			if (lenFile != len(file_name)) || (".csv" != filepath.Ext(file_name)) {                                            // Skip not counting file -> like "20171207.csv"
				if (lenFile+3 == len(file_name)) && (strings.HasSuffix(file_name, ".csv.gz")) {
					gziped = true
				} else {
					putil.GzipOldFile(fpath, conf, false)                                                       // Always gzip this files
					putil.PrintAndLog(fmt.Sprintf("--> SKIP file -> not wanted file!!\n"), conf.Verbose, false) // Verbose only
					return nil
				}
			}

			// Get date from file name
			file_suffix := ".csv"
			if gziped {
				file_suffix = ".csv.gz"
			}
			date_file := putil.GetDateFromComptipixFileName(strings.TrimSuffix(file_name, file_suffix))

			// Check file is in TimeSpan we want to test
			if conf.Nb_file > 0 {
				date_end := putil.GetNowMidnight()
				date_start := date_end.Add(time.Duration(-1*conf.Nb_file*24) * time.Hour)

				// Check skip today
				if conf.Skip_today && (date_file.Year() == date_end.Year()) && (date_file.Month() == date_end.Month()) && (date_file.Day() == date_end.Day()) {
					putil.PrintAndLog(fmt.Sprintf("--> SKIP file date(%s) -> because skip_today(%s)!!\n", date_file.Format("2006-01-02 15:04:05"), date_end.Format("2006-01-02 15:04:05")), conf.Verbose, false) //Verbose only
					return nil
				}

				// Check date is in day wanted
				if !putil.InTimeSpan(date_start, date_end.Add(12*time.Hour), date_file) { // Add 12hours to date_end which is today at midnight (to be sure today file will be processed)
					putil.GzipOldFile(fpath, conf, true)                                                                                                                                                                                                                  // Gzip file only if is after n day
					putil.PrintAndLog(fmt.Sprintf("--> SKIP file date(%s) -> not in nb file wanted (from:%s to %s)!!\n", date_file.Format("2006-01-02 15:04:05"), date_start.Format("2006-01-02 15:04:05"), date_end.Format("2006-01-02 15:04:05")), conf.Verbose, false) //Verbose only
					return nil
				}
			}

			// Open file to parse
			f, err := os.Open(fpath)
			if err != nil {
				putil.PrintAndLog(fmt.Sprintf("--> ERROR Opening : %s\n", err.Error()), conf.Verbose, true) // Verbose + Do Log
				return nil
			}
			defer f.Close()

			// fmt.Printf("Here %s \n",file_name)//DEBUG

			// Get parsed result  (from a gziped file or a normale one)
			file_parsed := pparse.FileParsed{}
			if gziped {
				gz, gz_err := gzip.NewReader(f)
				if gz_err != nil {
					putil.PrintAndLog(fmt.Sprintf("--> ERROR unZIPing : %s\n", gz_err.Error()), conf.Verbose, true) // Verbose + Do log
					return nil
				}
				file_parsed = pparse.GetParsedFile(gz, &date_file, conf.Out_resolution) // Get sum + file info from gziped file
				gz.Close()
			} else {
				file_parsed = pparse.GetParsedFile(f, &date_file, conf.Out_resolution) // Get sum + file info from vanilla file
			}
			f.Close()

			// Sum files
			sumKey := ""
			if len(conf.Out_sum_file) > 1 {
				// Check this serial must be sumed or substracted
				mustSumOrSub, _ := putil.SliceStringContains(conf.Out_sum_file, last_dir)
				if mustSumOrSub {
					// Check not already sumed
					sort.Strings(conf.Out_sum_file)
					sumKey = strings.Join(conf.Out_sum_file, "_")
					sumed.mux.RLock()
					_, sumDone := sumed.files[sumKey]
					sumed.mux.RUnlock()
					if !sumDone {
						// Prepare list of file to sum or substract
						sumParsed := []pparse.FileParsed{}
						sumParsed = append(sumParsed, file_parsed)

						// mark file position to substract
						toSub := []bool{}
						mustSub, _ := putil.SliceStringContains(conf.Out_sub_file, last_dir)
						if mustSub {
							toSub = append(toSub, true)
						} else {
							toSub = append(toSub, false)
						}

						// Get and parse all file need for sum
						for toSum := range conf.Out_sum_file {
							// Check file exist
							fSumPath := putil.SanitizeDir(conf.In_dir) + putil.SanitizeDir(conf.Out_sum_file[toSum]) + file_name
							if putil.GetFileSize(fSumPath) > 0 {
								// Open file to parse
								fSum, fSumErr := os.Open(fSumPath) //TODO mnanage gziped
								if fSumErr != nil {
									putil.PrintAndLog(fmt.Sprintf("--> ERROR Opening sum : %s\n", err.Error()), conf.Verbose, true) // Verbose + Do Log
									continue                                                                                        // skip to next file
								}
								parsedToSum := pparse.GetParsedFile(fSum, &date_file, conf.Out_resolution) // Get sum + file info from vanilla file
								sumParsed = append(sumParsed, parsedToSum)
								fSum.Close()

								// Add file to subsctract, or not
								mustSub, _ = putil.SliceStringContains(conf.Out_sub_file, conf.Out_sum_file[toSum])
								if mustSub {
									toSub = append(toSub, true)
								} else {
									toSub = append(toSub, false)
								}
							}
						}

						// Sum parsed file :
						file_parsed = pparse.SumParsedFile(sumParsed, toSub, conf.Out_fix_occupancy_positive)
						sumed.mux.Lock()
						sumed.files[sumKey] = true
						sumed.mux.Unlock()
					} else {
						sumKey = "" // Sum already done
					}
				}
			}

			// Write new file in destination
			out_name, out_content := file_parsed.GetOutFile(conf)
			out_name_ok_dir, out_name_ok := filepath.Split(out_name)
			dst_path := putil.SanitizeDir(conf.Out_dir) + putil.SanitizeDir(last_dir) + putil.SanitizeDir(out_name_ok_dir)
			if "" != sumKey {
				fmt.Printf("Sum file %s/%s\n", last_dir, out_name_ok_dir)
				dst_path = putil.SanitizeDir(conf.Out_dir) + putil.SanitizeDir(sumKey) + putil.SanitizeDir(out_name_ok_dir)
			}
			fmt.Printf("PATH: %s, FILE: %s, DATE: %s\n", dst_path, out_name_ok, date_file.Format("2006-01-02 15:04:05"))
			os.MkdirAll(dst_path, 0777) // always check and create directory before write
			f_path_to_write := dst_path + out_name_ok
			f_result, f_result_err := os.Create(f_path_to_write)
			if f_result_err != nil {
				putil.PrintAndLog(fmt.Sprintf("--> ERROR Writing result : to : %s -> %s \n", f_path_to_write, f_result_err.Error()), conf.Verbose, true) // Verbose + do LOG
				return nil
			}
			f_result.Write([]byte(out_content))
			f_result.Close()
			putil.GzipOldFile(fpath, conf, true) // Gzip file only if is after n day

			putil.PrintAndLog(fmt.Sprintf(" -> Outfile to %s \n", out_name), conf.Verbose, true) // Verbose + NO LOG
			return nil
		}
	}
}

// Parse all files in directory
func parseAllFiles(conf *pconf.Config) (err error) {
	sumed := SumedFile{files: make(map[string]bool)}
	err = filepath.Walk(conf.In_dir, walkParseFile(conf, sumed))
	if err != nil {
		putil.PrintAndLog(fmt.Sprintf("Error : parseAllFiles: %s \n", err.Error()), conf.Verbose, true) // Verbose + do LOG
	}

	return err
}

// Main
// Parse option and start to read files
func main() {
	// Read arg
	// ----------------
	help_demo := "This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE"
	default_json_config_path := "./parser.json"
	version := "0.2"

	// Help message
	Help_in_format := fmt.Sprintf("Format file input :\n") + pformat.PrintAllInFormat()
	Help_out_format := fmt.Sprintf("Format file output :\n") + pformat.PrintAllOutFormat()
	Help_out_channel := fmt.Sprintf("Channel to keep in output :\n") + pformat.PrintAllOutChannel()
	Help_out_date_format_type := fmt.Sprintf("Date format (exemple for on Jan 2 2006 at 15:04:05, GMT-0700) (NOTE: set this to override file definition) :\n") + pformat.PrintAllDateFormatType()
	Help_out_time_format_type := fmt.Sprintf("Time format (exemple for 15:04:05, GMT-0700) (NOTE: set this to override file definition) :\n") + pformat.PrintAllTimeFormatType()
	Help_file_end_line := fmt.Sprintf("File end line char (NOTE: set this to override file definition) :\n") + pformat.PrintAllFileEOL()

	// Flags
	In_dir := flag.String("in_dir", "./Data-in/", "Input source directory that store files to parse")
	Out_dir := flag.String("out_dir", "./Data-out/", "Output destination directory for resulting files")
	In_format := flag.Int("in_format", 0, Help_in_format)
	Nb_file := flag.Int("nb_file", 33, "Number of file to consider from today, -1 for all days")
	Out_format := flag.Int("out_format", 0, Help_out_format)
	Out_resolution := flag.Int64("out_resolution", 1800, "Data output resolution wanted in seconds")
	Out_sum_file_src := flag.String("out_sum_file_src", "", "List of directories (Sensor serial) to sum file")
	Out_fix_occupancy_positive := flag.Bool("out_fix_occupancy_positive", true, "Fix output occupancy to be always positive (add correction+ to fix it)")
	Out_channel := flag.Int("out_channel", 0, Help_out_channel)
	Out_channel_select_src := flag.String("out_channel_select_src", "", "Output channel id use separated by ',' (channel id are integer), exemple '1,3' -> use channel 1 and 3")
	Out_strip_header := flag.Bool("out_strip_header", false, "Set this flag to strip header from output file (Default: disabled)")
	Out_strip_null_line := flag.Bool("out_strip_null_line", false, "Set this flag to strip line with no counting (Default: disabled)")
	Out_date_format_type := flag.Int("out_date_format_type", 0, Help_out_date_format_type)
	Out_date_format_sep := flag.String("out_date_format_sep", "", "Output date format separator\n \t ex: for Mon Jan 2 2006 : -out_date_format_type=2 -out_date_format_sep='-' -> 2006-01-02 (NOTE: set this to override file definition)")
	Out_time_format_type := flag.Int("out_time_format_type", 0, Help_out_time_format_type)
	Out_time_format_sep := flag.String("out_time_format_sep", "", "Output time format separator\n \tex: 15:04:05 : -out_time_format_type=2 -out_time_format_sep=':' -> 15:04:05 (NOTE: set this to override file definition)")
	Out_date_time_sep := flag.String("out_date_time_sep", ",", "Output date time format separator\n \tex: '_' -> 2006-01-02_15:04:05 (NOTE: set this to override file definition)")
	Out_date_time_order := flag.Int("out_date_time_order", 0, "Output date and time order : \n\t0:Date before time (ex: 2006-01-02 15:04:05)\n\t1:time before date (ex: 15:04:05 2006-01-02) (NOTE: set this to override file definition)")
	Out_file_end_line := flag.Int("out_file_end_line", 0, Help_file_end_line)
	Out_opening_hour_src := flag.String("out_opening_hour_src", "", "Opening hour source separted by ',' like 8h30,12h30,14h,18h30 -> for an opening of 8h30 to 12h30 then 14h to 18h30")
	Skip_today := flag.Bool("skip_today", false, "Add this flag to skip today file (file in destination will arrive only when the day is ove)")
	Verbose := flag.Bool("verbose", false, "Set parser to be verbose")
	Gzip_older_than := flag.Int("gzip_older_than", 35, "Set a value greater than 0 to gzip source files older than this value (value is number of day from today)")
	Gzip_ext_blacklist_src := flag.String("gzip_ext_blacklist_src", ".jpg,.gz,.zip,.gzip,.size", "List of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')")
	Periodic_minute := flag.Int("periodic_minute", -1, "Minute periodicity to restart parsing process (a value negative or 0 will not restart process")
	Log_file := flag.String("log", "./log.log", "Destination log file")
	Config_json := flag.Bool("config_json", false, "Set this flag to read config from parser.json file (Default no read from parser.json)")
	Config_json_path := flag.String("config_json_path", default_json_config_path, "full path to file config including his name (ex:'my/super/path/to/my_config.json')")
	Version := flag.Bool("version", false, "Show version")
	// help := flag.Bool("help", false, "Print help usage and demo")
	flag.Parse()

	// Version
	if *Version {
		fmt.Printf("Parser version : %s\n", version)
		return // do nothing else
	}

	// Always show this message
	fmt.Println(help_demo)

	// read config from JSON
	var conf pconf.Config
	if (true == *Config_json) || (default_json_config_path != *Config_json_path) {
		// Read config file
		fmt.Printf("\n* Read json file %s :\n", *Config_json_path)
		if !pconf.ReadConfFile(&conf, *Config_json_path) {
			fmt.Printf("Error reading json file ABORT\n")
			return
		}
		fmt.Printf("Read json config (%s) OK\n", *Config_json_path)
	} else {
		// Because we have to do it after flag.Parse()
		conf.In_dir = filepath.ToSlash(*In_dir) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
		conf.Out_dir = filepath.ToSlash(*Out_dir)
		conf.In_format = *In_format
		conf.Nb_file = *Nb_file
		conf.Out_format = *Out_format
		conf.Out_resolution = *Out_resolution
		pconf.SetSumSubFile(&conf)
		conf.Out_sum_file_src = *Out_sum_file_src
		conf.Out_fix_occupancy_positive = *Out_fix_occupancy_positive
		conf.Out_channel = *Out_channel
		conf.Out_channel_select_src = *Out_channel_select_src
		pconf.SetConfOutChannelSelected(&conf)
		conf.Out_strip_header = *Out_strip_header
		conf.Out_strip_null_line = *Out_strip_null_line
		conf.Out_date_format_type = *Out_date_format_type
		conf.Out_date_format_sep = *Out_date_format_sep
		conf.Out_date_time_sep = *Out_date_time_sep
		conf.Out_date_time_order = *Out_date_time_order
		conf.Out_time_format_type = *Out_time_format_type
		conf.Out_time_format_sep = *Out_time_format_sep
		conf.Out_file_end_line = *Out_file_end_line
		conf.Out_opening_hour_src = *Out_opening_hour_src
		conf.Skip_today = *Skip_today
		conf.Verbose = *Verbose
		conf.Gzip_older_than = *Gzip_older_than
		conf.Gzip_ext_blacklist_src = *Gzip_ext_blacklist_src
		conf.Gzip_ext_blacklist = strings.Split(*Gzip_ext_blacklist_src, ",")
		conf.Periodic_minute = *Periodic_minute
		pconf.SetConfDateTimeFormat(&conf)
	}
	pconf.DisplayConfig(&conf)

	// Set log file
	if *Log_file != "" {
		f, f_err := os.OpenFile(*Log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if f_err != nil {
			fmt.Printf("Error opening log file(%s): %v !!!\n", *Log_file, f_err)
		} else {
			defer f.Close()
			log.SetOutput(f)
			fmt.Printf("Parser started at %s → in: %s , out: %s , nb_file: %d\n ---\n", time.Now().Format("2006-01-02 15:04:05"), conf.In_dir, conf.Out_dir, conf.Nb_file)
		}
	}

	// Parse all file according to config
	if conf.Periodic_minute > 0 {
		// Deamon mode : run each periodic_minute
		putil.PrintAndLog(fmt.Sprintf("Started with periodicity = %d minutes\n", conf.Periodic_minute), conf.Verbose, false) // Verbose + no LOG
		for {
			if (true == *Config_json) || (default_json_config_path != *Config_json_path) { // Try to re-read json config
				if !pconf.ReadConfFile(&conf, *Config_json_path) {
					fmt.Printf("Error reading json file : config remain unchanged\n")
				}
			}

			// Re-parse file
			putil.PrintAndLog(fmt.Sprintf("Re-Started because periodicity = %d minutes\n", conf.Periodic_minute), conf.Verbose, false) // Verbose only
			parseAllFiles(&conf)
			time.Sleep(time.Duration(conf.Periodic_minute) * time.Minute)
		}
	}

	// One shoot mode : run just once
	parseAllFiles(&conf)
}
