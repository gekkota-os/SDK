package pconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pformat"
	"strconv"
	"strings"
	"time"
)

// OpeningHour define a couple of an open and close hour
type OpeningHour struct {
	Open  time.Duration `json:"open"`  // Opening hour
	Close time.Duration `json:"close"` // Closing hour
}

// Config struct
type Config struct {
	In_dir                     string                 `json:"in_dir"`                     // Input source directory
	Out_dir                    string                 `json:"out_dir"`                    // Destination directory
	Nb_file                    int                    `json:"nb_file"`                    // Number of file to try to sum since today
	In_format                  int                    `json:"in_format"`                  // Input format (see \ref Input)
	Out_format                 int                    `json:"out_format"`                 // Output format (see \ref OutFormat)
	Out_resolution             int64                  `json:"out_resolution"`             // Data resolution wanted in seconds
	Out_sum_file_src           string                 `json:"out_sum_file_src"`           // Sum file source
	Out_sum_file               []string               `json:"out_sum_file"`               // Array of file to sum
	Out_sub_file               []string               `json:"out_sub_file"`               // Array of file to sub
	Out_fix_occupancy_positive bool                   `json:"out_fix_occupancy_positive"` // true to fix output occupancy to a positive value
	Out_channel                int                    `json:"out_channel"`                // Output channel to use (see \ref OutChannel)
	Out_channel_select_src     string                 `json:"out_channel_select_src"`     // Output channel separated by ',' -> user source string
	Out_channel_select         []int                  `json:"out_channel_select"`         // (internal) -> computed from Out_channel_select_src : Output channel to use
	Out_strip_header           bool                   `json:"out_strip_header"`           // true to set output file without header
	Out_strip_null_line        bool                   `json:"out_strip_null_line"`        // true to strip line with no counting
	Out_date_format_type       int                    `json:"out_date_format_type"`       // Output date format to use (see \ref DateFormatType)
	Out_date_format_sep        string                 `json:"out_date_format_sep"`        // Output date format separator (ex : "-" to have '2018-02-30')
	Out_time_format_type       int                    `json:"out_time_format_type"`       // Output type format to use (see \ref TimeFormatType)
	Out_time_format_sep        string                 `json:"out_time_format_sep"`        // Output type format separator (ex : ":" to have '15:04:05')
	Out_date_time_sep          string                 `json:"out_date_time_sep"`          // Output type date time separator (ex : "_" to have '2018-02-30_15:04:05')
	Out_date_time_order        int                    `json:"out_date_time_order"`        // Output date-time order
	Out_file_end_line          int                    `json:"out_file_end_line"`          // LF, CRLF or CR
	Out_opening_hour_src       string                 `json:"out_opening_hour_src"`       // Opening hour source separted by , like 8h30,12h30,14h,18h30
	Out_opening_hour           []OpeningHour          `json:"out_opening_hour"`           // Opening hour struct of time
	Date_time_format           pformat.DateTimeFormat `json:"date_time_format"`           // (internal) -> computed from Out_date_format_type/Out_date_format_sep/Out_time_format_type/Out_time_format_sep : date and time format type see \ref DateTimeFormat -> result of out_format + out_date_format_type...out_date_time_order
	Skip_today                 bool                   `json:"skip_today"`                 // true to skip today files
	Verbose                    bool                   `json:"verbose"`                    // true to be verbose
	Gzip_older_than            int                    `json:"gzip_older_than"`            // >0 to gzip files that are older than this value (value is number of day from today)
	Gzip_ext_blacklist_src     string                 `json:"gzip_ext_blacklist_src"`     // list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip,.size')
	Gzip_ext_blacklist         []string               `json:"gzip_ext_blacklist"`         // list of extension comma separated that shouldn't be zipped (ex: '.jpg,.gz,.zip,.gzip')
	Periodic_minute            int                    `json:"periodic_minute"`            // a value >=0 to restart process every x minute
}

// Public Functions
// --------------------------

