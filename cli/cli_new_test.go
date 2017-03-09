package cli

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/urfave/cli.v2"
	"testing"
)

func Args(params ...string) []string {
	return params
}

func Test_New_Nil_Error(t *testing.T) {
	passed := false
	fn := func(c *cli.Context) error {
		passed = true
		return nil
	}

	err := New(Processing{BaseAction: fn}).Run(Args("go-gh"))
	assert.Nil(t, err)
	assert.True(t, passed)
}
