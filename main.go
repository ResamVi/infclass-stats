/*
	Copyright (C) 2018  Julien Midedji

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Package infclass-stats initializes a server that provides a websocket endpoint to
// gain access to a selected amount of statistics collected via an econ connection of a teeworlds server.
package main

import (
	"github.com/op/go-logging"

	"infclass-stats/config"
	"infclass-stats/infrastructure"
	"infclass-stats/operations"
)

var log = logging.MustGetLogger("main")
var format = logging.MustStringFormatter(`%{color}%{time:15:04:05} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`)

func main() {

	if config.DEBUG_MODE {
		logging.SetLevel(logging.DEBUG, "main")
	} else {
		logging.SetLevel(logging.INFO, "main")
	}

	log.Info("Connecting to econ at " + config.SERVER_IP + ":" + config.ECON_PORT)
	file, rcon := infrastructure.Authenticate()

	log.Info("Initialize database")
	gormDB := infrastructure.NewGormDB()
	serverState := infrastructure.NewServerState()

	controller := operations.Controller{}
	controller.ClientRepository = gormDB
	controller.RoleRepository = gormDB
	controller.MapRepository = gormDB
	controller.ClassRepository = gormDB
	controller.PlayerRepository = gormDB
	controller.ActivityRepository = gormDB
	controller.StatusRepository = serverState
	controller.Rcon = rcon

	infrastructure.Listen(file, controller)

	if config.SERVE_API {
		log.Info("Starting HTTP server on localhost:8002")
		infrastructure.Serve(controller)
	}

	log.Info("Start keeping score")
	controller.Track()
}
