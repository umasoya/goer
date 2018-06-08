package main

import (
	"os"

	"github.com/umasoya/goer/pkg/goer"
)

func main() {
	exitCode := goer.Run()
	os.Exit(exitCode)
}
