package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) install() (undo bool, err error) {
	err = p.Exec(exe, drone.Dir(p.Folder))

	return
}
