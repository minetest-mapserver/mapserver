package bundle

import (
	"encoding/json"
	"mapserver/vfs"
)

func getManifest(useLocal bool) *Manifest {
	manifestBytes := vfs.FSMustByte(useLocal, "/manifest.js")
	manifest := &Manifest{}
	err := json.Unmarshal(manifestBytes, manifest)
	if err != nil {
		panic(err)
	}

	return manifest
}

func createBundle(useLocal bool, files []string) []byte {
	script := make([]byte, 0)

	for _, name := range files {
		script = append(script, []byte("\n/*File: "+name+"*/\n")...)
		content, err := vfs.FSByte(useLocal, name)
		if err != nil {
			panic("vfs-file not found: " + name)
		}

		script = append(script, content...)
	}

	return script
}
