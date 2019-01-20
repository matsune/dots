package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/matsune/dots"
)

const (
	exitOK = iota
	exitError
)

const version = "1.0"

type options struct {
	Version bool     `short:"v" long:"version" description:"Show version"`
	Tags    []string `short:"t" long:"tag" description:"Tag of targets"`
}

var opts options

var parser = flags.NewParser(&opts, flags.Default)

func main() {
	parser.Usage = "[OPTIONS] REPO [TARGETS]"
	args, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(exitOK)
		} else {
			fmt.Fprint(os.Stderr, err)
			os.Exit(exitError)
		}
	}

	if opts.Version {
		fmt.Printf("Version %s\n", version)
		os.Exit(exitOK)
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "REPO is not passed")
		os.Exit(exitError)
	}

	repo := args[0]
	targets := args[1:len(args)]

	dots.SetResolver(dots.NewGithubResolver(repo))
	exit := dots.Run(targets, opts.Tags)
	os.Exit(exit)
}
