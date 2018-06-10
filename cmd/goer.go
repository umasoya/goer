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

	fmt.Printf("%v\n", params)
	fmt.Printf("%v\n", params.Args)
	fmt.Printf("%v\n", params.Options)
}
