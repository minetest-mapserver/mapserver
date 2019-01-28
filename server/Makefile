
STATIC_VFS=vfs/static.go
OUT_DIR=output
ENV=GO111MODULE=on

XGO = sudo docker run -it --rm --entrypoint /bin/bash -v "$(shell pwd):/mnt/build"
XGO += -e "http_proxy=$(shell echo ${http_proxy})"
XGO += -e "https_proxy=$(shell echo ${https_proxy})"
XGO += karalabe/xgo-1.11.1 -c

docker-build: $(STATIC_VFS)
	$(XGO) "cd /mnt/build && make build"

all: build

$(OUT_DIR):
	mkdir $@

test: $(STATIC_VFS) $(OUT_DIR)
	$(ENV) go test ./...

clean:
	rm -rf $(STATIC_VFS)
	rm -rf $(OUT_DIR)

$(STATIC_VFS):
	test -f esc || $(ENV) go get github.com/mjibson/esc
	${HOME}/go/bin/esc -o $@ -prefix="static/" -pkg vfs static

build: $(STATIC_VFS) $(OUT_DIR)
	CGO=1 GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/mapserver-linux-x86
	CGO=1 GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-linux-x86_64
	CGO=1 GOOS=linux GOARCH=arm go build -o $(OUT_DIR)/mapserver-linux-arm
	CGO=1 GOOS=darwin GOARCH=386 go build -o $(OUT_DIR)/mapserver-darwin-x86
	CGO=1 GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-darwin-x86_64
	CGO=1 GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/mapserver-windows-x86.exe
	CGO=1 GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-windows-x86_64.exe
