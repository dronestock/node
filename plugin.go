package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 代码
	Source string `default:"${SOURCE=.}" validate:"required"`
	// 脚本列表
	Scripts []string `default:"${SCRIPTS=['build']}" validate:"required,dive"`

	card *card
}

func newPlugin() drone.Plugin {
	return &plugin{
		card: new(card),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.install, drone.Name(`依赖`)),
		drone.NewStep(p.scripts, drone.Name(`脚本`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Source),
		field.Strings(`scripts`, p.Scripts...),
	}
}
