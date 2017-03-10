package resources

import (
	"github.com/lcaballero/go-gh/conf"
	"github.com/lcaballero/go-gh/cli"
)

type Loader interface {
	LoadToken() (string, error)
	LoadLocals() (conf.Locals, error)
}

type Resources struct {
	loader cli.Loader
}

func New() Resources {
	return Resources {
		loader: cli.NewLoader(),
	}
}

func (r Resources) LoadToken() (string, error) {
	file := r.loader.FileProvider(cli.GO_GH_TOKEN)
	return r.loader.LoadToken(file)
}

func (r Resources) LoadLocals() (conf.Locals, error) {
	file := r.loader.FileProvider(cli.GO_GH)
	return r.loader.LoadLocals(file)
}
