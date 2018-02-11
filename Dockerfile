FROM golang:latest

RUN go get github.com/jteeuwen/go-bindata/...
RUN go get gopkg.in/yaml.v2