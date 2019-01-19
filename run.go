package dots

import (
	"fmt"
	"os"
)

const (
	exitOK = iota
	exitError
)

func Run(c cmd) int {
	l := localResolver{
		repo: c.Repo,
	}
	t, err := l.ReadYaml()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}
	fmt.Println(t)
	return exitOK
}
