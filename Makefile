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

MAIN_FILE_PATH:=cmd/goscaffold/main.go
WORK_DIR := "/go/src/github.com/aperezg/goscaffold"
BUILD_DIR := bin
BINARY_NAME := goscaffold
BINARY_LINUX := ${BINARY_NAME}_linux

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
	go build -o ${BINARY_NAME} ${MAIN_FILE_PATH}

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
	go-bindata -pkg=data -o data/bindata.go doc

##cross compilation

##build-linux: Build the project to linux platforms
.PHONY: build-linux
build-linux: bindata
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR} ${MAIN_FILE_PATH}
##docker-build: Build the binary into a docker container
.PHONY: docker-build
docker-build: bindata
	docker build -t goscaffold_image .
	docker run --rm -it -v "${PWD}":${WORK_DIR} -w ${WORK_DIR} goscaffold_image go build -o $(BUILD_DIR)/${BINARY_LINUX} ${MAIN_FILE_PATH}
