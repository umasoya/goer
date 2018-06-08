package main

import (
	"fmt"

	"github.com/umasoya/goer/goer"
)

func main() {
	args, opts := goer.Parse()

	fmt.Printf("%v\n", opts.Outfile)
	fmt.Printf("%#v\n", args)
}
