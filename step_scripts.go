package main

import (
	"context"

	"github.com/goexl/gox/args"
)

type stepScripts struct {
	*plugin
}

func newScriptsStep(plugin *plugin) *stepScripts {
	return &stepScripts{
		plugin: plugin,
	}
}

func (s *stepScripts) Runnable() bool {
	return true
}

func (s *stepScripts) Run(_ context.Context) (err error) {
	sa := args.New().Build()
	if 1 == len(s.Scripts) {
		sa.Subcommand(s.Scripts[0])
	} else {
		sa.Subcommand(s.Scripts[0], s.Scripts[1:]...)
	}
	_, err = s.Command(exe).Args(sa.Build()).Dir(s.Source).Build().Exec()

	return
}
