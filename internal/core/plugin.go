package core

import (
	"reflect"

	"github.com/dronestock/drone"
	"github.com/dronestock/node/internal/config"
	"github.com/dronestock/node/internal/internal/constant"
	"github.com/dronestock/node/internal/internal/core"
	"github.com/dronestock/node/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Plugin struct {
	drone.Base

	// 执行程序
	Binary config.Binary `default:"${BINARY}"`
	// 代码
	Source string `default:"${SOURCE=.}" validate:"required"`
	// 脚本列表
	Scripts []string `default:"${SCRIPTS=['build']}" validate:"required,dive"`
	// 类型
	Type core.Type `default:"${TYPE=pnpm}" validate:"oneof=npm yarn pnpm"`
}

func New() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Steps() (steps drone.Steps) {
	steps = drone.Steps{
		drone.NewStep(step.NewInstall(&p.Base, p.Source, p.Type, &p.Binary)).Name("依赖").Build(),
	}

	scripts := drone.NewStep(step.NewScripts(&p.Base, p.Source, p.Type, &p.Binary, p.Scripts...)).Name("脚本")
	if reflect.DeepEqual(p.Scripts, []string{"build"}) {
		scripts.Interrupt()
	}
	steps.Add(scripts.Build())

	return
}

func (p *Plugin) Setup() (err error) {
	p.Cleanup().Name("清理模块").File(constant.NodeModules).Build()

	return
}

func (p *Plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("scripts", p.Scripts),
	}
}
