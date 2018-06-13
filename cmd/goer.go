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

	res,err := goer.Run(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}
