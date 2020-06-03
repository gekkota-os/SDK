// To build you will need golang installed on your PC : https://golang.org/
// * BUILD :
// go build
// * BUILD for windows from linux :
// * And other OS is also possible (see https://golang.org/doc/install/source#environment for other OS support) :
// GOOS=windows GOARCH=386 go build -o relay.exe relay.go
// * OR :
// GOOS=windows GOARCH=amd64 go build -o relay.exe relay.go
//
// * RUN :
// ./relay -help
//

package main

import (
	"log"
	"strings"
	"path/filepath"
	//"compress/gzip"
	"bytes"
	"encoding/json"
	"encoding/base64"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"time"
	"path"
	"flag"
	"sync"
	"reflect"
)

// Request base struct
// Method POST or GET to use on host
// HOST the host to relay request like '192.168.0.100' or 'fs-read' to read a file on relay server
// URL url parameter to use like '/CONFIG?uptime'
// Content content of answer or content to send for a POST request
// Content_type content type to set
type Req struct {
	Method			string	`json:"method"`
	Host				string	`json:"host"`
	Url				string	`json:"url"`
	Status			int		`json:"status"`
	Content			string	`json:"content"`
	Content_type	string	`json:"content_type"`
}

// A request cached
type ReqCached struct {
	Timestamp		int64
	Status			int
	Content			string
	Content_type	string
}

// Config struct
type Config struct {																			// DEFAULT:
	Port 								int 			`json:"port"`							// 8888
	Verbose							bool 			`json:"verbose"`						// false
	Verbose_cache					bool 			`json:"verbose_cache"`				// false
	Web_dir							string		`json:"web_dir"`						// "./relay_web/"
	Web_dir_default				string		`json:"web_dir_default"`			// "../../../relay_web/"
	No_web							bool			`json:"no_web"`						// false
	Cache_time_tkn					int64			`json:"cache_time_tkn"`				// 600
	Cache_time						int64			`json:"cache_time"`					// 2
	Cache_time_remove				int64 		`json:"cache_time_remove"`			// 1200
	Cache_strip_tkn				bool			`json:"cache_strip_tkn"`			// false
	Timeout							int			`json:"timeout"`						// 10
	Srv_read_timeout				int			`json:"srv_read_timeout"`			// 40
	Srv_write_timeout				int			`json:"srv_write_timeout"`			// 40
}


// Functions
// --------------------------

// Display config
func displayConfig(conf *Config) {
	// This is the same than a %+v but nicer
	fmt.Printf("CONFIG:\n ----\n")
	fmt.Printf("port:                     %d\n", conf.Port)
	fmt.Printf("verbose:                  %t\n", conf.Verbose)
	fmt.Printf("verbose_cache:            %t\n", conf.Verbose_cache)
	fmt.Printf("web_dir:                  %s\n", conf.Web_dir)
	fmt.Printf("web_dir_default:          %s\n", conf.Web_dir_default)
	fmt.Printf("no_web:                   %t\n", conf.No_web)
	fmt.Printf("cache_time_tkn:           %d\n", conf.Cache_time_tkn)
	fmt.Printf("cache_time:               %d\n", conf.Cache_time)
	fmt.Printf("cache_strip_tkn:          %t\n", conf.Cache_strip_tkn)
	fmt.Printf("timeout:                  %d\n", conf.Timeout)
	fmt.Printf("srv_read_timeout:         %d\n", conf.Srv_read_timeout)
	fmt.Printf("srv_write_timeout:        %d\n", conf.Srv_write_timeout)
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
				fmt.Printf("Export current config as relay.json.new OK\n")
			}
		} else {
			if (conf.Verbose) {
				fmt.Printf("Export current config as sdcard_ferelaytcher.json.new Error writing file !\n")
			}
			return export_conf_f_err
		}
	} else {
		if (conf.Verbose) {
			fmt.Printf("Export current config as relay.json.new Error generating JSON !\n")
		}
	}

	return export_conf_b_err
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
				conf.Web_dir = filepath.ToSlash(conf.Web_dir) // Transform path to be always '/' , this will be easier to transform any filepath (go transform all '/' to the OS separator)
				conf.Web_dir_default = filepath.ToSlash(conf.Web_dir_default)
				return true
			} else {
				fmt.Printf("BAD JSON config file: %s\n", json_err.Error())
			}
		}
	}
	return false
}

