package dots

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"golang.org/x/sync/errgroup"
)

const (
	exitOK = iota
	exitError
)

var r Resolver

func SetResolver(res Resolver) {
	r = res
}

func Run(repo string, targets []string) int {
	if r == nil {
		r = NewGithubResolver(repo)
	}
	ts, err := r.Targets()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	if len(targets) == 0 {
		// do all targets
		return do(ts)
	} else {
		// do specific targets
		tmap := map[string]Target{}
		for _, ct := range targets {
			if _, ok := tmap[ct]; ok {
				// already added
				continue
			}

			var tar *Target
			for _, t := range ts {
				if t.Name == ct {
					tar = &t
					break
				}
			}
			if tar != nil {
				tmap[ct] = *tar
			} else {
				fmt.Fprintln(os.Stderr, fmt.Errorf("Could not find target %s", ct))
			}
		}
		list := make([]Target, 0, len(tmap))
		for _, t := range tmap {
			list = append(list, t)
		}
		return do(list)
	}
}

func do(ts []Target) int {
	eg := errgroup.Group{}
	for _, t := range ts {
		t := t
		eg.Go(func() error {
			reader, err := r.ReadFile(t)
			if err != nil {
				return err
			}
			defer reader.Close()

			dstPath, err := homedir.Expand(t.Dst)
			if err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(reader)
			fmt.Printf("write to %s\n", dstPath)
			return ioutil.WriteFile(dstPath, buf, 0644)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}
	return exitOK
}
