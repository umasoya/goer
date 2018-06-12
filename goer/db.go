package goer

import (
	"database/sql"
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	Database	string
	Host	string
	User	string
	Password	string
	Port	string
}

func parseYaml(params Args) (Conn, error) {
	var conn Conn

	buf,err := ioutil.ReadFile(params.Args[0])
	if err != nil {
		return conn,err
	}

	err = yaml.Unmarshal(buf, &conn)
	return conn,err
}

func yamlToDsn(conn Conn) (string, error) {
	var addr string

	if len(conn.User) == 0 {
		err := errors.New("undefined user")
		return addr, err
	}
	addr = conn.User

	if len(conn.Password) != 0 {
		addr += ":" + conn.Password
	}

	if conn.Host != "localhost" {
		addr += "@tcp(" + conn.Host + ":" + conn.Port + ")"
	} else {
		addr += "@"
	}

	if len(conn.Database) == 0 {
		err := errors.New("undefined database")
		return "", err
	}
	addr += "/" + conn.Database

	return addr,nil
}

func Connection(params Args) (*sql.DB,error) {
	var db *sql.DB

	conn,err := parseYaml(params)
	if err != nil {
		return db,err
	}

	dsn, err := yamlToDsn(conn)
	if err != nil {
		return db,err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db,err
	}
	defer db.Close()

	return db, nil
}
