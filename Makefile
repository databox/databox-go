all: test

#dependencies:
#	go get -u github.com/gorilla/mux
#	go get -u github.com/gorilla/pat


test:
	go test ./... -v -p=1 -gcflags "-N -l"

build:
	go build ./...
