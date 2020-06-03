//----------------------------------------------------------------------------------------------------------------------------------//
//			Eurecam Demo program receiving our POST protocol, fell free to use in your product.
//			This is a Demo program.
//			This program is distributed in the hope that it will be useful,
//			but WITHOUT ANY WARRANTY; without even the implied warranty of
//			MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//----------------------------------------------------------------------------------------------------------------------------------//

// This program is a full example of POST send protocol, managing some optional case as :
// - Allowing a get request to get file saved on server
// - Basic Auth (user/pass)
// - saving file at different location
// - change file name before saving it
// - transform file before saving it
// - loging to a log.log file

// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build -o http_post_send
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o http_post_send.exe http_post_send.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o http_post_send.exe http_post_send.go
//
// * RUN :
// ./http_post_send -help
//

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"pconf"
	"pformat"
	"preq"
	"time"
)

// Main
// --------------------------

// Main program
// read argument : address to get info
func main() {

	// Default
	default_json_config_path := "./http_post_send.json"

	// Read arg
	// ----------------

	help_demo := "This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE"
	Help_out_date_format_type := fmt.Sprintf("Date format (exemple for on Jan 2 2006 at 15:04:05, GMT-0700) (NOTE: set this to override file definition) :\n") + pformat.PrintAllDateFormatType()
	Help_file_end_line := fmt.Sprintf("File end line char (NOTE: set this to override file definition) :\n") + pformat.PrintAllFileEOL()

	// Read config
	Port := flag.Int("port", 8080, "Port to listen (Default: 8080)")
	Verbose := flag.Bool("verbose", false, "To be verbose on reception an response (Default: no verbose)")
	Out_dir := flag.String("out_dir", "Data/", "Directory to save files (Default: 'Data/')")
	Out_dir_add_path := flag.Bool("out_dir_add_path", false, "Add the sender path url to Out_dir directory (Default: don't add directory)")
	Stop_size := flag.Bool("stop_size", false, "Stop receiving other files when there is same size files")
	Trust_size_file := flag.Bool("trust_size_file", false, "Check the size info inside a file name having '.size' extension (and containing only the size in byte)")
	Generate_size_file := flag.Bool("generate_size_file", false, "generate a file with extension '.size' containing only the size in byte, this file can be used to check file size with option -trust_size_file")
	Check_path := flag.String("check_path", "", "Check path match a value")
	Pass := flag.String("pass", "", "Set pass value to use basic Auth (Default: no basic Auth)")
	User := flag.String("user", "", "Set user value to use basic Auth (Default: no basic Auth)")
	Default_type := flag.String("default_type", "CPX3", "Default type to use if it's not set in request. Note: only CPX3 version < 1.3.0 doesn't send the 'type' parameter (Default: 'CPX3)")
	Get_counting := flag.Bool("get_counting", false, "Get counting file (ex: '20170822.csv')")
	Get_occupancy := flag.Bool("get_occupancy", false, "Get occupancy file (ex: '20170822_presence.csv')")
	Get_log := flag.Bool("get_log", false, "Get log file (ex: 'log_20170822.csv')")
	Get_jpg := flag.Bool("get_jpg", false, "Get image file (ex: '20170822.jpg')")
	Get_config := flag.Bool("get_config", false, "Get config file (ex: '20170822.txt')")
	Get_line := flag.Bool("get_line", false, "Get line file (ex: '20170822_line.csv') (Comptipix-2D only)")
	Get_area := flag.Bool("get_area", false, "Get area	file (ex: '20170822_area.csv') (Comptipix-3D only)")
	Get_track := flag.Bool("get_track", false, "Get track file (ex: '20170822_track.csv') (Comptipix-3D only)")
	Get_heat := flag.Bool("get_heat", false, "Get heatmap file (ex: '20170822_heat.csv') (Comptipix-3D only)")
	Get_depth := flag.Bool("get_depth", false, "Get depthmap file (ex: '20170822_depth.bmp') (Comptipix-3D only)")
	Get_all := flag.Bool("get_all", true, "Get all files types (Default: Accept all files)")
	Enable_get := flag.Bool("enable_get", false, "Enable answer to 'get' request to return file content. Request must contains get flags + serial, type and file name (Default: Disabled)")
	Log_suffix := flag.Bool("log_suffix", false, "Transform log file to have 'log_' as suffix and no more as prefix (Default: don't transform log name)")
	Out_date_format_type := flag.Int("out_date_format_type", 0, Help_out_date_format_type)
	Out_date_format_sep := flag.String("out_date_format_sep", "", "Output date format separator\n \t ex: for Mon Jan 2 2006 : -out_date_format_type=2 -out_date_format_sep='-' -> 2006-01-02 (NOTE: set this to override file definition)")
	Out_file_end_line := flag.Int("out_file_end_line", 0, Help_file_end_line)
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
	No_log := flag.Bool("no_log", false, "Don't log error in a file log.log (Default: log enabled)")
	Srv_read_timeout := flag.Int("srv_read_timeout", 40, "Web server read request timeout in seconds (Default: 40)")
	Srv_write_timeout := flag.Int("srv_write_timeout", 40, "Web server write request timeout in seconds (Default: 40)")
	// Read flags
	export_this_config_json := flag.Bool("export_this_config_json", false, "Export the current config as http_post_send.json.new (Default don't export)")
	config_json := flag.Bool("config_json", false, "Set this flag to read config from http_post_send.json file (Default no read from http_post_send.json)")
	config_json_path := flag.String("config_json_path", default_json_config_path, "Full path to file config including his name (ex:'my/super/path/to/my_config.json')")
	flag_help := flag.Bool("help_usage", false, "Print help") // Override default help to add tips and usage
	flag.Parse()

	// Print help if user asked for it
	if *flag_help != false {
		fmt.Println("Http POST file receiver")
		fmt.Println("This program implement the file HTTP post reception with Comptipix-V3 protocol")
		fmt.Printf("-h → See all options\n")

		fmt.Printf("\nUSAGE EXAMPLES:\n")
		fmt.Printf("----------\n")
		fmt.Printf("Store files in 'type-serial' directory, stop sending as soon as a file sent has same size\n")
		fmt.Printf(" ./http_post_send -stop_size\n")
		fmt.Printf("Same as preceding + use basic auth user='foo' and password='bar' + check that query path match 'hello'\n")
		fmt.Printf(" ./http_post_send -stop_size -user foo -pass bar -check_path hello\n")
		fmt.Printf("Do not store files in 'type-serial' directory + transform file date as yymmdd + replace comma inside file by semicolon\n")
		fmt.Printf(" ./http_post_send -no_serial_dir -yyyymmdd_to_yymmdd -serial_prefix -replace_comma_by ';'\n")
		fmt.Printf("Store files in 'Type-Serial/yyyy/mm' directory\n")
		fmt.Printf(" ./http_post_send -save_dir_yyyy -save_dir_mm\n")
		fmt.Printf("Store files in 'yy-mm' directory and set it serial prefix\n")
		fmt.Printf(" ./http_post_send -no_serial_dir -serial_prefix -save_dir_yy -save_dir_mm -save_dir_sep '-'\n")
		fmt.Printf("Store only counting and log files even if sensor is sending all files type\n")
		fmt.Printf(" ./http_post_send -fcounting -flog\n")

		fmt.Printf("\nTIPS:\n")
		fmt.Printf("----------\n")
		fmt.Printf("To run this on Linux listening on port <= 3000, you can run (as root) once before use :\n")
		fmt.Printf("↳ setcap 'cap_net_bind_service=+ep' http_post_send\n")

		fmt.Printf("\nLIMITATIONS:\n")
		fmt.Printf("----------\n")
		fmt.Printf("%s\n", help_demo)
		return // if we print help, we do nothing else
	}
	fmt.Printf("NOTE: %s\n", help_demo) // diplay demo notice at each run

	// Read config
	// ----------------

	// Set default files to get : by default accept all files, but if some are enabled disable the -get_all flags
	if *Get_counting || *Get_occupancy || *Get_log || *Get_jpg || *Get_config || *Get_area || *Get_track || *Get_heat || *Get_depth {
		*Get_all = false
	}

	// read config from JSON
	var conf pconf.Config
	if (true == *config_json) || (default_json_config_path != *config_json_path) {
		// Read config file
		read_json_ok := pconf.ReadConfFile(&conf, *config_json_path)
		if !read_json_ok {
			fmt.Printf("Error reading json file ABORT\n")
			return
		}
	} else {
		// Because we have to do it after flag.Parse()
		conf.Port = *Port
		conf.Verbose = *Verbose
		conf.Out_dir = *Out_dir
		conf.Out_dir_add_path = *Out_dir_add_path
		conf.Stop_size = *Stop_size
		conf.Trust_size_file = *Trust_size_file
		conf.Generate_size_file = *Generate_size_file
		conf.Check_path = *Check_path
		conf.Pass = *Pass
		conf.User = *User
		conf.Default_type = *Default_type
		conf.Get_counting = *Get_counting
		conf.Get_occupancy = *Get_occupancy
		conf.Get_log = *Get_log
		conf.Get_jpg = *Get_jpg
		conf.Get_config = *Get_config
		conf.Get_line = *Get_line
		conf.Get_area = *Get_area
		conf.Get_track = *Get_track
		conf.Get_heat = *Get_heat
		conf.Get_depth = *Get_depth
		conf.Get_all = *Get_all
		conf.Enable_get = *Enable_get
		conf.Log_suffix = *Log_suffix
		conf.Out_date_format_type = *Out_date_format_type
		conf.Out_date_format_sep = *Out_date_format_sep
		conf.Out_file_end_line = *Out_file_end_line
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
		conf.No_log = *No_log
		conf.Srv_read_timeout = *Srv_read_timeout
		conf.Srv_write_timeout = *Srv_write_timeout
	}
	// Display config
	pconf.DisplayConfig(&conf)

	// Export current config
	if *export_this_config_json {
		pconf.ExportConfigJSON(&conf, "http_post_send.json.new")
	}

	// Set log file
	// ----------------

	if !conf.No_log {
		f, f_err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if f_err != nil {
			fmt.Printf("Error opening log file: %v !!!\n", f_err)
		} else {
			defer f.Close()
			log.SetOutput(f)
			log.Printf("Server Up and running on port %d using plain/text http !!\n---\n", conf.Port)
		}
	}

	// Handle http
	// ----------------

	// Set a http
	h := http.NewServeMux() // Set a specific handler mux (it's better if you import other lib to not polute the default http handler, see: https://blog.cloudflare.com/exposing-go-on-the-internet/)
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get answer
		answ_code, answ_byte, answ_err := preq.ProcessRequest(r, &conf)

		// Send answer
		if (200 == answ_code) && (nil == answ_err) {
			w.Write(answ_byte) // OK (request command understood and proccessed)
		} else {
			http.Error(w, string(answ_byte[:]), answ_code) // Error (request command not understood, or wrong data in request)
		}
	})

	// Set a http server with timeout, do not use default httpserver (see: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
	port_str := fmt.Sprintf(":%d", conf.Port)
	srv := &http.Server{
		Addr:         port_str,
		ReadTimeout:  time.Duration(conf.Srv_read_timeout) * time.Second,
		WriteTimeout: time.Duration(conf.Srv_write_timeout) * time.Second,
		Handler:      h,
	}

	fmt.Printf("%s\n\nServer Up and running on port %s using plain/text http !! \n \n", help_demo, port_str)
	srv.ListenAndServe()
}
