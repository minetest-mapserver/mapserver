package postgres

func (this *PostgresAccessor) intQuery(q string, params ...interface{}) int {
	rows, err := this.db.Query(q, params...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var result int
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}

		return result
	}

	panic("no result!")
}
