package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterateMapblock(t *testing.T) {
	seen := map[[3]int]int{}
	var first, last [3]int
	n := 0

	IterateMapblock(func(x, y, z int) {
		if n == 0 {
			first = [3]int{x, y, z}
		}
		last = [3]int{x, y, z}
		n++
		seen[[3]int{x, y, z}]++
		assert.True(t, x >= 0 && x < 16 && y >= 0 && y < 16 && z >= 0 && z < 16)
	})

	assert.Equal(t, 4096, n)
	assert.Len(t, seen, 4096, "every position is visited exactly once")
	assert.Equal(t, [3]int{0, 0, 0}, first)
	assert.Equal(t, [3]int{15, 15, 15}, last)
}

func TestScanRecursive(t *testing.T) {
	root := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(root, "a", "b"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(root, "skipme"), 0755))
	for _, f := range []string{"top.txt", "a/one.txt", "a/b/two.txt", "skipme/three.txt"} {
		require.NoError(t, os.WriteFile(filepath.Join(root, f), []byte("x"), 0644))
	}

	folders, files := scan_recursive(root, []string{"skipme"})

	assert.ElementsMatch(t, []string{
		root,
		filepath.Join(root, "a"),
		filepath.Join(root, "a", "b"),
	}, folders)
	assert.ElementsMatch(t, []string{
		filepath.Join(root, "top.txt"),
		filepath.Join(root, "a", "one.txt"),
		filepath.Join(root, "a", "b", "two.txt"),
	}, files)
}

func TestScanRecursiveMissingDir(t *testing.T) {
	folders, files := scan_recursive(filepath.Join(t.TempDir(), "nope"), nil)
	assert.Empty(t, folders)
	assert.Empty(t, files)
}
