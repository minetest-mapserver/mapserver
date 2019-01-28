
STATIC_VFS=vfs/static.go
OUT_DIR=output
ENV=GO111MODULE=on

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
	#CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/mapserver-linux-x86
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-linux-x86_64
	# apt install gcc-5-arm-linux-gnueabihf, TODO: container
	GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc-5 CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" -o $(OUT_DIR)/mapserver-linux-arm
