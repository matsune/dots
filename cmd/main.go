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

	cmd, err := dots.Parse(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		os.Exit(exitError)
	}

	exit := dots.Run(cmd)
	os.Exit(exit)
}
