
# Building the mapserver

Instructions to build the mapserver from source

## Build dependencies

* go >= 1.11

Ubuntu install: https://github.com/golang/go/wiki/Ubuntu

## Compile

```bash
# generate the static web files
go generate

# build the binary for the current playtform
go build

# (optionally) run the unit-tests
go test ./...

```
