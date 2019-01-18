package dots

import (
	"fmt"
)

type cmd struct {
	Self    string
	Repo    string
	Targets []string
}

func Parse(args []string) (cmd, error) {
	var c cmd
	var err error
	if len(args) < 2 {
		err = fmt.Errorf("Invalid command args count")
	} else if len(args) == 2 {
		c.Self = args[0]
		c.Repo = args[1]
	} else {
		c.Self = args[0]
		c.Repo = args[1]
		c.Targets = args[2:len(args)]
	}
	return c, err
}
