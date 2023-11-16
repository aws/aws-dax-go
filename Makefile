MAKEFILE    := $(realpath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR    := $(shell dirname $(MAKEFILE))

SOURCES     := $(wildcard *.go */*.go) $(SRC_LINK) $(MAKEFILE)
REVISION    := $(shell git log -n 1 --pretty=format:%h -- $(SOURCES))
BUILD_FLAGS := -a -ldflags "-X revision=$(REVISION) -w -extldflags=$(LDFLAGS)" -tags "$(TAGS)"
__PKGS      = $(or $($PKG), $(shell go list ./... | grep -v "vendor" | grep -v "gopath"))
# Memoize PKGS so it is executed only once on-demand.
PKGS        = $(if $(__PKGS),,$(eval __PKGS := $$(__PKGS)))$(__PKGS)
# Allow target test packages to be overridden, e.g. $ make test TEST=TestBasic
TEST        := .

unexport GOBIN
export GO111MODULE=on

all: test 

target:
	mkdir -p $@

release-all: clean test 

$(SRC_LINK):
	mkdir -p $(shell dirname $(SRC_LINK))
	ln -sf $(ROOT_DIR) $(SRC_LINK)

vendor: 
	 go mod vendor

test: $(SOURCES) vendor
	SHELL=/bin/sh GOOS= go test  -coverprofile=coverageunit.out -covermode=atomic -v -tags "$(TAGS)" -run $(TEST) $(PKGS)

benchmark: $(SOURCES) vendor
	SHELL=/bin/sh GOOS= go test  -tags "$(TAGS)" \
		  -run=__XXX__ -bench=$(TEST) $(PKGS) -benchtime=1s -cpu 1

clean:
	rm -rf target

build: $(SOURCES) vendor
	go build -mod=vendor $(BUILD_FLAGS) $(PKGS) 

.PHONY: all release release-all test build clean benchmark
