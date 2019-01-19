package dots

import (
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

const (
	exitOK = iota
	exitError
)

var r resolver

func Run(c cmd) int {
	r = &localResolver{
		repo: c.Repo,
	}
	ts, err := r.ReadYaml()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	if len(c.Targets) == 0 {
		// do all targets
		return do(ts)
	} else {
		// do specific targets
		tmap := map[string]target{}
		for _, ct := range c.Targets {
			if _, ok := tmap[ct]; ok {
				// already added
				continue
			}

			var tar *target
			for _, t := range ts {
				if t.Name == ct {
					tar = &t
					break
				}
			}
			if tar != nil {
				tmap[ct] = *tar
			} else {
				fmt.Fprintln(os.Stderr, fmt.Errorf("could not find target %s\n", ct))
			}
		}
		list := make([]target, 0, len(tmap))
		for _, t := range tmap {
			list = append(list, t)
		}
		return do(list)
	}
}

func do(ts []target) int {
	eg := errgroup.Group{}
	for _, t := range ts {
		eg.Go(func() error {
			return r.do(t)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}
	return exitOK
}
