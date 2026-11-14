package tilerendererjob

import (
	"mapserver/settings"
	"mapserver/types"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// in-memory settings
type testSettings map[string]string

func (s testSettings) GetString(key string, defaultValue string) string {
	if v, ok := s[key]; ok {
		return v
	}
	return defaultValue
}
func (s testSettings) SetString(key string, value string) { s[key] = value }
func (s testSettings) GetInt(key string, defaultValue int) int {
	return int(s.GetInt64(key, int64(defaultValue)))
}
func (s testSettings) SetInt(key string, value int) { s.SetInt64(key, int64(value)) }
func (s testSettings) GetInt64(key string, defaultValue int64) int64 {
	v, err := strconv.ParseInt(s.GetString(key, ""), 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}
func (s testSettings) SetInt64(key string, value int64) { s[key] = strconv.FormatInt(value, 10) }
func (s testSettings) GetBool(key string, defaultValue bool) bool {
	return s.GetString(key, strconv.FormatBool(defaultValue)) == "true"
}
func (s testSettings) SetBool(key string, value bool) { s[key] = strconv.FormatBool(value) }

func TestLastMtimePos(t *testing.T) {
	s := testSettings{}

	// nothing stored yet (fresh or upgraded installation)
	assert.Nil(t, loadLastMtimePos(s, 0))
	s.SetInt64(settings.SETTING_LAST_MTIME, 1000)
	assert.Nil(t, loadLastMtimePos(s, 1000))

	// roundtrip, including negative coordinates
	saveLastMtimePos(s, 2000, types.NewMapBlockCoords(-5, 10, -2047))
	assert.Equal(t, int64(2000), s.GetInt64(settings.SETTING_LAST_MTIME, 0))
	assert.Equal(t, types.NewMapBlockCoords(-5, 10, -2047), loadLastMtimePos(s, 2000))

	// position belongs to another mtime ("last_mtime" changed manually): ignored
	assert.Nil(t, loadLastMtimePos(s, 1500))
	assert.Nil(t, loadLastMtimePos(s, 0))

	// no position: only the mtime is stored, the old position does not apply anymore
	saveLastMtimePos(s, 3000, nil)
	assert.Equal(t, int64(3000), s.GetInt64(settings.SETTING_LAST_MTIME, 0))
	assert.Nil(t, loadLastMtimePos(s, 3000))

	// garbage
	s.SetString(settings.SETTING_LAST_MTIME_POS, "3000/1/2")
	assert.Nil(t, loadLastMtimePos(s, 3000))
	s.SetString(settings.SETTING_LAST_MTIME_POS, "abc")
	assert.Nil(t, loadLastMtimePos(s, 3000))
}
