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

echo $GO_PATH
# Change default binary name
DEFAULT=`grep module go.mod | rev | cut -f1 -d/ | rev`
echo $DEFAULT
mv -v $GO_PATH/bin/$DEFAULT $GO_PATH/bin/$BIN_NAME
