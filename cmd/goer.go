package main

import (
	"fmt"
	"log"

	"github.com/umasoya/goer/goer"
)

func main() {
	params, err := goer.Parse()
	if err != nil {
		log.Fatal(err)
	}

	db,err := goer.Connection(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", dsn)
}
