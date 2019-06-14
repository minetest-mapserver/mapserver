package media

import (
	"io/ioutil"
	"os"
	"strings"
)

func ScanDir(repo map[string][]byte, path string, ignore []string) error {
	_, files := scan_recursive(path, ignore)

	for _, filename := range files {
		if strings.HasSuffix(filename, ".png") {

			file, err := os.Open(filename)

			if err != nil {
				return err
			}

			content, err := ioutil.ReadAll(file)

			if err != nil {
				return err
			}

			simplefilename := filename
			lastSlashIndex := strings.LastIndex(filename, "/")

			if lastSlashIndex >= 0 {
				simplefilename = filename[lastSlashIndex+1:]
			}

			repo[simplefilename] = content
		}

	}

	return nil
}
