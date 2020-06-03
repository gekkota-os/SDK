// To build you will need golang installed on your PC : https://golang.org/
// BUILD :
// go build -o http_demo
// RUN :
// ./http_demo
//
// BUILD for windows from linux :
// $ GOOS=windows GOARCH=386 go build -o http_demo.exe http_api_demo_get_sdcard_file_simple.go

// NOTE: this is a simple version just intended to demonstarte usage of golang http client to request ComptipixV3 API
// You have to manually change TEST value (user / pass / addr / file) and re-build

package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"time"
)

func main() {
	//TEST value
	user := "reader" 							// TEST value : change it to suit your need
	pass := "reader"							// TEST value : change it to suit your need
	addr := "http://192.168.0.141/"		// TEST value : change it to suit your need
	file := "20160525.csv"					// TEST value : change it to suit your need

	// Declare a netclient with 40 second timeout
	var netClient = &http.Client{
		Timeout: time.Second * 40, //NOTE: It's very important to set a default timeout that fit your application need (Golang default is no Timeout, we discourage to keep the default in production)
	}


	// Get a Token
	// ----------------

	// Send POST in octet-stream to get token
	str_cred := "user="+user+"&pass="+pass
	var credential = []byte(str_cred)
	req, err := http.NewRequest("POST", addr+"CONFIG?get_tkn", bytes.NewBuffer(credential))
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := netClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check token is okay
	if 200 != resp.StatusCode {
		if (401 == resp.StatusCode)||(418 == resp.StatusCode) {
			fmt.Println("Error Getting token : login/pass wrong")
		} else {
			fmt.Println("Error Getting token : ", resp.Status)
		}

		// If token is wrong no need to continue
		return;
	}

	body, _ := ioutil.ReadAll(resp.Body)
	token := string(body)
	resp.Body.Close()
	fmt.Println("Token OK:", token)



	// Check file exist
	// ----------------

	resp, err = http.Get(addr+"CONFIG?tkn="+token+"&sdcard_info="+file)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		if 404 == resp.StatusCode {
			fmt.Println("There is no file ", file)
		} else {
			fmt.Printf("Error Getting file %s, error code: %s", file, resp.Status)
		}

		// If file not ok no need to continue
		return;
	}

	// Get file info
	body, _ = ioutil.ReadAll(resp.Body)
	file_info := string(body)
	resp.Body.Close()
	fmt.Println("File info:", file_info)

	// Check file info OK
	//TODO


	// Get file
	// ----------------

	resp, err = http.Get(addr+"CONFIG?tkn="+token+"&sdcard_read="+file)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		fmt.Printf("Error reading file %s, error code: %s", file, resp.Status)

		// If file can't be read stop here
		return;
	}

	// Get file info
	body, _ = ioutil.ReadAll(resp.Body)
	file_read := string(body)
	resp.Body.Close()
	fmt.Println("File :", file_read)
}
