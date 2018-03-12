.PHONY: test build
all: test

test:
	go test ./... -v -p=1 -gcflags "-N -l"

build:
	go build ./...

check:
	gofmt -s -w .
	go vet .
