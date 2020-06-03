#!/bin/bash

# set GOPATH for this project
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
CMD='export GOPATH="'$GOPATH:$DIR'"'
eval $CMD

# build go server
# NOTE: "-tags netgo" to force usage of go net package and not system network -> goal is to limit number of cpu used (wich is impossible if we use system networking : that create thread)
cd src/
if [ "$1" == "all" ] # rebuild all package (needed only if you updated package)
then
	go build -tags netgo -v -a -o ../http_post_send
elif [ "$1" == "prod" ] # Build all production version
then
	cd ..
	sh prod_build.sh
else
	go build -tags netgo -v -o ../http_post_send
fi
