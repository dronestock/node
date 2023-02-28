package main

import (
	"context"
	"strings"
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

func (s *stepScripts) Run(_ context.Context) error {
	return s.Command(exe).Args(strings.Join(s.Scripts, space)).Dir(s.Source).Exec()
}
