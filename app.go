package dots

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	homedir "github.com/mitchellh/go-homedir"
)

const (
	exitOK = iota
	exitError
)

var (
	r     Resolver
	force bool
)

func SetResolver(res Resolver) {
	r = res
}

func SetForce(f bool) {
	force = f
}

func Run(targets, tags []string) int {
	all, err := getTargets("")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	ts := filter(all, targets, tags)
	doTargets(ts)
	return exitOK
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

func doTargets(ts []Target) {
	wg := &sync.WaitGroup{}
	var dones struct {
		paths []string
		mux   sync.Mutex
	}
	var errs struct {
		errors []error
		mux    sync.Mutex
	}
	for _, t := range ts {
		wg.Add(1)
		go func(tar Target) {
			defer wg.Done()

			if dstPath, err := doTarget(tar); err != nil {
				errs.mux.Lock()
				errs.errors = append(errs.errors, err)
				errs.mux.Unlock()
			} else {
				dones.mux.Lock()
				dones.paths = append(dones.paths, dstPath)
				dones.mux.Unlock()
			}
		}(t)
	}
	wg.Wait()

	for _, p := range dones.paths {
		fmt.Printf("write to %s\n", p)
	}
	for _, err := range errs.errors {
		fmt.Fprintln(os.Stderr, err)
	}
}

func doTarget(t Target) (string, error) {
	reader, err := r.ReadFile(t.Sub, t.File)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	dstPath, err := homedir.Expand(t.Dst)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dstPath); err != nil {
		dstDir := filepath.Dir(dstPath)
		if _, err := os.Stat(dstDir); err != nil {
			err = os.MkdirAll(dstDir, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
	} else {
		if !force {
			return "", fmt.Errorf("already exists %s", dstPath)
		}
	}

	buf, err := ioutil.ReadAll(reader)

	return dstPath, ioutil.WriteFile(dstPath, buf, os.ModePerm)
}
