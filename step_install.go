package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
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
	if le := i.link(ctx); nil != le {
		err = le
	} else if ie := i.install(ctx); nil != ie {
		err = ie
	}

	return
}

func (i *stepInstall) link(_ context.Context) (err error) {
	name := cryptor.New(os.Getenv(repositoryLink)).Md5().Hex()
	link := filepath.Join(i.Source, nodeModules)
	modules := filepath.Join(os.Getenv(modulePath), name)
	if _, se := os.Stat(modules); nil != se && os.IsNotExist(se) {
		err = os.MkdirAll(modules, os.ModePerm)
	}
	if _, se := os.Lstat(link); nil != se && os.IsNotExist(se) && nil == err {
		err = os.Symlink(modules, link)
	}
	if nil == err {
		err = i.setenv(modules)
	}

	return
}

func (i *stepInstall) install(ctx context.Context) (err error) {
	switch i.Type {
	case typeYarn:
		err = i.yarn(ctx)
	case typePnpm:
		err = i.pnpm(ctx)
	}

	return
}

func (i *stepInstall) pnpm(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(install).Flag("no-frozen-lockfile")
	if i.Verbose {
		ia.Flag(verbose)
	}
	_, err = i.Command(i.Binary.Pnpm).Args(ia.Build()).Dir(i.Source).Build().Exec()

	return
}

func (i *stepInstall) yarn(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(install).Flag("prefer-offline", "parallel", "frozen-lockfile")
	if i.Verbose {
		ia.Flag(verbose)
	}
	_, err = i.Command(i.Binary.Yarn).Args(ia.Build()).Dir(i.Source).Build().Exec()

	return
}

func (i *stepInstall) setenv(modules string) (err error) {
	newPath := gox.StringBuilder(os.Getenv(path), filepath.ListSeparator, filepath.Join(modules, bin)).String()
	if sne := os.Setenv(nodePath, modules); nil != sne {
		err = sne
	} else if spe := os.Setenv(path, newPath); nil != spe {
		err = spe
	}

	return
}