// SetSumSubFile set sum and substract file
func SetSumSubFile(conf *Config) (read_ok bool) {
	read_ok = true
	conf.Out_sum_file = nil //Reset

	fmt.Printf("!!! SetSumFile %s\n", conf.Out_sum_file_src)

	// Read file to sum
	read_ok = true
	if "" != conf.Out_sum_file_src {
		// Mark all file to sum or sub
		filesName := strings.Split(conf.Out_sum_file_src, ",")
		if len(filesName) < 2 {
			read_ok = false
		} else {
			for i := range filesName {
				if strings.HasPrefix(filesName[i], "--") {
					// Set file to sum + sub
					conf.Out_sub_file = append(conf.Out_sub_file, strings.TrimPrefix(filesName[i], "--"))
					conf.Out_sum_file = append(conf.Out_sum_file, strings.TrimPrefix(filesName[i], "--"))
				} else {
					// Set file to sum
					conf.Out_sum_file = append(conf.Out_sum_file, filesName[i])
				}
			}
		}
	}

	// fmt.Printf("%+v\n", conf.Out_sum_file)

	return read_ok
}

// SetConfOutChannelSelected set Config Out_channel_selected from Out_channel_selected_src
func SetConfOutChannelSelected(conf *Config) (read_ok bool) {
	read_ok = true
	conf.Out_channel_select = nil // Reset

	if "" != conf.Out_channel_select_src {
		tmp_out_channel_select := strings.Split(conf.Out_channel_select_src, ",")
		for k := range tmp_out_channel_select {
			k_i, k_err := strconv.Atoi(tmp_out_channel_select[k])
			if nil == k_err {
				conf.Out_channel_select = append(conf.Out_channel_select, k_i)
			} else {
				fmt.Printf("Error Out_channel_select_src %d : %s not an integer (%s)\n", k, tmp_out_channel_select[k], k_err.Error())
				read_ok = false
			}
		}
	}

	return read_ok
}

// SetOpeningHour set couples of open/close hour reading Out_opening_hour_src
func SetOpeningHour(conf *Config) (read_ok bool) {
	read_ok = true
	conf.Out_opening_hour = nil //Reset

	if "" != conf.Out_opening_hour_src {
		// Read opening and convert it to second
		tmp_out_opening_hour := strings.Split(conf.Out_opening_hour_src, ",")
		tmp_opening_open := time.Duration(0)
		for k := range tmp_out_opening_hour {
			Duration, DurationErr := time.ParseDuration(tmp_out_opening_hour[k])
			if nil == DurationErr {
				if 0 != k%2 {
					tmp_opening := &OpeningHour{Open: tmp_opening_open, Close: Duration}
					conf.Out_opening_hour = append(conf.Out_opening_hour, *tmp_opening)
					// fmt.Printf("%d (%s) - Close = %s (%d)\n", k, tmp_out_opening_hour[k], Duration, tmp_opening.Close)
				} else {
					tmp_opening_open = Duration
					// fmt.Printf("%d (%s) - Open = %s (%d)\n", k, tmp_out_opening_hour[k], Duration, tmp_opening_open)
				}
			} else {
				fmt.Printf("Error reading open/close hour at pos %d (%s), %s\n", k, tmp_out_opening_hour[k], DurationErr.Error())
				read_ok = false
			}
		}
	}

	// fmt.Printf("%+v\n", conf.Out_opening_hour)

	return read_ok
}

// SetConfDateTimeFormat set user DateTime format override
func SetConfDateTimeFormat(conf *Config) {
	conf.Date_time_format.Date_format_type = pformat.DateFormatType(conf.Out_date_format_type)
	conf.Date_time_format.Date_format_sep = conf.Out_date_format_sep
	conf.Date_time_format.Time_format_type = pformat.TimeFormatType(conf.Out_time_format_type)
	conf.Date_time_format.Time_format_sep = conf.Out_time_format_sep
	conf.Date_time_format.Date_time_sep = conf.Out_date_time_sep
	conf.Date_time_format.Date_time_order = pformat.DateTimeFormatOrder(conf.Out_date_time_order)
}

