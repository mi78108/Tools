# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
SRC_NAME=./src/transmit.go
BINARY_NAME=./bin/transmit
BINARY_UNIX=$(BINARY_NAME)_linux
BINARY_NT=$(BINARY_NAME)_windows
BINARY_DARWIN=$(BINARY_NAME)_mac

all: build-linux build-darwin build-nt
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC_NAME)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_DARWIN)
	rm -f $(BINARY_NT)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC_NAME)
	./$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(SRC_NAME)
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) -v $(SRC_NAME)
build-nt:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NT) -v $(SRC_NAME)
