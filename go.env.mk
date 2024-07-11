ifeq (, $(shell which go))
export GOROOT ?= /usr/lib/go
export GOPATH ?= $(HOME)/go
else
export GOROOT ?= $(shell go env GOROOT)
export GOPATH ?= $(shell go env GOPATH)
endif

export PATH := $(GOPATH)/bin:$(GOROOT)/bin:$(PATH)