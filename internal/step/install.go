package step

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dronestock/drone"
	"github.com/dronestock/node/internal/config"
	"github.com/dronestock/node/internal/internal/constant"
	"github.com/dronestock/node/internal/internal/core"
	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
)

type Install struct {
	base    *drone.Base
	source  string
	typ     core.Type
	binary  *config.Binary
	scripts []string
}

func NewInstall(base *drone.Base, source string, typ core.Type, binary *config.Binary) *Install {
	return &Install{
		base:   base,
		source: source,
		typ:    typ,
		binary: binary,
	}
}

func (i *Install) Runnable() bool {
	return true
}

func (i *Install) Run(ctx context.Context) (err error) {
	if le := i.link(ctx); nil != le {
		err = le
	} else if ie := i.install(ctx); nil != ie {
		err = ie
	}

	return
}

func (i *Install) link(_ context.Context) (err error) {
	if core.TypePnpm == i.typ {
		return
	}

	name := cryptor.New(os.Getenv(constant.RepositoryLink)).Md5().Hex()
	link := filepath.Join(i.source, constant.NodeModules)
	modules := filepath.Join(os.Getenv(constant.ModulePath), name)
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

func (i *Install) install(ctx context.Context) (err error) {
	switch i.typ {
	case core.TypeNpm:
		err = i.npm(ctx)
	case core.TypeYarn:
		err = i.yarn(ctx)
	case core.TypePnpm:
		err = i.pnpm(ctx)
	}

	return
}

func (i *Install) npm(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(constant.Install)
	if i.base.Verbose {
		ia.Flag(constant.Verbose)
	}
	_, err = i.base.Command(i.binary.Npm).Args(ia.Build()).Dir(i.source).Build().Exec()

	return
}

func (i *Install) pnpm(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(constant.Install)
	if i.base.Verbose {
		ia.Flag(constant.Verbose)
	}
	_, err = i.base.Command(i.binary.Pnpm).Args(ia.Build()).Dir(i.source).Build().Exec()

	return
}

func (i *Install) yarn(_ context.Context) (err error) {
	ia := args.New().Build().Subcommand(constant.Install).Flag("prefer-offline", "parallel", "frozen-lockfile")
	if i.base.Verbose {
		ia.Flag(constant.Verbose)
	}
	_, err = i.base.Command(i.binary.Yarn).Args(ia.Build()).Dir(i.source).Build().Exec()

	return
}

func (i *Install) setenv(modules string) (err error) {
	path := os.Getenv(constant.Path)
	newPath := gox.StringBuilder(path, filepath.ListSeparator, filepath.Join(modules, constant.Bin)).String()
	if sne := os.Setenv(constant.NodePath, modules); nil != sne {
		err = sne
	} else if spe := os.Setenv(constant.Path, newPath); nil != spe {
		err = spe
	}

	return
}
