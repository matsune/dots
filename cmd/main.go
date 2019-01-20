package main

import (
	"fmt"
	"os"

	"github.com/matsune/dots"
)

const (
	exitOK = iota
	exitError
)

func usage() {
	fmt.Print(`Usage:
	dots REPO [TARGET]

Help Options:
	-h, --help		Show this help message
`)
}

func main() {
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			usage()
			os.Exit(exitOK)
		}
	}

	c, err := parse(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		os.Exit(exitError)
	}

	// - FIXME: localResolver is for dev
	// dots.SetResolver(&localResolver{
	// 	repo: c.repo,
	// })

	exit := dots.Run(c.repo, c.targets)
	os.Exit(exit)
}
