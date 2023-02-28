package main

import (
	"context"
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
	for _, script := range s.Scripts {
		if err = s.Command(exe).Args(run, script).Dir(s.Source).Exec(); nil != err {
			return
		}
	}

	return
}
