package sqlite



func (db *Sqlite3Accessor) GetSetting(key string, defaultvalue string) (string, error) {
  rows, err := db.db.Query(getSettingQuery, key)
	if err != nil {
		return "", err
	}

	defer rows.Close()

  value := defaultvalue

	if rows.Next() {

		err = rows.Scan(&value)
		if err != nil {
			return "", err
		}
	}

	return value, nil

}

func (db *Sqlite3Accessor) SetSetting(key string, value string) error {
  _, err := db.db.Exec(setSettingQuery, key, value)
	return err
}