// DisplayOpeningHour display opening hour
func DisplayOpeningHour(conf *Config) {
	if len(conf.Out_opening_hour) > 0 {
		fmt.Printf("Opening hour : \n")
		for k := range conf.Out_opening_hour {
			fmt.Printf("\t%d: From %s to %s\n", k, conf.Out_opening_hour[k].Open, conf.Out_opening_hour[k].Close)
		}
	}
}

// DisplayConfig Display config nicely
func DisplayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("in_dir:                      %s\n", conf.In_dir)
	fmt.Printf("out_dir:                     %s\n", conf.Out_dir)
	fmt.Printf("nb_file:                     %d\n", conf.Nb_file)
	fmt.Printf("in_format:                   %d -> %s\n", conf.In_format, pformat.PrintInFormat(conf.In_format))
	fmt.Printf("out_format:                  %d -> %s\n", conf.Out_format, pformat.PrintOutFormat(conf.Out_format))
	fmt.Printf("out_resolution:              %d -> %s\n", conf.Out_resolution, pformat.PrintResolution(conf.Out_resolution))
	fmt.Printf("out_sum_file:                %+v\n", conf.Out_sum_file)
	fmt.Printf("out_sub_file:                '-> %+v\n", conf.Out_sub_file)
	fmt.Printf("out_channel:                 %d -> %s\n", conf.Out_channel, pformat.PrintOutChannel(conf.Out_channel, conf.Out_channel_select_src))
	fmt.Printf("out_channel_select_src:      %s\n", conf.Out_channel_select_src)
	fmt.Printf("out_channel_select:          %+v\n", conf.Out_channel_select)
	fmt.Printf("out_strip_header:            %t\n", conf.Out_strip_header)
	fmt.Printf("out_strip_null_line:         %t\n", conf.Out_strip_null_line)
	fmt.Printf("out_date_format_type:        %d -> %s\n", conf.Out_date_format_type, pformat.PrintDateFormatType(conf.Out_date_format_type))
	fmt.Printf("out_date_format_sep:         %s\n", conf.Out_date_format_sep)
	fmt.Printf("out_file_end_line:           %d -> %s\n", conf.Out_file_end_line, pformat.PrintFileEOL(conf.Out_file_end_line))
	fmt.Printf("out_opening_hour:            %s -> ", conf.Out_opening_hour_src)
	DisplayOpeningHour(conf)
	fmt.Printf("skip_today:                  %t\n", conf.Skip_today)
	fmt.Printf("verbose:                     %t\n", conf.Verbose)
	fmt.Printf("gzip_older_than:             %d\n", conf.Gzip_older_than)
	fmt.Printf("gzip_ext_blacklist_src:      %s\n", conf.Gzip_ext_blacklist_src)
	fmt.Printf("gzip_ext_blacklist:          %+v\n", conf.Gzip_ext_blacklist)
	fmt.Printf("periodic_minute:             %d\n", conf.Periodic_minute)
	fmt.Printf(" ----\n")
}

// ReadConfFile Read config from conf file
func ReadConfFile(conf *Config, json_path string) (read_ok bool) {
	_, fi_err := os.Stat(json_path)
	read_ok = false

	if nil == fi_err {
		// Read file
		f, f_err := ioutil.ReadFile(json_path)
		if nil == f_err {
			// Decode JSON
			json_err := json.Unmarshal(f, conf)
			if nil == json_err {
				conf.In_dir = filepath.ToSlash(conf.In_dir) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
				conf.Out_dir = filepath.ToSlash(conf.Out_dir)
				conf.Gzip_ext_blacklist = strings.Split(conf.Gzip_ext_blacklist_src, ",")
				SetConfDateTimeFormat(conf)
				SetSumSubFile(conf)
				read_ok = SetConfOutChannelSelected(conf)
				if read_ok {
					read_ok = SetOpeningHour(conf)
				}
			} else {
				fmt.Printf("BAD JSON config file: %s\n", json_err.Error())
			}
		}
	}

	return read_ok
}
