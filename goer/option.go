package goer

import (
	"os"
	"log"

	flags "github.com/jessevdk/go-flags"
)

// Options
type Options struct {
	Format	string	`short:"T" description:"Set output file format(default: pdf)" default:"pdf"`
	Outfile	string	`short:"o" description:"Write output to file outfile"`
}

var opts Options

func Parse() ([]string, Options) {
	args, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return args,opts
}
