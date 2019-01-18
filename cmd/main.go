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

func main() {
	cmd, err := dots.Parse(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitError)
	}
	fmt.Println(cmd)
}
