package goer

import "database/sql"

func getTables(db *sql.DB) (tables []Table, err error) {
	tables = []Table{}

	rows, err := db.Query(`SHOW TABLES`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err = rows.Scan(&table); err != nil {
			return
		}
		tables = append(tables, Table{Name: table})
	}
	return
}

// Analyze from Tables, Foreign Keys
func Analyze(db *sql.DB) (tables []Table, err error) {
	tables, err = getTables(db)
	if err != nil {
		return
	}

	return
}
