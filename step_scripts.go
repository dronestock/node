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

func (s *stepScripts) Run(ctx context.Context) (err error) {
	switch s.Type {
	case typeYarn:
		err = s.yarn(ctx)
	case typePnpm:
		err = s.pnpm(ctx)
	}

	return
}

func (s *stepScripts) pnpm(_ context.Context) (err error) {
	sa := args.New().Build()
	if 1 == len(s.Scripts) {
		sa.Subcommand(s.Scripts[0])
	} else {
		sa.Subcommand(s.Scripts[0], s.Scripts[1:]...)
	}
	_, err = s.Command(s.Binary.Pnpm).Args(sa.Build()).Dir(s.Source).Build().Exec()

	return
}

func (s *stepScripts) yarn(_ context.Context) (err error) {
	for _, script := range s.Scripts {
		sa := args.New().Build()
		sa.Subcommand(script)
		if _, err = s.Command(s.Binary.Yarn).Args(sa.Build()).Dir(s.Source).Build().Exec(); nil != err {
			break
		}
	}

	return
}
