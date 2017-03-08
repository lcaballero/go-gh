package cli

import (
	"github.com/lcaballero/go-gh/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tag_Cli(t *testing.T) {
	app := ToApp(conf.Config{})
	assert.NotNil(t, app)
}
