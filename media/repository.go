package media

import (
	"io/ioutil"
	"os"
	"strings"
)

func ScanDir(repo map[string][]byte, path string, ignore []string) (int, error) {
	_, files := scan_recursive(path, ignore)
	size := 0

	for _, filename := range files {
		if strings.HasSuffix(filename, ".png") {

			file, err := os.Open(filename)

			if err != nil {
				return 0, err
			}

			content, err := ioutil.ReadAll(file)

			if err != nil {
				return 0, err
			}

			size += len(content)

			simplefilename := filename
			lastSlashIndex := strings.LastIndex(filename, "/")

			if lastSlashIndex >= 0 {
				simplefilename = filename[lastSlashIndex+1:]
			}

			repo[simplefilename] = content
		}

	}

	return size, nil
}
