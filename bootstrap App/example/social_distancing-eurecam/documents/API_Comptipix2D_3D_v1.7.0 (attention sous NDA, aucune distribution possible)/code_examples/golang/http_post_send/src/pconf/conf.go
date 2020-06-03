package pconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pformat"
)

// Declaration
// --------------------------

// Config struct
type Config struct { // DEFAULT:
	Port                     int                    `json:"port"`                     // 8080
	Verbose                  bool                   `json:"verbose"`                  // false
	Out_dir                  string                 `json:"out_dir"`                  // "Data"
	Out_dir_add_path         bool                   `json:"out_dir_add_path"`         // false
	Stop_size                bool                   `json:"stop_size"`                // true
	Trust_size_file          bool                   `json:"trust_size_file"`          // false
	Generate_size_file       bool                   `json:"generate_size_file"`       // false
	Check_path               string                 `json:"check_path"`               // ""
	Default_type             string                 `json:"default_type"`             // "CPX3"
	Pass                     string                 `json:"pass"`                     // ""
	User                     string                 `json:"user"`                     // ""
	Enable_get               bool                   `json:"enable_get"`               // false
	Get_counting             bool                   `json:"get_counting"`             // false
	Get_occupancy            bool                   `json:"get_occupancy"`            // false
	Get_log                  bool                   `json:"get_log"`                  // false
	Get_jpg                  bool                   `json:"get_jpg"`                  // false
	Get_config               bool                   `json:"get_config"`               // false
	Get_line                 bool                   `json:"get_line"`                 // false
	Get_area                 bool                   `json:"get_area"`                 // false
	Get_track                bool                   `json:"get_track"`                // false
	Get_heat                 bool                   `json:"get_heat"`                 // false
	Get_depth                bool                   `json:"get_depth"`                // false
	Get_all                  bool                   `json:"get_all"`                  // true
	Log_suffix               bool                   `json:"log_suffix"`               // false
	Out_date_format_type     int                    `json:"out_date_format_type"`     // 0
	Out_date_format_sep      string                 `json:"out_date_format_sep"`      // ""
	Out_file_end_line        int                    `json:"out_file_end_line"`        // 0
	Replace_csv_by           string                 `json:"replace_csv_by"`           // ""
	Replace_txt_by           string                 `json:"replace_txt_by"`           // ""
	Replace_comma_by         string                 `json:"replace_comma_by"`         // ""
	No_replace_comma_replace bool                   `json:"no_replace_comma_replace"` // false
	Serial_prefix            bool                   `json:"serial_prefix"`            // false
	Serial_suffix            bool                   `json:"serial_suffix"`            // false
	Serial_suffix_final      bool                   `json:"serial_suffix_final"`      // false
	Save_dir_yyyy            bool                   `json:"save_dir_yyyy"`            // false
	Save_dir_yy              bool                   `json:"save_dir_yy"`              // false
	Save_dir_mm              bool                   `json:"save_dir_mm"`              // false
	Save_dir_dd              bool                   `json:"save_dir_dd"`              // false
	Save_dir_sep             string                 `json:"save_dir_sep"`             // "/"
	No_serial_dir            bool                   `json:"no_serial_dir"`            // false
	No_size_check            bool                   `json:"no_size_check"`            // false
	No_sanitize              bool                   `json:"no_sanitize"`              // false
	No_log                   bool                   `json:"no_log"`                   // false
	Srv_read_timeout         int                    `json:"srv_read_timeout"`         // 40
	Srv_write_timeout        int                    `json:"srv_write_timeout"`        // 40
	Date_time_format         pformat.DateTimeFormat `json:"date_time_format"`         // (internal) -> computed from Out_date_format_type/Out_date_format_sep/Out_time_format_type/Out_time_format_sep : date and time format type see \ref DateTimeFormat -> result of out_format + out_date_format_type...out_date_time_order
}

// Public Functions
// --------------------------

// Display config
func DisplayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("port:                     %d\n", conf.Port)
	fmt.Printf("verbose:                  %t\n", conf.Verbose)
	fmt.Printf("out_dir:                  %s\n", conf.Out_dir)
	fmt.Printf("out_dir_add_path:         %t\n", conf.Out_dir_add_path)
	fmt.Printf("stop_size:                %t\n", conf.Stop_size)
	fmt.Printf("trust_size_file:          %t\n", conf.Trust_size_file)
	fmt.Printf("generate_size_file:       %t\n", conf.Generate_size_file)
	fmt.Printf("check_path:               %s\n", conf.Check_path)
	fmt.Printf("default_type:             %s\n", conf.Default_type)
	fmt.Printf("user:                     %s\n", conf.User)
	fmt.Printf("pass:                     %s\n", conf.Pass)
	fmt.Printf("enable_get:               %t\n", conf.Enable_get)
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
	fmt.Printf("log_suffix:               %t\n", conf.Log_suffix)
	fmt.Printf("out_date_format_type:     %d -> %s\n", conf.Out_date_format_type, pformat.PrintDateFormatType(conf.Out_date_format_type))
	fmt.Printf("out_date_format_sep:      %s\n", conf.Out_date_format_sep)
	fmt.Printf("out_file_end_line:        %d -> %s\n", conf.Out_file_end_line, pformat.PrintFileEOL(conf.Out_file_end_line))
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
	fmt.Printf("srv_read_timeout:         %d\n", conf.Srv_read_timeout)
	fmt.Printf("srv_write_timeout:        %d\n", conf.Srv_write_timeout)
	fmt.Printf(" ----\n")
}

// Set user DateTime format override
func SetConfDateTimeFormat(conf *Config) {
	conf.Date_time_format.Date_format_type = pformat.DateFormatType(conf.Out_date_format_type)
	conf.Date_time_format.Date_format_sep = conf.Out_date_format_sep
	conf.Date_time_format.Time_format_type = pformat.TimeFormatType(0)     // We don't use time
	conf.Date_time_format.Time_format_sep = ""                             //conf.Out_time_format_sep	// We don't use time
	conf.Date_time_format.Date_time_sep = ""                               //conf.Out_date_time_sep		// We don't use time
	conf.Date_time_format.Date_time_order = pformat.DateTimeFormatOrder(0) // We don't use time
}

// Read config from conf file
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
				conf.Out_dir = filepath.ToSlash(conf.Out_dir) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
				SetConfDateTimeFormat(conf)
				read_ok = true
			} else {
				fmt.Printf("BAD JSON config file: %s\n", json_err.Error())
			}
		}
	}

	return read_ok
}

// Export current config
func ExportConfigJSON(conf *Config, file_name string) error {
	export_conf_b, export_conf_b_err := json.Marshal(conf)
	if nil == export_conf_b_err {
		export_conf_f, export_conf_f_err := os.Create(file_name)
		if nil == export_conf_f_err {
			export_conf_f.Write(export_conf_b)
			export_conf_f.Close()
			if conf.Verbose {
				fmt.Printf("Export current config as config.json.new OK\n")
			}
		} else {
			if conf.Verbose {
				fmt.Printf("Export current config as config.json.new Error writing file !\n")
			}
			return export_conf_f_err
		}
	} else {
		if conf.Verbose {
			fmt.Printf("Export current config as config.json.new Error generating JSON !\n")
		}
	}

	return export_conf_b_err
}
