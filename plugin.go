package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 类型
	Type string `default:"${PLUGIN_TYPE=${TYPE=npm}}" validate:"required,oneof=npm yarn"`
	// 脚本列表
	Scripts []string `default:"${PLUGIN_SCRIPTS=${SCRIPTS=['build']}}" validate:"required,dive"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.install, drone.Name(`安装依赖`)),
		drone.NewStep(p.scripts, drone.Name(`执行脚本`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Folder),
		field.String(`type`, p.Type),
		field.Strings(`scripts`, p.Scripts...),
	}
}
