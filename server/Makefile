
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
	test -f ${HOME}/go/bin/esc || $(ENV) go get github.com/mjibson/esc
	${HOME}/go/bin/esc -o $@ -prefix="static/" -pkg vfs static

build: $(STATIC_VFS) $(OUT_DIR)
	$(ENV) GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/mapserver-linux-x86
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-linux-x86_64
	GOOS=linux GOARCH=arm go build -o $(OUT_DIR)/mapserver-linux-arm
	GOOS=darwin GOARCH=386 go build -o $(OUT_DIR)/mapserver-darwin-x86
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-darwin-x86_64
	GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/mapserver-windows-x86
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-windows-x86_64
