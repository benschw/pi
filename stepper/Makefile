SHELL=/bin/bash
VERSION := 0.0.1
ITTERATION := $(shell date +%s)

# # drone build
# sudo apt-get update
# sudo apt-get install ruby-dev build-essential rubygems wget curl
# sudo gem install fpm
# make deps test build deb gzip

all: build

deps: golang-crosscompile golang-buildsetup
	go get -t -v ./...

test: 
	go test -v

clean:
	rm -rf build
	rm -rf piled
	rm -rf golang-crosscompile

build: 
	mkdir -p build
	source golang-crosscompile/crosscompile.bash; \
	go-linux-arm build -o stepper-demo

golang-buildsetup: golang-crosscompile
	source golang-crosscompile/crosscompile.bash; \
	go-crosscompile-build linux/arm

golang-crosscompile:
	git clone https://github.com/davecheney/golang-crosscompile.git



push:
	scp stepper-demo pi@192.168.0.116:/home/pi/

.PHONY: build
