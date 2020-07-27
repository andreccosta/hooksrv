# Go parameters
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
BUILDDIR=bin
NAME=hooksrv

.PHONY: all
all: deps test build

.PHONY: build
build:
	$(GOBUILD) -o $(BUILDDIR)/$(NAME) -v

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f bin/*

.PHONY: run
run:
	$(GOBUILD) -o $(BUILDDIR)/$(NAME) -v
	./$(BUILDDIR)/$(NAME)

.PHONY: deps
deps:
	$(GOGET) -v -t -d ./...
