#!/bin/bash -e
cd $(dirname $0)
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

go test -v ./...
gofmt -s -w -l `find . -name \*.go`
