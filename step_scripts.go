package main

import (
	"context"
	"strings"

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
	sa := args.New().Build().Add(strings.Join(s.Scripts, space)).Build()
	_, err = s.Command(exe).Args(sa).Dir(s.Source).Build().Exec()

	return
}
