package settings

import (
	"mapserver/mapobjectdb"
	"strconv"
)

const (
	SETTING_LAST_MTIME       = "last_mtime"
	SETTING_LASTX            = "last_x"
	SETTING_LASTY            = "last_y"
	SETTING_LASTZ            = "last_z"
	SETTING_INITIAL_RUN      = "initial_run"
	SETTING_LEGACY_PROCESSED = "legacy_processed"
)

type Settings struct {
	db mapobjectdb.DBAccessor
}

func New(db mapobjectdb.DBAccessor) *Settings {
	return &Settings{
		db: db,
	}
}

func (this *Settings) GetString(key string, defaultValue string) string {
	str, err := this.db.GetSetting(key, defaultValue)
	if err != nil {
		panic(err)
	}

	return str
}

func (this *Settings) SetString(key string, value string) {
	err := this.db.SetSetting(key, value)
	if err != nil {
		panic(err)
	}
}

func (this *Settings) GetInt(key string, defaultValue int) int {
	str, err := this.db.GetSetting(key, strconv.Itoa(defaultValue))
	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}

func (this *Settings) SetInt(key string, value int) {
	err := this.db.SetSetting(key, strconv.Itoa(value))
	if err != nil {
		panic(err)
	}
}

func (this *Settings) GetInt64(key string, defaultValue int64) int64 {
	str, err := this.db.GetSetting(key, strconv.FormatInt(defaultValue, 10))
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return value
}

func (this *Settings) SetInt64(key string, value int64) {
	err := this.db.SetSetting(key, strconv.FormatInt(value, 10))
	if err != nil {
		panic(err)
	}
}

func (this *Settings) GetBool(key string, defaultValue bool) bool {
	defStr := "false"
	if defaultValue {
		defStr = "true"
	}

	str, err := this.db.GetSetting(key, defStr)
	if err != nil {
		panic(err)
	}

	return str == "true"
}

func (this *Settings) SetBool(key string, value bool) {
	defStr := "false"
	if value {
		defStr = "true"
	}

	err := this.db.SetSetting(key, defStr)
	if err != nil {
		panic(err)
	}
}
