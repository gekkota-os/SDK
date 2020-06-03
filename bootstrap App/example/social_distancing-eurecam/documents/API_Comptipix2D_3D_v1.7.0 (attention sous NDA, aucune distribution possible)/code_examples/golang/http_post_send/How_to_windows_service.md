# This document is for Windows user wanting to start http_post_send as a service

## Install the service thanks to [nssm](http://nssm.cc/)

NOTE: easier to install when there is no space in installation path (see nssm ["quoting issue"](http://nssm.cc/usage)) &rarr; Put http_post_send.exe in a directory path with no spaces   
NSSM usage : To use nssm in console first do **cd C:\path\to\the\downloaded\directory\containing\nssm.exe**

## Install service :

**nssm install http_post_send "C:\path\to\http_post_send.exe" -config_json -config_json_path "C:\path\to\http_post_send.json"**

- now in Windows menu *configuration panel/local service* you should see *http_post_send*

## UNinstall service :

**nssm remove http_post_send**
