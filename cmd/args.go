package main

import "fmt"

type cmdArgs struct {
	cmd     string
	repo    string
	targets []string
}

func parse(args []string) (cmdArgs, error) {
	var c cmdArgs
	var err error
	if len(args) < 2 {
		err = fmt.Errorf("Invalid command args count")
	} else if len(args) == 2 {
		c.cmd = args[0]
		c.repo = args[1]
	} else {
		c.cmd = args[0]
		c.repo = args[1]
		c.targets = args[2:len(args)]
	}
	return c, err
}
