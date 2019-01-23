
STATIC_VFS=vfs/static.go
OUT_DIR=output

all: build

$(OUT_DIR):
	mkdir $@

test: $(STATIC_VFS) $(OUT_DIR)
	go test ./...

clean:
	rm -rf $(STATIC_VFS)
	rm -rf $(OUT_DIR)

$(STATIC_VFS):
	test -f ${HOME}/go/bin/esc || go get github.com/mjibson/esc
	${HOME}/go/bin/esc -o $@ -prefix="static/" -pkg vfs static

build: $(STATIC_VFS) $(OUT_DIR)
	GOOS=linux GOARCH=386 go build -o $(OUT_DIR)/mapserver-linux-x86
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-linux-x86_64
	GOOS=linux GOARCH=arm go build -o $(OUT_DIR)/mapserver-linux-arm
	GOOS=darwin GOARCH=386 go build -o $(OUT_DIR)/mapserver-darwin-x86
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-darwin-x86_64
	GOOS=windows GOARCH=386 go build -o $(OUT_DIR)/mapserver-windows-x86
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/mapserver-windows-x86_64

profile:
	go test -cpuprofile=cprof ./tilerenderer
	go tool pprof --text cprof
