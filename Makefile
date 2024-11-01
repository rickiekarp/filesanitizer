GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_PATH=build
BINARY_NAME=filesanitizer

test:
	$(GOTEST) -v ./...

releaseTag:
	git tag v$(shell date +1.%y%j.$(shell date -d "1970-01-01 UTC $(shell date +%T)" +%s))
	git push --tags