// Remove too old cache entries
func removeAllOldCache(conf *Config) {
	RequestCache.Range(func(key, value interface{}) (bool) {
		val := reflect.ValueOf(value)
		val_timestamp := val.FieldByName("Timestamp")
		if !val_timestamp.IsValid() {
			// Remove key because it's not valid
			RequestCache.Delete(key)

			if conf.Verbose_cache {
				fmt.Printf("removeAllOldCache() -> Remove invalid key : %s\n", key)
			}
		} else {
			cache_timestamp := val_timestamp.Int()
			now := time.Now().Unix()
			if ((now - conf.Cache_time_remove) > cache_timestamp) {
				// Remove key because it's too old
				RequestCache.Delete(key)

				if conf.Verbose_cache {
					fmt.Printf("removeAllOldCache() -> Remove too old key : %s\n", key)
				}
			}
		}

		return true
	})
}


// Do a relay request
// Set content and Status in request passed
// Manage requests cache
func doRelayReq(r *Req, conf *Config) (http_code int) {
	// Declare a http_client with a timeout
	var http_client = &http.Client{
		Timeout: time.Second * time.Duration(conf.Timeout), //NOTE: It's very important to set a default timeout that fit your application need (Golang default is no Timeout, we discourage to keep the default in production)
	}

	// Sanitize Host : add 'http://' if missing
	if (false == strings.HasPrefix(strings.ToLower(r.Host), "http://"))&&(false == strings.HasPrefix(strings.ToLower(r.Host), "https://")) {
		r.Host = "http://" + r.Host
	}

	// Check addr ends with '/' if url don't start with '/'
	if (false == strings.HasSuffix(r.Host, "/"))&&(false == strings.HasPrefix(r.Url, "/")) {
		r.Host += "/"
	}

	// Manage cache
	// ---

	req_is_get_tkn := false
	if strings.Contains(r.Url, "CONFIG?get_tkn") {
		req_is_get_tkn = true
	}

	cache_used := false
	cache_used_ok := false
	cache_key := ""
	if req_is_get_tkn && (conf.Cache_time_tkn > 0) {
		cache_used = true
	} else if !req_is_get_tkn && (conf.Cache_time > 0) {
		cache_used = true
	}

	// Use cache
	if cache_used {
		// Valid cache time in seconds
		var valid_time int64 = conf.Cache_time
		if req_is_get_tkn {
			valid_time = conf.Cache_time_tkn
		}

		// Strip tkn
		url := r.Url
		if conf.Cache_strip_tkn {
			i_tkn := strings.Index(url, "CONFIG?tkn=")
			i_and := strings.Index(url, "&")
			if (i_tkn > -1) && (i_and > -1) {
				url = strings.Replace(url, r.Url[i_tkn:i_and], "", 1)
			}
		}

		// Check for request result in cache
		cache_key = r.Method + r.Host + url + r.Content + r.Content_type
		cache_result, cache_ok := RequestCache.Load(cache_key)
		if cache_ok {
			cache_used_ok = true
			val := reflect.ValueOf(cache_result)
			val_timestamp := val.FieldByName("Timestamp")
			val_status := val.FieldByName("Status")
			val_content := val.FieldByName("Content")
			val_content_type := val.FieldByName("Content_type")

			if val_timestamp.IsValid() && val_status.IsValid() && val_content.IsValid() && val_content_type.IsValid() {
				cache_timestamp := val_timestamp.Int()
				now := time.Now().Unix()
				if ((now - valid_time) > cache_timestamp) {
					// Outdated cache = need to update it
					// Update the current timestamp to now,
					// to not have too many request trying to update the cache at the same time during the update request is running
					// which is what cache try to limit

					update_cache := ReqCached{Timestamp:now}
					update_cache.Status = r.Status
					update_cache.Content = r.Content
					update_cache.Content_type = r.Content_type

					RequestCache.Store(cache_key, update_cache) // Update cache to now, so this cache will be updated by one request at a time
				} else {
					// Directly get the cache info
					r.Status = int(val_status.Int())
					r.Content = val_content.String()
					r.Content_type = val_content_type.String()

					if conf.Verbose_cache {
						fmt.Printf("Cache Used for %s\n", cache_key)
					}

					// Stop here : cache is used
					return r.Status
				}
			} else {
				// This cache is incorrect
				if conf.Verbose_cache {
					fmt.Printf("Cache incorrect[%t %t %t %t] for %s\n", val_timestamp.IsValid(), val_status.IsValid(), val_content.IsValid(), val_content_type.IsValid(), cache_key)
				}

				RequestCache.Delete(cache_key)
				cache_ok = false
			}
		}
	}

	// Do request
	var req *http.Request
	var req_err error
	if "POST" == strings.ToUpper(r.Method) {
		req, req_err = http.NewRequest("POST", r.Host + r.Url, bytes.NewBuffer([]byte(r.Content)))
	} else if "GET" == strings.ToUpper(r.Method) {
		req, req_err = http.NewRequest("GET", r.Host + r.Url, nil)
	} else {
		return 400
	}

	// Set header
	if "" != r.Content_type {
		req.Header.Set("Content-Type", r.Content_type)
	}

	// Check for error of req creation answer
	if nil != req_err {
		log.Printf("ERR doRelayReq %s\n ", req_err)//DEBUG
		return 400
	}

	// Send POST or GET request
	resp, http_err := http_client.Do(req)
	if http_err != nil {
		return 400
	}
	defer resp.Body.Close()

	// fmt.Printf("-> %+v\n\n", resp)
	// fmt.Printf("--> %+v = %d\n\n", resp.Header["Content-Type"], len(resp.Header["Content-Type"]))

	// Read Content-Type
	if len(resp.Header["Content-Type"]) > 0 {
		r.Content_type = strings.Join(resp.Header["Content-Type"], ";")
	} else {
		r.Content_type = ""
	}

	// Read code
	r.Status = resp.StatusCode

	// Read answer
	body, _ := ioutil.ReadAll(resp.Body)
	if ("" == r.Content_type) || (strings.HasPrefix(r.Content_type, "image")) {
		r.Content = base64.StdEncoding.EncodeToString(body)	// Image or binary content
	} else {
		r.Content = string(body)	// Text content
	}
	resp.Body.Close()


	if cache_used {
		// Create/Update cache entry
		new_cache := ReqCached{Timestamp:time.Now().Unix()}
		new_cache.Status = r.Status
		new_cache.Content = r.Content
		new_cache.Content_type = r.Content_type

		if conf.Verbose_cache {
			if cache_used_ok {
				fmt.Printf("Update cache for %s\n", cache_key)
			} else {
				fmt.Printf("Set new cache for %s\n", cache_key)
			}
		}

		RequestCache.Store(cache_key, new_cache)
	}

	return 200
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

// Do a relay request that read file from File System host
// Set content and Status in request passed
func readFile(r *Req) (http_code int) {
	fi_size := GetFileSize(r.Url)
	if fi_size >= 0 {
		// Read file
		f, f_err := ioutil.ReadFile(r.Url)
		if f_err != nil {
			r.Status = 418 // Read error : send a teaPot !!
		}

		// Set result
		r.Content = string(f) // convert content to a 'string'
		r.Status = 200
	} else if fi_size == -1 {
		r.Status = 404 // File not found !!
	} else {
		r.Status = 401 // Right permission reading file
	}

	return r.Status
}

// Do a relay request that write file in File System host
// Set content and Status in request passed
func writeFile(r *Req, append bool) (http_code int) {
	// Create Directory and file if not exist
	os.MkdirAll(path.Dir(r.Url), 0777) // always check and create directory before write
	fi_size := GetFileSize(r.Url)
	var f *os.File
	var f_err error
	if fi_size < 0 {
		f, f_err = os.Create(r.Url)
	} else if append {
		f, f_err = os.OpenFile(r.Url, os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		f, f_err = os.OpenFile(r.Url, os.O_WRONLY, 0600)
	}

	// Check for create/open file error
	if f_err != nil {
		r.Status = 418 // create/open error : send a teaPot !!
	} else {
		// Write file
		defer f.Close()
		_, f_err = f.WriteString(r.Content)

		// Check for write error
		if f_err != nil {
			r.Status = 500 // Write error : send a 500  (because it's not normal to have error here)
		} else {
			r.Status = 200 // Write OK
		}
	}

	return r.Status
}

// Decode JSON and get answer
func processRequest(r io.ReadCloser, conf *Config) (http_code int, answ []byte) {
	// Decode JSON
	var req Req
	json_b, json_err := ioutil.ReadAll(r)
	json_err = json.Unmarshal(json_b, &req)
	if nil != json_err {
		log.Printf("BAD API req: %s \n", json_err.Error())
		return 400, []byte("ERR: Bad request format (API Decode err) !")
	}

	if (conf.Verbose) {
		fmt.Println("\n---------------")
		fmt.Println("Send request :")
		fmt.Printf("%+v \n", req)
	}

	if "fs-read" == req.Host {
		readFile(&req) 				// Read file from File System
	} else if "fs-write" == req.Host {
		writeFile(&req, false) 		// Write file on File System
	} else if "fs-write-append" == req.Host {
		writeFile(&req, true) 		// Write append to a file on File System
	} else {
		doRelayReq(&req, conf)	// Do a relay request
	}

	if (conf.Verbose) {
		fmt.Printf("\n↳ Answer :")
		fmt.Printf("%+v \n", req)
		fmt.Println("---------------")
	}

	// Build a JSON answer
	json_b, json_err = json.Marshal(req)
	if nil != json_err {
		return 500, []byte("ERR: Formating json")
	}

	return 200, []byte(json_b)
}


// A global request cache
// ----------------

// Used to store request result, so even if there is many client requesting 1 comptipix via relay,
// the relay will not send many requests to the comptipix (just send request in case of outdated cache or not present cache)
var RequestCache sync.Map

// Main server
// start a static file server and the relay request server
func main() {
	// Read arg
	// ----------------

	default_json_config_path := "./relay.json"
	def_web_dir := "./relay_web/"
	def_web_dir_default := "../../../relay_web/"	// When web_dir is not available (when user unzip soft webdir is at root and soft is in /linux/amd64 for exemple)
	help_demo := "This is a demo program distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE"

	// Config Flags
	Port := flag.Int("port", 8888, "Port to use to serve html+js content")
	Verbose := flag.Bool("verbose", false, "Be verbose on request (Default: Off)")
	Verbose_cache := flag.Bool("verbose_cache", false, "Be verbose on cached request (Default: Off)")
	Web_dir := flag.String("web_dir", def_web_dir, "Directory of static file to host")
	Web_dir_default := flag.String("web_dir_default",  def_web_dir_default, "Default directory of static file to host, if the provided one don't exist")		// Handy for user that just download the .exe to eurecam.net/telechargement/API_demo/ but don't copy relay_xeb/ in same dir as exe
	No_web := flag.Bool("no_web", false, "Run the relay server without web interface : relay server just relay request in JSON format (Default: enable web interface)")
	Timeout := flag.Int("timeout", 10, "Max time to wait on relayed request error (request relayed to Comptipix)")
	Cache_time_tkn := flag.Int64("cache_time_tkn", 600, "Time in second to keep a 'get_tkn' request in cache, -1 to disable it")
	Cache_time := flag.Int64("cache_time", 2, "Time in second to keep request in cache, -1 to disable it")
	Cache_time_remove := flag.Int64("cache_time_remove", 1200, "Time in second to remove element from cache as outdated, -1 to disable it")
	Cache_strip_tkn := flag.Bool("cache_strip_tkn", false, "Use it to strip the 'tkn=' part in cache (So all client request will use same cache)")
	Srv_read_timeout := flag.Int("srv_read_timeout", 40, "Web server read request timeout in seconds")
	Srv_write_timeout := flag.Int("srv_write_timeout", 40, "Web server write request timeout in seconds")
	// Config management flags
	Export_this_config_json := flag.Bool("export_this_config_json", false, "Export the current config as relay.json.new (Default don't export)")
	Config_json := flag.Bool("config_json", false, "Set this flag to read config from relay.json file (Default no read from relay.json)")
	Config_json_path := flag.String("config_json_path", default_json_config_path, "full path to file config (ex:'my/super/path/to/my_config.json')")
	help := flag.Bool("help_usage", false, "Print help usage and demo")
	flag.Parse()

	// Print help if user asked for it
	if *help {
		fmt.Printf("\nHELP\n")
		fmt.Printf("----------\n")
		fmt.Printf("see -h\n")

		fmt.Printf("\nUSAGE\n")
		fmt.Printf("----------\n")
		fmt.Printf("Run relay : ./relay, then open your browser and go to http://localhost:8888/ (change the port 8888 if you used -port option)\n")
		fmt.Printf("* Use the url loader command : If you want to :\n")
		fmt.Printf("- relay request to comptipix in 192.168.0.148                   ( → ?host=192.168.0.148 )\n")
		fmt.Printf("- display Occupancy                                             ( → &demo=2 ) \n")
		fmt.Printf("- user as reader and password is 'my_super_pass'                ( → &user=reader&pass=my_super_pass )\n")
		fmt.Printf("- hide buttons                                                  ( → &hide_buttons=1 )\n")
		fmt.Printf("- hide mouse                                                    ( → &hide_mouse=1 )\n")
		fmt.Printf("- hide eurecam logo                                             ( → &hide_eurecam=1 )\n")
		fmt.Printf("- hide hour                                                     ( → &hide_hour=1 )\n")
		fmt.Printf("- force gauge max to 55                                         ( → &occ_conf.thres_max=55 )\n")
		fmt.Printf("- force gauge min to 55, So gauge will work in reverse mode     ( → &occ_conf.thres_min=55 )\n")
		fmt.Printf("- force gauge min cliping to 1                                  ( → &occ_conf.min=1 )\n")
		fmt.Printf("- force gauge max cliping to 1001                               ( → &occ_conf.max=1001 )\n")
		fmt.Printf(" Open: http://localhost:8888/?host=192.168.0.148&demo=2&user=reader&pass=my_pass_for_reader&hide_buttons=1\n")
		fmt.Printf("* If you want same as above but plot file from SDcard :\n")
		fmt.Printf(" Open: http://localhost:8888/?host=192.168.0.148&demo=3&user=reader&pass=my_pass_for_reader&hide_buttons=1\n")
		fmt.Printf("* If you want same as above but display counting + hide mouse + hide eurecam logo + hide hour :\n")
		fmt.Printf(" Open: http://localhost:8888/?host=192.168.0.148&demo=4&user=reader&pass=my_pass_for_reader&hide_buttons=1&hide_mouse=1&hide_hour=1&hide_eurecam=1\n")
		fmt.Printf("* If you want to display anything from our API :\n")
		fmt.Printf(" Hack index.html + demo.js (you can re-use app_relay.js and app_request.js),\n or you can put your new file to another directory (must have an 'index.html') and use -web_dir to your code. \nWe wish you a good programming :)\n")

		fmt.Printf("\nLIMITATIONS:\n")
		fmt.Printf("----------\n")
		fmt.Printf("* %s\n", help_demo)
		return // if we print help we do nothing else
	}

	// Always show this message
	fmt.Printf("* %s\n", help_demo)

	// read config from JSON
	var conf Config
	if true == *Config_json {
		// Read config file
		read_json_ok := true
		_, fi_err := os.Stat(*Config_json_path)
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
		conf.Port = *Port
		conf.Verbose = *Verbose
		conf.Verbose_cache = *Verbose_cache
		conf.Web_dir = *Web_dir
		conf.Web_dir_default = *Web_dir_default
		conf.No_web = *No_web
		conf.Cache_time_tkn = *Cache_time_tkn
		conf.Cache_time = *Cache_time
		conf.Cache_time_remove = *Cache_time_remove
		conf.Cache_strip_tkn = *Cache_strip_tkn
		conf.Timeout = *Timeout
		conf.Srv_read_timeout = *Srv_read_timeout
		conf.Srv_write_timeout = *Srv_write_timeout
	}
	// Display config
	displayConfig(&conf)

	// Export current config
	if *Export_this_config_json {
		exportConfigJSON(&conf, "relay.json.new")
	}

	// Auto adapt web_dir
	// ----------------------

	// Use default dir if dir doesn't exist or don't contain an index.html
	// Only if user has not changed the default directory
	if def_web_dir == conf.Web_dir {
		file_test_1 := conf.Web_dir + "index.html"
		if _, err := os.Stat(file_test_1); os.IsNotExist(err) {
			// There is no index.html in this path, so use default path
			fmt.Printf("There is no index.html in %s, so use default path: %s as web_dir\n", conf.Web_dir, conf.Web_dir_default)
			conf.Web_dir = conf.Web_dir_default
		}
	}

	// Check the directory contains
	file_test_2 := conf.Web_dir + "index.html"
	use_web_interface := true
	if _, err := os.Stat(file_test_2); os.IsNotExist(err) {
		fmt.Printf("There is no index.html in %s\n", conf.Web_dir)
		use_web_interface = false
	}


	// Verbose log
	// ----------------------

	// We want to have the file and line when a log occur
	log.SetFlags(log.Lshortfile)


	// Every minutes, do some redundant task
	// ----------------------

	// Timer to call it every minute
	var f func()
	var t *time.Timer
	f = func () {
		// Remove to old cache entries (just to not have a cache that becomes too huge)
		removeAllOldCache(&conf)
		t = time.AfterFunc(time.Duration(1) * time.Minute, f)
	}
	t = time.AfterFunc(time.Duration(1) * time.Minute, f)
	defer t.Stop()


	// Public file serve
	// ----------------------

	if conf.No_web {
		use_web_interface = false
	}

	// Static file server (to serve login HTML + JS + CSS)
	h := http.NewServeMux() // Set a specific handler mux (it's better if you import other lib to not polute the default http handler, see: https://blog.cloudflare.com/exposing-go-on-the-internet/)
	if use_web_interface {
		h.Handle("/", http.FileServer(http.Dir(conf.Web_dir)))
	} else {
		fmt.Printf("Relay server run without web interface !!!\n")
	}


	// API route
	// ----------------------

	// Read api request
	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Decode URL
		_, query_err := url.QueryUnescape(r.URL.RawQuery)
		if nil != query_err {
			log.Printf("ERROR reading request: %s\n", query_err)
			http.Error(w, "Bad webapp query format", 400)
		} else {
			// Process request
			http_code, http_answ := processRequest(r.Body, &conf)
			if 200 == http_code {
				w.Write(http_answ)
			} else {
				http.Error(w, "ERROR Request", 400)
			}
		}
	})

	// Set a http server with timeout, do not use default httpserver (see: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
	port_str := fmt.Sprintf(":%d", conf.Port)
	srv := &http.Server{
		Addr: 			port_str,
		ReadTimeout: 	time.Duration(conf.Srv_read_timeout) * time.Second,
		WriteTimeout: 	time.Duration(conf.Srv_write_timeout) * time.Second,
		Handler:      	h,
	}


	// Start public static file server + API
	// ----------------------

	t_start := time.Now()
	fmt.Printf("%s → Relay server Up and running on port %s using plain/text http \n", t_start.Format("2006-01-02 15:04:05"), port_str)
	srv.ListenAndServe()
}
