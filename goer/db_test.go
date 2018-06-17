package goer

import (
	"fmt"
	"os"
	"testing"
)

func TestParseYaml(t *testing.T) {
	exp := Conn{Database: "test_goer", Host: "127.0.0.1", User: "root", Password: "password", Port: "3306"}
	cases := []struct {
		params Args
		err    error
	}{
		{Args{Args: []string{os.Getenv("GOPATH") + "/src/github.com/umasoya/goer/sample/db_test.yaml"}, Options: Options{}}, nil},
		{Args{Args: []string{"./nothing.yaml"}, Options: Options{}}, fmt.Errorf("")},
	}

	for _, c := range cases {
		conn, err := parseYaml(c.params)
		if err == nil && conn != exp {
			t.Errorf("Failed parseYaml %#v,\n%#v\n%#v", conn, exp, err)
		}
		if err == nil && c.err != nil {
			t.Errorf("Failed parseYaml")
		}
	}

}

func TestYamlToDsn(t *testing.T) {
	// checking err is nil only
	cases := []struct {
		data Conn
		dsn  string
		err  error
	}{
		{Conn{}, "", fmt.Errorf("")},
		{Conn{Database: "", Host: "127.0.0.1", User: "user", Password: "", Port: "3306"}, "", fmt.Errorf("")},
		{Conn{Database: "db", Host: "127.0.0.1", User: "user", Password: "", Port: "3306"}, "user@tcp(127.0.0.1:3306)/db", nil},
		{Conn{Database: "db", Host: "localhost", User: "user", Password: "", Port: ""}, "user@/db", nil},
		{Conn{Database: "db", Host: "127.0.0.1", User: "user", Password: "", Port: ""}, "user@tcp(127.0.0.1:3306)/db", nil},
		{Conn{Database: "db", Host: "", User: "user", Password: "", Port: ""}, "user@/db", nil},
		{Conn{Database: "db", Host: "localhost", User: "user", Password: "pass", Port: ""}, "user:pass@/db", nil},
	}

	for i, c := range cases {
		dsn, err := yamlToDsn(c.data)
		if c.dsn != dsn {
			t.Errorf("yamlToDsn(%v) want %+v but %+v", i, c.dsn, dsn)
		}
		if err == nil && c.err != nil {
			t.Errorf("yamlToDsn(%v) want %#v but %#v", i, c.err, err)
		}
	}
}
