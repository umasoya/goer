package goer

import (
	flags "github.com/jessevdk/go-flags"
)

// Options
type Options struct {
	Format	string	`short:"T" description:"Set output file format(default: pdf)" default:"pdf"`
	Outfile	string	`short:"o" description:"Write output to file outfile"`
}

type Args struct {
	Args []string
	Options
}

func Parse() (Args, error) {
	var opts Options

	args,err := flags.Parse(&opts)

	if err != nil {
		return Args{},err
	}

	params := Args{Args: args, Options : opts}
	return params,err
}
