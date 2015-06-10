all: test

#dependencies:
#	go get -u github.com/stretchr/testify

test:
	go test ./... -v -p=1 -gcflags "-N -l"

build:
	go build ./...
