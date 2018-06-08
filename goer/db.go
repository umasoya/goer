package goer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func parseYaml() {
}

func Connection() {
	conf := parseYaml()

	db, err := sql.Open("mysql", /* dataSourceName string */)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
