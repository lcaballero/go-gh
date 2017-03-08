package conf

import (
	"errors"
	"path/filepath"
	"strings"
)

var ErrUnableToFindConventions = errors.New("error unable to find conventions used")

func CwdConventions(getwd func() (string, error)) (org, repo string, err error) {
	var dir string
	dir, err = getwd()
	if err != nil {
		return "", "", err
	}
	return Conventions(dir)
}

func Conventions(dir string) (org, repo string, err error) {
	if dir != "" {
		repo = filepath.Base(dir)
		dir = dir[:len(dir)-len(repo)]
		dir = strings.Trim(dir, " \t\r\n/")
	}

	if dir != "" {
		org = filepath.Base(dir)
		dir = dir[:len(dir)-len(org)]
		dir = strings.Trim(dir, " \t\r\n/")
	}

	if org == "" || repo == "" {
		return "", "", ErrUnableToFindConventions
	}

	return org, repo, nil
}
