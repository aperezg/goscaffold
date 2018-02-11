all: help

##
##    ___   ___  __            __  __       _     _
##   / _ \ /___\/ _\ ___ __ _ / _|/ _| ___ | | __| |
##  / /_\///  //\ \ / __/ _` | |_| |_ / _ \| |/ _` |
## / /_\\/ \_// _\ \ (_| (_| |  _|  _| (_) | | (_| |
## \____/\___/  \__/\___\__,_|_| |_|  \___/|_|\__,_|
##
##
## @author: Adrian Perez <@>
## @description: Makefile for using goscaffold
##
##

MAIN_FILE_PATH:=cmd/goscaffold.go

.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

##install: Download the require dependencies for the project
.PHONY: install
install:
	go get github.com/jteeuwen/go-bindata/...
	go get gopkg.in/yaml.v2

##build: Compile on your own architecture
.PHONY: build
build: bindata
	go build -v ${MAIN_FILE_PATH}

##test: Execute the go test command
.PHONY: test
test:
	go test -v ./...

##clean: Execute go clean and remove the binaries
.PHONY: clean
clean:
	go clean

##bindata: Transform to binary data the require file for the application
.PHONY: bindata
bindata:
	go-bindata -pkg data -o data/bindata.go doc