GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_PATH=build
BINARY_NAME=filesanitizer

test:
	$(GOTEST) -v ./...

check-tag:
ifndef TAG
	$(error TAG is undefined. Add TAG=v0.0.1 argument)
endif

release: check-tag
	git tag $(TAG)
	git push --tags