package shared

import (
	"bytes"
	"os/exec"
	"strings"
)

func CurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = "."
	out := bytes.NewBuffer([]byte{})
	cmd.Stdout = out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	branch := strings.Trim(out.String(), " \r\n\t")

	return branch, nil
}
