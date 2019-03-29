
OUT_DIR=output
MOD_ZIP=$(OUT_DIR)/mapserver-mod.zip

all: $(OUT_DIR) $(MOD_ZIP)
	# build the docker image with all dependencies
	$(MAKE) -C docker_builder build
	# build all with the docker image
	sudo docker run --rm -it\
	 -v $(shell pwd)/server/:/app\
	 -v mapserver-volume:/root/go\
	 -w /app\
	 mapserver-builder\
	 make test all
	# copy generated files to output dir
	cp server/output/* $(OUT_DIR)/

$(OUT_DIR):
	mkdir $@

$(MOD_ZIP): $(OUT_DIR)
	zip -r $(OUT_DIR)/mapserver-mod.zip mapserver_mod

clean:
	rm -rf $(OUT_DIR)
	$(MAKE) -C server clean
