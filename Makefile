MAKEFILE    := $(realpath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR    := $(shell dirname $(MAKEFILE))
GOPATH      := $(ROOT_DIR)/gopath
SRC_LINK    := $(GOPATH)/src/github.com/aws/aws-dax-go

DEP_INSTALL_DIRECTORY := $(GOPATH)/bin
INSTALL_DIRECTORY     := $(DEP_INSTALL_DIRECTORY)
DEP_RELEASE_TAG       := v0.4.1
DEP_INSTALLER_URL     := "https://raw.githubusercontent.com/golang/dep/master/install.sh"

DEP_EXE     := $(DEP_INSTALL_DIRECTORY)/dep
DEP_TOML    := Gopkg.toml
DEP_LOCK    := Gopkg.lock
SOURCES     := $(wildcard *.go */*.go) $(SRC_LINK) $(DEP_LOCK) $(MAKEFILE)
REVISION    := $(shell git log -n 1 --pretty=format:%h -- $(SOURCES))
BUILD_FLAGS := -a -ldflags "-X revision=$(REVISION) -w -extldflags=$(LDFLAGS)" -tags "$(TAGS)"
__PKGS      = $(or $($PKG), $(shell cd $(SRC_LINK) && \
				env GOPATH=$(GOPATH) go list ./... | grep -v "vendor" | grep -v "gopath"))
# Memoize PKGS so it is executed only once on-demand.
PKGS        = $(if $(__PKGS),,$(eval __PKGS := $$(__PKGS)))$(__PKGS)
# Allow target test packages to be overridden, e.g. $ make test TEST=TestBasic
TEST        := .

export GOPATH
unexport GOBIN

all: test 

target:
	mkdir -p $@

release-all: clean test 

$(SRC_LINK):
	mkdir -p $(shell dirname $(SRC_LINK))
	ln -sf $(ROOT_DIR) $(SRC_LINK)

vendor: $(DEP_TOML) $(DEP_EXE)
	cd $(SRC_LINK) && $(DEP_EXE) ensure && touch $@

$(DEP_EXE):
	mkdir -p $(DEP_INSTALL_DIRECTORY)
	curl $(DEP_INSTALLER_URL) | sh

test: $(SOURCES) vendor
	SHELL=/bin/sh GOOS= go test -v -tags "$(TAGS)" -run $(TEST) $(PKGS)

benchmark: $(SOURCES) vendor
	SHELL=/bin/sh GOOS= go test  -tags "$(TAGS)" \
		  -run=__XXX__ -bench=$(TEST) $(PKGS) -benchtime=1s -cpu 1

clean:
	rm -rf target

build: $(SOURCES) vendor
	go build $(BUILD_FLAGS) $(PKGS) 

.PHONY: all release release-all test build clean benchmark
