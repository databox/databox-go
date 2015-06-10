all: test

#checkin:
#	go test -run "TestCheckin*" ./client

test:
	go test ./... -v -p=1 -gcflags "-N -l"
build:
	go build ./...
