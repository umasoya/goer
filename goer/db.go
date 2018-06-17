package goer

import (
	"database/sql"
	"errors"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"gopkg.in/yaml.v2"
)

// Conn is DB connection configs
type Conn struct {
	Database string
	Host     string
	User     string
	Password string
	Port     string
}

// Column is struct of table column
type Column struct {
	Field   string
	Type    string
	Null    bool
	Key     string
	Extra   string
	Comment string
}

// Table is struct of table
type Table struct {
	Name   string
	Column []Column
}

func parseYaml(params Args) (Conn, error) {
	var conn Conn

	buf, err := ioutil.ReadFile(params.Args[0])
	if err != nil {
		return conn, err
	}

	err = yaml.Unmarshal(buf, &conn)
	return conn, err
}

func yamlToDsn(conn Conn) (string, error) {
	var addr string

	if len(conn.User) == 0 {
		err := errors.New("undefined user")
		return "", err
	}
	addr = conn.User

	if len(conn.Password) != 0 {
		addr += ":" + conn.Password
	}

	if conn.Host != "localhost" && len(conn.Host) != 0 {
		if len(conn.Port) == 0 {
			addr += "@tcp(" + conn.Host + ":3306)"
		} else {
			addr += "@tcp(" + conn.Host + ":" + conn.Port + ")"
		}
	} else {
		addr += "@"
	}

	if len(conn.Database) == 0 {
		err := errors.New("undefined database")
		return "", err
	}
	addr += "/" + conn.Database

	return addr, nil
}

// Run parseYaml,yamlToDsn and Analyze
func Run(params Args) ([]Table, error) {
	conn, err := parseYaml(params)
	if err != nil {
		return []Table{}, err
	}

	dsn, err := yamlToDsn(conn)
	if err != nil {
		return []Table{}, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return []Table{}, err
	}
	defer db.Close()

	return Analyze(db)
}
