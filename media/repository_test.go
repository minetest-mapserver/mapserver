package media

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	repo := make(map[string][]byte)

	ScanDir(repo, "../", []string{"mapserver.tiles", ".git"})

	for key := range repo {
		fmt.Println(key)
	}
}
