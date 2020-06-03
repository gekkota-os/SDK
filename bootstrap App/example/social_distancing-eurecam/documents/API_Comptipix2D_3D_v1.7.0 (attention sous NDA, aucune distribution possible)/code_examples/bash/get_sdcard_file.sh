#!/bin/bash

# just some bash color
COLOR_DEFAULT="\e[39m"
COLOR_BLUE="\e[96m"
COLOR_GREEN="\e[92m"
COLOR_RED="\e[91m"
COLOR_YELLOW="\e[93m"

# address + file example (Change it to suit your need)
USER="reader"
PASS="reader"
ADDR="http://192.168.0.100"
NOW="$(date +'%Y%m%d')"
FILE="$NOW.csv"


################
## Get TOKEN
################

# address to get token
GET_TKN="$ADDR/CONFIG?get_tkn"

# check user/pass ok
USER_PASS_OK=$(curl --write-out %{http_code} --silent --output /dev/null -d "user=$USER&pass=$PASS" $GET_TKN)
if [ "200" != $USER_PASS_OK ]
then
	if [ "418" == $USER_PASS_OK ]
	then
		printf "$COLOR_RED→ Wrong user/pass ! HTTP error code is :$COLOR_DEFAULT $USER_PASS_OK\n"
	else
		printf "$COLOR_RED→ Wrong address... or something get wrong ! HTTP error code is :$COLOR_DEFAULT $USER_PASS_OK\n"
	fi
	exit
fi

# Get a token (we redo the same http request with curl)
TKN=$(curl -d "user=$USER&pass=$PASS" $GET_TKN)
printf "$COLOR_BLUE→ Token :$COLOR_GREEN $TKN\n\n $COLOR_DEFAULT"


################
## Get File info
################

# And now get 1 file info :
GET_FILE_INFO="$ADDR/CONFIG?tkn=$TKN&sdcard_info=$FILE"
RESULT=$(curl $GET_FILE_INFO)
printf "$COLOR_BLUE→ File info:$COLOR_GREEN $RESULT\n$COLOR_DEFAULT"

# check file info
if [ "sdcard_info=-1,," == $RESULT ]
then
	printf "$COLOR_RED↳ There is no file : $FILE :$COLOR_DEFAULT\n"
	exit
elif [ "error" == $RESULT ]
then
	printf "$COLOR_RED↳ SD card is in error !!!!!$COLOR_DEFAULT\n"
	exit
elif [ "sdcard_info=busy,," == $RESULT ]
then
	printf "$COLOR_YELLOW↳ SD card is busy !! $COLOR_DEFAULT→ Wait 5seconds before downloading file $FILE\n"
	sleep 5
fi
printf "\n"


################
## Get File
################

# And now get 1 file :
GET_FILE="$ADDR/CONFIG?tkn=$TKN&sdcard_read=$FILE"
RESULT=$(curl $GET_FILE)
printf "$COLOR_BLUE→ File$COLOR_DEFAULT $FILE :$COLOR_GREEN\n $RESULT\n$COLOR_DEFAULT"

# Save file
echo "... Saving file $FILE ..."
echo $RESULT > $FILE
echo "Done"
