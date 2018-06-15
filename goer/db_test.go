package goer

import (
	"fmt"
	"unsafe"
	"testing"

	"gopkg.in/yaml.v2"
)

const (
	db   = "database: test_goer"
	host = "host: 127.0.0.1"
	user = "user: root"
	pw   = "password: secret"
	port = "port: 3306"
)

func TestParseYaml(t *testing.T) {
	var act Conn
	exp := Conn{Database: "test_goer", Host: "127.0.0.1", User: "root", Password: "secret", Port: "3306"}

	fstr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", db, host, user, pw, port)
	fbyte := *(*[]byte)(unsafe.Pointer(&fstr))

	err := yaml.Unmarshal(fbyte, &act)
	if err != nil {
		t.Errorf("failed yaml parse\n%#v", err)
	}

	if act != exp {
		t.Errorf("failed yaml parse %#v", err)
	}
}

func TestYamlToDsn(t *testing.T) {
	// checking err is nil only
	cases := []struct {
		data Conn
		dsn string
		err error
	} {
		{Conn{}, "", fmt.Errorf("")},
		{Conn{Database: "", Host: "128.0.0.1", User: "user", Password: "", Port: "3306"}, "", fmt.Errorf("")},
		{Conn{Database: "db", Host: "127.0.0.1", User: "user", Password: "", Port: "3306"}, "user@tcp(127.0.0.1:3306)/db", nil},
		{Conn{Database: "db", Host: "localhost", User: "user", Password: "", Port: ""}, "user@/db", nil},
		{Conn{Database: "db", Host: "", User: "user", Password: "", Port: ""}, "user@/db", nil},
		{Conn{Database: "db", Host: "localhost", User: "user", Password: "pass", Port: ""}, "user:pass@/db", nil},
	}

	for i,c := range cases {
		dsn,err := yamlToDsn(c.data)
		if c.dsn != dsn {
			t.Errorf("yamlToDsn(%v) want %+v but %+v", i, c.dsn, dsn)
		}
		if err == nil && c.err != nil {
			t.Errorf("yamlToDsn(%v) want %#v but %#v", i, c.err, err)
		}
	}
}
