OUT_DIR=output
VERSION=git-$(shell git rev-parse HEAD)


# -ldflags="-X mapserver/app.Version=1.0"
GO_LDFLAGS=-ldflags "-linkmode external -extldflags -static -X mapserver/app.Version=$(VERSION)"
GO_LDFLAGS_WIN=-ldflags "-X mapserver/app.Version=$(VERSION)"
GO_BUILD=CGO_ENABLED=1 go build

BINARIES =  $(OUT_DIR)/mapserver-linux-x86_64
BINARIES += $(OUT_DIR)/mapserver-windows-x86-64.exe
BINARIES += $(OUT_DIR)/mapserver-linux-arm

JS_BUNDLE = public/js/bundle.js

all: $(JS_BUNDLE) $(OUT_DIR)/mapserver-linux-x86_64

$(OUT_DIR):
	mkdir $@

fmt:
	go fmt ./...

test: $(OUT_DIR)
	go vet ./...
	go test ./...

clean:
	rm -rf $(JS_BUNDLE) test-output
	rm -rf $(OUT_DIR)

jshint:
	cd public/js && jshint .

$(JS_BUNDLE):
	cd public/js && rollup -c rollup.config.js

$(OUT_DIR)/mapserver-linux-x86_64: $(OUT_DIR)
	# native (linux x86_64)
	GOOS=linux GOARCH=amd64 CC=gcc $(GO_BUILD) $(GO_LDFLAGS) -o $@

$(OUT_DIR)/mapserver-windows-x86-64.exe: $(OUT_DIR)
	GOARCH=amd64 GOOS=windows CC=x86_64-w64-mingw32-gcc $(GO_BUILD) $(GO_LDFLAGS_WIN) -o $@

$(OUT_DIR)/mapserver-linux-arm: $(OUT_DIR)
	# apt install gcc-5-arm-linux-gnueabihf
	GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc-8 $(GO_BUILD) $(GO_LDFLAGS) -o $@


release: builder_image $(OUT_DIR) $(MOD_ZIP)
	# build all with the docker image
	sudo docker run --rm -i \
	 -v $(shell pwd):/app \
	 -v mapserver-volume:/root/go \
	 -w /app \
	 mapserver-builder \
	 make test jshint release-all VERSION=$(VERSION)
	# copy generated files to output dir

builder_image:
	# build the docker image with all dependencies
	$(MAKE) -C docker_builder build

release-all: $(JS_BUNDLE) $(BINARIES)
