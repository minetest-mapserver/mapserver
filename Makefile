
STATIC_VFS=vfs/static.go

test: $(STATIC_VFS)
	go test ./...

$(STATIC_VFS):
	go get -u github.com/mjibson/esc
	${HOME}/go/bin/esc -o $@ -prefix="static/" -pkg vfs static

build: $(STATIC_VFS)
	go build
