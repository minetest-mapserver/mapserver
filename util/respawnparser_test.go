package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const respawnJSON = `{
	"spawn": {"full_name": "Main spawn", "pos": {"x": 10, "y": -5, "z": 300}, "look": {"h": 1.5, "v": -0.5}, "color": "red", "icon": "home"},
	"minimal": {"pos": {"x": 1, "y": 2, "z": 3}}
}`

func TestParseRespawn(t *testing.T) {
	f, err := ParseRespawn([]byte(respawnJSON))
	require.NoError(t, err)
	require.Len(t, f, 2)

	spawn := f["spawn"]
	assert.Equal(t, "Main spawn", spawn.FullName)
	assert.Equal(t, RespawnPos{X: 10, Y: -5, Z: 300}, spawn.Pos)
	require.NotNil(t, spawn.Look)
	assert.Equal(t, 1.5, spawn.Look.H)
	assert.Equal(t, -0.5, spawn.Look.V)
	assert.Equal(t, "red", spawn.Color)
	assert.Equal(t, "home", spawn.Icon)

	minimal := f["minimal"]
	assert.Equal(t, RespawnPos{X: 1, Y: 2, Z: 3}, minimal.Pos)
	assert.Nil(t, minimal.Look)
	assert.Equal(t, "", minimal.FullName)
}

func TestParseRespawnInvalid(t *testing.T) {
	_, err := ParseRespawn([]byte(`{not json`))
	assert.Error(t, err)

	_, err = ParseRespawn([]byte(`[1,2,3]`))
	assert.Error(t, err)
}

func TestRespawnPosUnmarshalLenient(t *testing.T) {
	// wrong types and missing keys fall back to 0
	f, err := ParseRespawn([]byte(`{"a": {"pos": {"x": "str", "y": 4.9}}}`))
	require.NoError(t, err)
	assert.Equal(t, RespawnPos{X: 0, Y: 4, Z: 0}, f["a"].Pos)

	// pos is not an object
	_, err = ParseRespawn([]byte(`{"a": {"pos": "nope"}}`))
	assert.Error(t, err)
}

func TestGetInt(t *testing.T) {
	assert.Equal(t, 5, getInt(5.0))
	assert.Equal(t, -3, getInt(-3.7))
	assert.Equal(t, 0, getInt("5"))
	assert.Equal(t, 0, getInt(nil))
}

func TestParseRespawnFile(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "respawn.json")
	require.NoError(t, os.WriteFile(filename, []byte(respawnJSON), 0644))

	f, err := ParseRespawnFile(filename)
	require.NoError(t, err)
	assert.Len(t, f, 2)

	_, err = ParseRespawnFile(filepath.Join(t.TempDir(), "missing.json"))
	assert.Error(t, err)

	bad := filepath.Join(t.TempDir(), "bad.json")
	require.NoError(t, os.WriteFile(bad, []byte("garbage"), 0644))
	_, err = ParseRespawnFile(bad)
	assert.ErrorContains(t, err, "bad.json")
}
