SHELL := /bin/bash
PATH := bin:$(PATH)
MAIN := main.go
PKG := $(go list ./... | grep -v /vendor/)

ifeq ($(RACE),1)
	GOFLAGS=-race
endif

VERSION?=$(shell git version > /dev/null 2>&1 && git describe --dirty=-dirty --always 2>/dev/null || echo NO_VERSION)
LDFLAGS=-ldflags "-X=main.version=$(VERSION)"

all: tools fix fmt vet rebuild

tools:
	@go get -u github.com/smartystreets/goconvey
	@go get -u honnef.co/go/staticcheck/cmd/staticcheck
	@go get -u github.com/rancher/trash

fix:
	@go get -u .

fmt:
	@gofmt -l -w -s `go list -f {{.Dir}}`

vet:
	@go vet $(PKG)
	@$(GOPATH)/bin/staticcheck

deps:
	@$(GOPATH)/bin/trash

rebuild:
	@go build -a $(LDFLAGS) $(GOFLAGS) $(MAIN)

build:
	@go build $(LDFLAGS) $(GOFLAGS) $(MAIN)

run:
	@echo "Compiling"
	@go run $(LDFLAGS) $(GOFLAGS) $(MAIN) -config=config.cfg -verbose

test:
	@go test $(LDFLAGS) $(GOFLAGS) $(PKG)

test-short:
	@go test $(LDFLAGS) $(GOFLAGS) -v -test.short -test.run="Test[^D][^B]" $(PKG) -verbose

convey:
	@$(GOPATH)/bin/goconvey