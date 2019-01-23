
STATIC_VFS=vfs/static.go

test: $(STATIC_VFS)
	go test ./...

clean:
	rm -rf $(STATIC_VFS)
	rm -rf mapserver

$(STATIC_VFS):
	go get github.com/mjibson/esc
	${HOME}/go/bin/esc -o $@ -prefix="static/" -pkg vfs static

build: $(STATIC_VFS)
	go build

profile:
	go test -cpuprofile=cprof ./tilerenderer
	go tool pprof --text cprof
