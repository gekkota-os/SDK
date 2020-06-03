#!/bin/bash

## This file set GOPATH to be here
## Just run it in the instalation directory (wich is "generic_parser")

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

CMD='export GOPATH="'$GOPATH:$DIR'"'
echo "$CMD"

##NOTE: that I don't know how to export var in bash : there is plenty of bad solution :
# http://stackoverflow.com/questions/16618071/can-i-export-a-variable-to-the-environment-from-a-bash-script-without-sourcing-i
