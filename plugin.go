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
		drone.NewStep(newInstallStep(p)).Name("依赖").Build(),
	}

	scripts := drone.NewStep(newScriptsStep(p)).Name("脚本")
	if reflect.DeepEqual(p.Scripts, []string{"build"}) {
		scripts.Interrupt()
	}
	steps.Add(scripts.Build())

	return
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("scripts", p.Scripts),
	}
}
