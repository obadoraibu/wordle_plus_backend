GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=app

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/app

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build clean
