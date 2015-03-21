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
	go-linux-arm build -o build/piled

golang-buildsetup: golang-crosscompile
	source golang-crosscompile/crosscompile.bash; \
	go-crosscompile-build linux/arm

golang-crosscompile:
	git clone https://github.com/davecheney/golang-crosscompile.git


deb:
	mkdir -p build/root/usr/bin
	mkdir -p build/root/etc/init.d
	cp build/piled build/root/usr/bin
	cp init-script.sh build/root/etc/init.d/piled
	fpm -s dir -t deb -n piled -v $(VERSION) -p build/piled.deb \
		--deb-priority optional \
		--category util \
		--force \
		--iteration $(ITTERATION) \
		--deb-compression bzip2 \
		--url https://github.com/benschw/pi \
		--description "raspberry pi golang demo" \
		-m "Ben Schwartz <benschw@gmail.com>" \
		--license "Apache License 2.0" \
		--vendor "fliglio.com" \
		-a armhf \
		build/root/=/

push:
	scp build/piled.deb pi@192.168.0.115:/home/pi/

.PHONY: build
