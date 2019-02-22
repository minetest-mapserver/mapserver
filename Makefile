
OUT_DIR=output
MOD_ZIP=$(OUT_DIR)/mapserver-mod.zip

all: $(OUT_DIR) $(MOD_ZIP)
	$(MAKE) -C server build-all-docker
	cp server/output/* $(OUT_DIR)/

$(OUT_DIR):
	mkdir $@

$(MOD_ZIP): $(OUT_DIR)
	zip -r $(OUT_DIR)/mapserver-mod.zip mapserver_mod

clean:
	rm -rf $(OUT_DIR)
	$(MAKE) -C server clean
