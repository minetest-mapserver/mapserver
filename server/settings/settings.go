package settings

const (
	SETTING_LAST_MTIME       = "last_mtime"
	SETTING_LASTX            = "last_x"
	SETTING_LASTY            = "last_y"
	SETTING_LASTZ            = "last_z"
	SETTING_INITIAL_RUN      = "initial_run"
	SETTING_LEGACY_PROCESSED = "legacy_processed"
)

type Settings interface {
	GetString(key string, defaultValue string) string
	SetString(key string, value string)
	GetInt(key string, defaultValue int) int
	SetInt(key string, value int)
	GetInt64(key string, defaultValue int64) int64
	SetInt64(key string, value int64)
	GetBool(key string, defaultValue bool) bool
	SetBool(key string, value bool)
}
