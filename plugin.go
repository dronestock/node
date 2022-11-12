package main

import (
	"reflect"

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

func (p *plugin) Steps() (steps drone.Steps) {
	steps = drone.Steps{
		drone.NewStep(p.install, drone.Name(`依赖`)),
	}

	options := drone.NewStepOptions(drone.Name(`脚本`))
	if reflect.DeepEqual(p.Scripts, []string{"build"}) {
		options = append(options, drone.Interrupt())
	}
	steps.Add(drone.NewStep(p.scripts, options...))

	return
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Source),
		field.Strings(`scripts`, p.Scripts...),
	}
}
