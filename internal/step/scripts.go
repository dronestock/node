package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/node/internal/config"
	"github.com/dronestock/node/internal/internal/constant"
	"github.com/dronestock/node/internal/internal/core"
	"github.com/goexl/gox/args"
)

type Scripts struct {
	base    *drone.Base
	source  string
	typ     core.Type
	binary  *config.Binary
	scripts []string
}

func NewScripts(base *drone.Base, source string, typ core.Type, binary *config.Binary, scripts ...string) *Scripts {
	return &Scripts{
		base:    base,
		source:  source,
		typ:     typ,
		binary:  binary,
		scripts: scripts,
	}
}

func (s *Scripts) Runnable() bool {
	return true
}

func (s *Scripts) Run(ctx context.Context) (err error) {
	switch s.typ {
	case core.TypeNpm:
		err = s.npm(ctx)
	case core.TypeYarn:
		err = s.yarn(ctx)
	case core.TypePnpm:
		err = s.pnpm(ctx)
	}

	return
}

func (s *Scripts) pnpm(_ context.Context) (err error) {
	sa := args.New().Build()
	if 1 == len(s.scripts) {
		sa.Subcommand(s.scripts[0])
	} else {
		sa.Subcommand(s.scripts[0], s.scripts[1:]...)
	}
	_, err = s.base.Command(s.binary.Pnpm).Args(sa.Build()).Dir(s.source).Build().Exec()

	return
}

func (s *Scripts) yarn(_ context.Context) (err error) {
	for _, script := range s.scripts {
		sa := args.New().Build()
		sa.Subcommand(script)
		if _, err = s.base.Command(s.binary.Yarn).Args(sa.Build()).Dir(s.source).Build().Exec(); nil != err {
			break
		}
	}

	return
}

func (s *Scripts) npm(_ context.Context) (err error) {
	for _, script := range s.scripts {
		sa := args.New().Build().Subcommand(constant.Run)
		sa.Subcommand(script)
		if _, err = s.base.Command(s.binary.Npm).Args(sa.Build()).Dir(s.source).Build().Exec(); nil != err {
			break
		}
	}

	return
}
