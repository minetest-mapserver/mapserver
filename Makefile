
OUT_DIR=output
VERSION=git-$(shell git rev-parse HEAD)

all: builder_image $(OUT_DIR) $(MOD_ZIP)
	# build all with the docker image
	sudo docker run --rm -it\
	 -v $(shell pwd)/server/:/app\
	 -v mapserver-volume:/root/go\
	 -w /app\
	 mapserver-builder\
	 make test jshint all VERSION=$(VERSION)
	# copy generated files to output dir
	cp server/output/* $(OUT_DIR)/

builder_image:
	# build the docker image with all dependencies
	$(MAKE) -C docker_builder build

$(OUT_DIR):
	mkdir $@

clean:
	rm -rf $(OUT_DIR)
	$(MAKE) -C server clean
