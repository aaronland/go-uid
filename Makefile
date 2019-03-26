CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/aaronland/go-uid
	cp *.go src/github.com/aaronland/go-uid/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/boltdb/bolt"
	@GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-artisanal-integers-proxy"
	@GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-string"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-pool-boltdb"
	mv src/github.com/aaronland/go-artisanal-integers-proxy/vendor/github.com/aaronland/* src/github.com/aaronland/
	mv src/github.com/aaronland/go-artisanal-integers-proxy/vendor/github.com/whosonfirst/* src/github.com/whosonfirst/
	rm -rf src/github.com/whosonfirst/go-whosonfirst-pool-boltdb/vendor/github.com/whosonfirst/go-whosonfirst-pool
	rm -rf src/github.com/whosonfirst/go-whosonfirst-pool-boltdb/vendor/github.com/boltdb
	cp -r vendor/* src/

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt *.go


bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/uid cmd/uid.go