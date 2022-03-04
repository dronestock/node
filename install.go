package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) install() (undo bool, err error) {
	err = p.Exec(pnpmExe, drone.Args(`install`), drone.Dir(p.Folder))

	return
}
