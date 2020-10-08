TARGETS           = darwin/amd64 linux/amd64 windows/amd64
DIST_DIRS         = find * -type d -exec

.NOTPARALLEL:

.PHONY: bootstrap build test test_fmt validate-copyright-headers fmt lint ci

ifdef DEBUG
GOFLAGS   := -gcflags="-N -l" -mod=vendor
else
GOFLAGS   := -mod=vendor
endif

# go option
GO              ?= go
TAGS            :=
LDFLAGS         :=
BINDIR          := $(CURDIR)/bin
PROJECT         := kubectl-capz
VERSION         ?= $(shell git rev-parse HEAD)
GITTAG          := $(shell git describe --exact-match --tags $(shell git log -n1 --pretty='%h') 2> /dev/null)
GOBIN           ?= $(shell $(GO) env GOPATH)/bin

ifdef DEBUG
LDFLAGS := -X main.version=$(VERSION)
else
LDFLAGS := -s -X main.version=$(VERSION)
endif

ifeq ($(OS),Windows_NT)
	EXTENSION = .exe
	SHELL     = cmd.exe
	CHECK     = where.exe
else
	EXTENSION =
	SHELL     = bash
	CHECK     = which
endif

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

all: build

.PHONY: build
build: go-build

.PHONY: go-build
go-build:
	$(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(PROJECT)$(EXTENSION)
