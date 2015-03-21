SHELL=/bin/bash

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
	rm -rf piled
	rm -rf golang-crosscompile

build: 
	source golang-crosscompile/crosscompile.bash; \
	go-linux-arm build -o piled

golang-buildsetup: golang-crosscompile
	source golang-crosscompile/crosscompile.bash; \
	go-crosscompile-build linux/arm

golang-crosscompile:
	git clone https://github.com/davecheney/golang-crosscompile.git

push:
	scp piled pi@192.168.0.115:/home/pi/