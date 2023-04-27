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

func (i *stepInstall) Run(_ context.Context) (err error) {
	switch i.Type {
	case typeYarn:
		err = i.yarn()
	case typePnpm:
		err = i.pnpm()
	}

	return
}

func (i *stepInstall) pnpm() (err error) {
	ia := args.New().Build().Subcommand(install)
	ia.Flag("no-frozen-lockfile")
	if i.Verbose {
		ia.Flag(verbose)
	}
	_, err = i.Command(i.Binary.Pnpm).Args(ia.Build()).Dir(i.Source).Build().Exec()

	return
}

func (i *stepInstall) yarn() (err error) {
	ia := args.New().Build().Subcommand(install).Flag("prefer-offline", "parallel")
	if i.Verbose {
		ia.Flag(verbose)
	}
	_, err = i.Command(i.Binary.Yarn).Args(ia.Build()).Dir(i.Source).Build().Exec()

	return
}
