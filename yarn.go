package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) yarnScript(script string) error {
	return p.Exec(yarnExe, drone.Args(script))
}

func (p *plugin) yarnInstall() error {
	return p.Exec(yarnExe, drone.Args(`install`))
}
