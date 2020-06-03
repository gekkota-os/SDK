package preq

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pconf"
	"putil"
	//"net/url"
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"
)

// Private Functions
// --------------------------

// Check a basic auth header
// true if auth is same as config, false otherwise
func checkAuth(r *http.Request, user string, pass string) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return (pair[0] == user) && (pair[1] == pass)
}

// Public Functions
// --------------------------

// Process http Request
// Create and update files according to sensor request
// Return answer as:
//  - http code (200 for success),
//  - answer string
// 		NOTE answer possible are :
// 		"2";	// skip all remaining files
// 		"1";	// skip this file
// 		"0"; 	// send this file
//  - error (nil for no error)
func ProcessRequest(r *http.Request, conf *pconf.Config) (int, []byte, error) {

	// Basic auth
	if ("" != conf.User) && (!checkAuth(r, conf.User, conf.Pass)) {
		basic_err := fmt.Errorf("Bad basic auth %s", r.URL.Query())
		if conf.Verbose {
			fmt.Printf("%s\n", basic_err.Error())
		}
		log.Printf("ERROR(401) - %s-%s : Unauthorized auth\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial")) // Log error
		return 401, []byte("Unauthorized auth"), basic_err
	}

	// Check path
	path_to_check := ""
	if "" != conf.Check_path {
		// get config path
		if conf.No_sanitize {
			path_to_check = conf.Check_path
		} else {
			path_to_check = putil.SanitizeURLPath(conf.Check_path)
		}

		// Check path
		if path_to_check != r.URL.Path { // NOTE: maybe could use url.PathUnescape, but url.PathUnescape need go version 1.8
			path_err := fmt.Errorf("Unauthorized path %s should match %s", r.URL.Path, path_to_check)
			if conf.Verbose {
				fmt.Printf("%s\n", path_err.Error())
			}
			log.Printf("ERROR(400) - %s-%s : Unauthorized path : %s, should match %s\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial"), r.URL.Path, path_to_check) // Log error
			return 401, []byte("Unauthorized path"), path_err
		}
	}

	// Read body
	body, body_err := ioutil.ReadAll(r.Body)
	if nil != body_err {
		fmt.Printf("ERROR reading body: %s\n", body_err.Error())
		log.Printf("ERROR(400) - %s-%s : Malformed body\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial")) // Log error
		return 400, []byte("Malformed body"), body_err
	}

	// Verbose info on request
	if conf.Verbose {
		// Print URL detail if verbose enabled
		fmt.Printf("\n\nRequest received: %s \n-----\n", r.URL.Query())
		fmt.Printf("Host: %s\n", r.Host)
		fmt.Printf("Path: %s\n", r.URL.Path)
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Headers: %+v\n", r.Header)
		fmt.Printf("\nAll: %+v\n-----\n\n", r)
		// fmt.Printf("Body: %s\n", body)
	}

	// Get query params
	_, serial_ok := r.URL.Query()["serial"]
	_, type_ok := r.URL.Query()["type"]
	_, file_ok := r.URL.Query()["file"]
	_, size_ok := r.URL.Query()["size"]
	_, check_ok := r.URL.Query()["check"]
	_, data_ok := r.URL.Query()["data"]
	_, get_ok := r.URL.Query()["get"]
	serial_val := r.URL.Query().Get("serial")
	tmp_type_val := r.URL.Query().Get("type")
	file_val := r.URL.Query().Get("file")
	tmp_size_val := r.URL.Query().Get("size")

	// Apply default type
	type_val := conf.Default_type
	if type_ok {
		type_val = tmp_type_val
	}

	// Answer to request
	answ_b := []byte("0") // 0 is send next file

	// Get file name
	file_path, _ := putil.GetFilePath(file_val, type_val, serial_val, r.URL.Path, conf) //TODO manage error
	file_full_path := file_path

	// Every request should have serial + file name
	if serial_ok && file_ok {
		file_full_path += putil.GetFileName(file_val, type_val, serial_val, conf)
	} else {
		req_err := fmt.Errorf("Missing serial or file info in query %s", r.URL.Query())
		if conf.Verbose {
			fmt.Printf("%s\n", req_err.Error())
		}
		log.Printf("ERROR(400) - %s-%s : Missing serial or file name info in request\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial")) // Log error
		return 400, []byte("Missing info"), req_err
	}

	// Check if file type is wanted
	wanted_file := true
	if !conf.Get_all {
		if !conf.Get_log && strings.HasPrefix(file_val, "log_") {
			wanted_file = false
		} else if !conf.Get_heat && strings.HasSuffix(file_val, "_heat.txt") {
			wanted_file = false
		} else if !conf.Get_depth && strings.HasSuffix(file_val, "_depth.bmp") {
			wanted_file = false
		} else if !conf.Get_track && strings.HasSuffix(file_val, "_track.csv") {
			wanted_file = false
		} else if !conf.Get_area && strings.HasSuffix(file_val, "_area.csv") {
			wanted_file = false
		} else if !conf.Get_occupancy && strings.HasSuffix(file_val, "_presence.csv") {
			wanted_file = false
		} else if !conf.Get_counting && strings.HasSuffix(file_val, ".csv") {
			wanted_file = false
		} else if !conf.Get_config && strings.HasSuffix(file_val, ".txt") {
			wanted_file = false
		} else if !conf.Get_jpg && strings.HasSuffix(file_val, ".jpg") {
			wanted_file = false
		} else if !conf.Get_line && strings.HasSuffix(file_val, "_line.csv") {
			wanted_file = false
		}
	}

	// Query type check : check file size
	if !wanted_file {
		answ_b = []byte("1") // Skip next file
		if conf.Verbose {
			fmt.Printf("Unwanted file: %s -> Skip (1)\n", file_val)
		}
	} else if check_ok && size_ok {
		// Convert size to int
		size_val, size_val_err := strconv.Atoi(tmp_size_val)
		if nil != size_val_err {
			req_size_err := fmt.Errorf("Wrong size (not an int) %s", tmp_size_val)
			if conf.Verbose {
				fmt.Printf("%s\n", req_size_err)
			}
			log.Printf("ERROR(400) - %s-%s : Invalid file size %d\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial"), size_val) // Log error
			return 400, []byte("Malformed size info"), req_size_err
		}

		// Get file size
		f_size := putil.GetFileSize(file_full_path)

		// Get size written in ".size" file
		if conf.Trust_size_file {
			f_trust, f_trust_err := ioutil.ReadFile(file_full_path + ".size")
			if nil != f_trust_err {
				log.Printf("ERROR - Cannot read size in %s : %s\n", file_full_path+".size", f_trust_err.Error()) // Log error
			} else {
				f_size, f_trust_err = strconv.ParseInt(string(f_trust[:]), 10, 64)
				if nil != f_trust_err {
					log.Printf("ERROR - Cannot convert read size value in %s : %s\n", file_full_path+".size", f_trust_err.Error()) // Log error
				} else {
					// fmt.Printf("Size from .size file = %d\n", f_size) //DEBUG
				}
			}
		}

		// Check size diff
		if (!conf.No_size_check) && (f_size == int64(size_val)) {
			if conf.Stop_size {
				answ_b = []byte("2") // Skip all remaining file
			} else {
				answ_b = []byte("1") // Skip next file
			}
		}
	}

	// Query type data : save file
	if data_ok {
		// Check directory
		if putil.GetFileSize(file_path) < 0 {
			os.MkdirAll(file_path, 0777)
		}

		// Create or update file
		f, f_err := os.Create(file_full_path)
		if nil != f_err {
			save_file_err := fmt.Errorf("Error saving file %s : check directory right", file_full_path)
			if conf.Verbose {
				fmt.Printf("Error %s : %s\n", f_err.Error(), save_file_err.Error())
			}
			log.Printf("ERROR(500) - %s-%s : saving file %s : check directory right\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial"), file_full_path) // Log error
			return 500, []byte("ERROR saving file"), save_file_err
		}
		defer f.Close()

		size_written := -1
		if ((0 != conf.Out_file_end_line) || conf.Replace_comma_by != "") && (!strings.HasSuffix(file_val, ".jpg")) && (!strings.HasSuffix(file_val, ".bmp")) { // Avoid messing jpg and bmp files
			file_content := putil.GetFileContent(file_val, string(body[:]), conf) // Transform file content according to config
			size_written = len([]byte(file_content))
			f.Write([]byte(file_content)) // Write transformed file
		} else {
			size_written = len(body)
			f.Write(body) // Write Vanilla file
		}
		f.Close()

		// If Generate size file enabled -> create the size file
		if conf.Generate_size_file {
			size_full_path := file_full_path + ".size"
			f, f_err = os.Create(size_full_path)
			if nil != f_err {
				size_file_err := fmt.Errorf("Error saving size file %s : check directory right", size_full_path)
				if conf.Verbose {
					fmt.Printf("Error %s : %s\n", f_err.Error(), size_file_err.Error())
				}
				log.Printf("ERROR(500) - %s-%s : saving SIZE file %s : check directory right\n", r.URL.Query().Get("type"), r.URL.Query().Get("serial"), file_full_path) // Log error
				return 500, []byte("ERROR saving size file"), size_file_err
			}
			f.Write([]byte(strconv.Itoa(size_written)))
			f.Close()
		}
	}

	// Query type get : read and return file content
	if get_ok && conf.Enable_get {
		if putil.GetFileSize(file_full_path) < 0 {
			return 404, []byte("File not found"), fmt.Errorf("Error geting file %s : check directory right, or file don't exist", file_full_path)
		} else {
			f_read_b, f_read_err := ioutil.ReadFile(file_full_path)
			if nil != f_read_err {
				return 500, []byte("Error reading file"), fmt.Errorf("Error reading file %s : check directory right", file_full_path)
			}
			answ_b = f_read_b
		}
	}

	// Verbose result
	if conf.Verbose && check_ok {
		fmt.Printf("Check file %s OK -> %s\n", file_full_path, string(answ_b[:]))
	} else if conf.Verbose && data_ok {
		fmt.Printf("Data(%d) file %s Received -> %s\n", len(body), file_full_path, string(answ_b[:]))
	}

	return 200, answ_b, nil
}
