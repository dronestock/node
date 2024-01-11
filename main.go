package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/node/internal/core"
)

func main() {
	drone.New(core.New).Boot()
}
