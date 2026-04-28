package settings

import (
	"mapserver/mapobjectdb"
	"strconv"
)

type DBSettings struct {
	db mapobjectdb.DBAccessor
}

func New(db mapobjectdb.DBAccessor) Settings {
	return &DBSettings{
		db: db,
	}
}

func (s *DBSettings) GetString(key string, defaultValue string) string {
	str, err := s.db.GetSetting(key, defaultValue)
	if err != nil {
		panic(err)
	}

	return str
}

func (s *DBSettings) SetString(key string, value string) {
	err := s.db.SetSetting(key, value)
	if err != nil {
		panic(err)
	}
}

func (s *DBSettings) GetInt(key string, defaultValue int) int {
	str, err := s.db.GetSetting(key, strconv.Itoa(defaultValue))
	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}

func (s *DBSettings) SetInt(key string, value int) {
	err := s.db.SetSetting(key, strconv.Itoa(value))
	if err != nil {
		panic(err)
	}
}

func (s *DBSettings) GetInt64(key string, defaultValue int64) int64 {
	str, err := s.db.GetSetting(key, strconv.FormatInt(defaultValue, 10))
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return value
}

func (s *DBSettings) SetInt64(key string, value int64) {
	err := s.db.SetSetting(key, strconv.FormatInt(value, 10))
	if err != nil {
		panic(err)
	}
}

func (s *DBSettings) GetBool(key string, defaultValue bool) bool {
	defStr := "false"
	if defaultValue {
		defStr = "true"
	}

	str, err := s.db.GetSetting(key, defStr)
	if err != nil {
		panic(err)
	}

	return str == "true"
}

func (s *DBSettings) SetBool(key string, value bool) {
	defStr := "false"
	if value {
		defStr = "true"
	}

	err := s.db.SetSetting(key, defStr)
	if err != nil {
		panic(err)
	}
}
