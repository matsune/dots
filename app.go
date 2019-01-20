package dots

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func Run(targets, tags []string) int {
	all, err := getTargets("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	ts := filter(all, targets, tags)
	return doTargets(ts)
}

// Read dots.yml under sub directory and return targets.
// This method will be called recursively if sub directory has dots.yml.
func getTargets(sub string) ([]Target, error) {
	reader, err := r.ReadFile(sub, "dots.yml")
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	yml, err := ParseYaml(data)
	if err != nil {
		return nil, err
	}

	for i := range yml.Targets {
		yml.Targets[i].Sub = sub
	}

	res := yml.Targets
	for _, s := range yml.Sub {
		ts, err := getTargets(filepath.Join(sub, s))
		if err != nil {
			return nil, err
		}
		for _, t := range ts {
			res = append(res, t)
		}
	}
	return res, nil
}

func doTargets(ts []Target) int {
	eg := errgroup.Group{}
	for _, t := range ts {
		_t := t
		eg.Go(func() error {
			return doTarget(_t)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}
	return exitOK
}

func doTarget(t Target) error {
	reader, err := r.ReadFile(t.Sub, t.File)
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
}
