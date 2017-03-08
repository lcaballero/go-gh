package conf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeWd(cwd string) func() (string, error) {
	return func() (string, error) {
		return cwd, nil
	}
}

func Test_Conventions_Root(t *testing.T) {
	org, repo, err := CwdConventions(makeWd("/"))

	assert.NotNil(t, err)
	assert.Equal(t, "", org, "org")
	assert.Equal(t, "", repo, "repo")
}

func Test_Conventions_1(t *testing.T) {
	org, repo, err := CwdConventions(makeWd("github.com/lcaballero/go-gh"))

	t.Logf("org: %s, repo: %s, err: %v", org, repo, err)

	assert.Nil(t, err)
	assert.Equal(t, "lcaballero", org)
	assert.Equal(t, "go-gh", repo)
}
