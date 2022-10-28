package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) scripts() (undo bool, err error) {
	for _, script := range p.Scripts {
		if err = p.Exec(exe, drone.Args(`run`, script), drone.Dir(p.Source)); nil != err {
			return
		}
	}

	return
}
