
VERSION=2.0
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION}"

test:
	go test ./...

build:
	go build ${LDFLAGS}
