package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) npmScript(script string) error {
	return p.Exec(npmExe, drone.Args(`run`, script))
}

func (p *plugin) npmInstall() error {
	return p.Exec(npmExe, drone.Args(`install`))
}
