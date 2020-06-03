There are some examples files :

NOTE that 4 golang examples are really usefull and complete working program :
--------------------
- sdcard_fetcher : get sdcard file using http
 → Compiled result for OSX/Linux/Windows here : https://eurecam.net/telechargement/API_demo/sdcard_fetcher_all.7z

- http_post_send : Comptipix http/POST protocol server implementation
 → Compiled result for OSX/Linux/Windows here : https://eurecam.net/telechargement/API_demo/http_post_send.7z

- parser : a counting file parser that can transform eurecam file in other format or remove double line, change resolution, rebuild an occupancy file from a counting file etc
 → Compiled result for OSX/Linux/Windows here : https://eurecam.net/telechargement/API_demo/parser_all.7z

- relay : a sever web that relay request from a web page to one or many comptipix : so web developper can use it
 → Compiled result for OSX/Linux/Windows here : https://eurecam.net/telechargement/API_demo/relay_all.7z


* other/http_dav.conf :
--------------------
An example of configuration file to set a working http/dav server using Apache (or any server that is compatible with apache config)
↳ This is a good start example if you want to test webdav file sending

* python/http_post_send.py :
--------------------
A simple working python program that can be used to receive files sent by Comptipix-V3 using POST protocol
Note that this program use python BaseHTTPServer which, at the time of writing, is not a production grade server : use it for debuging not for production !!
↳ This is a good start example if you want to test http-POST file sending protocol

* php/http_post_send.php :
--------------------
A simple working php program that can be used to receive files sent by Comptipix-V3 using POST protocol
↳ This is a good start example if you want to test http-POST file sending protocol

* php/http_post_send_full.php :
--------------------
A complet working php program that can be used to receive files sent by Comptipix-V3 using POST protocol, can skip already present files, produce log and use basic auth
↳ This is a demonstration program of http-POST file sending protocol with php

* php/get_real_time_counting.php :
--------------------
A working php program that can be used to request a Comptipix-V3 using the http API and save timestamped counting data :
This program do not download files saved on SD card, it uses the real time total entries and exits counters to save timestamped data into a file.
↳ This is a good start example if you want to get counting data in almost real time (ie: not depending on SD card's configuration)

* php/sdcard_fetch.php :
--------------------
A working php program that can be used to request a Comptipix-V3 using the http API to get sdcard files :
This program download files saved on SD card.
↳ This is a good start example if you want to get SDcard file using http request with php (if you are looking for a more robust program check the Go version: https://eurecam.net/telechargement/API_demo/sdcard_fetcher_all.7z)

* bash/get_sdcard_file.sh :
--------------------
A simple demo scripting curl program with bash.
Feel free to translate it in .bat for Windows (or use Windows10 bash implementation : https://msdn.microsoft.com/en-us/commandline/wsl/about) or use cygwin http://www.cygwin.com
↳ This is just a simple demo, not a robust program

* cpp/reboot_udp.cpp :
--------------------
A c++ program (for linux) that send udp command to order a sensor reboot


----------------------------------------
## NOTE on Golang (https://golang.org/) and Go examples :
----------------------------------------

NOTE: We think that GoLang fit very well if you want to query our HTTP API (or any HTTP API) :
- it has net/http support (https://golang.org/pkg/net/http/), so you have nothing to invent to do GET/POST requests
- if you want to query many sensor you can easily use goroutine to run many sensors quering concurrently (1 goroutine per sensor) **
** : query many sensor at one time (query 1000 sensors concurrently is perfectly doable with go routine),
but NEVER do many query to one sensor at the same time : he will not be able to answer to all requests

* golang/sdcard_fetch_simple :
--------------------
A simple example to query a sensor :
Get a token, use it to get file_info, if file_info is okay it get the file
↳ This is a good start example to write your own Golang program

* golang/sdcard_fetcher.go :
--------------------
Example that generate a usable program to get sensor sdcard files :
Same as get_file_simple_request.go, but can ask many files, and save them to a location.
This program has various option like transform file_name before saving, replace comma inside file etc
Use "-help" flag to check option availavable
↳ This is a demonstration program
→ The compiled result for OSX/Linux/Windows can be downloaded here : https://eurecam.net/telechargement/API_demo/sdcard_fetcher_all.7z

* golang/http_post_send.go :
--------------------
A complet working program that can be used to receive files sent by Comptipix-V3 using POST protocol, can skip already present files, produce log and use basic auth
This program has various option like transform file_name before saving, replace comma inside file etc
Use "-help" flag to check option availavable
↳ This is a demonstration program
→ The compiled result for OSX/Linux/Windows can be downloaded here : https://eurecam.net/telechargement/API_demo/http_post_send.7z

* golang/parser :
--------------------
Read a directory entries containing eurecam counting files and transform file in an output directory. Possibilities are :
- eurecam counting to eurecam counting :
Parser can transform an eurecam file to another eurecam file : changing resolution, apply opening hour
An Eurecam to Eurecam file transformation will appropriately sum line with same timestamp, fix gap with 0 data
- eurecam counting to eurecam occupancy :
Will generate an occupancy file from a counting file (it take care to not have an occupancy < 0)
- eurecam to other format :
Will transform eurecam counting file to another file type see golang/parser/generic_parser/doc-files/ to check available file transformation
↳ This is a demonstration program
→ The compiled result for OSX/Linux/Windows can be downloaded here : https://eurecam.net/telechargement/API_demo/parser_all.7z

* golang/relay :
--------------------
Relay is a usable demo of a server relaying request to 1 or many Comptipix :
- Relay server define a JSON format to query a relay server to do a query to another address, and is used to query Comptipix
- Relay serve static html + js present in relay_web/ thus using html + javascript developer can display result of many Comptipix in a unique html page.
	For example: it is possible to display occupancy sum of many Comptipix
- Web part of relay are design to be reusable (but this is not mandatory) :
	a developer can use app_relay.js (and app_request.js that is needed by app_relay.js) to get token and query many comptipix : demo.js is just a demo of using it
	a developer can write his own index.html + js to display everything the Comptipix API provide
- A developer can also redesign totaly all javascript file and use directly JSON request to relay server
- A developer can also redesign everything considering relay_serv.go just as a proof of concept
↳ This is a usable demonstration program, design to be extended or modified (especialy the web part)
→ The compiled result for OSX/Linux/Windows can be downloaded here : https://eurecam.net/telechargement/API_demo/relay_all.7z
