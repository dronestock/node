package main

import (
	"context"

	"github.com/goexl/gox/args"
)

type stepInstall struct {
	*plugin
}

func newInstallStep(plugin *plugin) *stepInstall {
	return &stepInstall{
		plugin: plugin,
	}
}

func (i *stepInstall) Runnable() bool {
	return true
}

func (i *stepInstall) Run(ctx context.Context) (err error) {
	switch i.Type {
	case typeYarn:
		err = i.yarn(ctx)
	case typePnpm:
		err = i.pnpm(ctx)
	}

	return
}

func (i *stepInstall) pnpm(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(install).Flag("no-frozen-lockfile").Build()
	_, err = i.Command(exe).Args(ia).Dir(i.Source).Build().Exec()

	return
}

func (i *stepInstall) yarn(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(install).Build()
	_, err = i.Command(exe).Args(ia).Dir(i.Source).Build().Exec()

	return
}
