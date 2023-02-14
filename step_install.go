package main

import (
	"context"
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

func (i *stepInstall) Run(_ context.Context) error {
	return i.Command(exe).Args("--prefer-offline").Dir(i.Source).Exec()
}
