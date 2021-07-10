#! /usr/bin/bash

SRC_DIR=$PWD
BIN_DIR=$SRC_DIR/bin
BIN_NAME=tcgrws

GO_PATH=`go env GOPATH`


if [ "$1" == "--debug" ]; then
    go build -o $BIN_DIR/$BIN_NAME
else
    go install -ldflags "-w -s"
fi


