/*
	Copyright (C) 2018 Julien Midedji

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

// Package operations orchestrate the flow of data to and from the model
package operations

import "infclass-stats/model"
import "os"

// RconCommands are all commands in use by infclass-stats
type RconCommands interface {
	Say(message string)
	Whisper(clientID int64, message string)
	Status()
}

// Controller handles the data coming from the models to the database
// The CRUD operations are defined in the Repositories of package model and implemented in package infrastructure
// Implementations vary in SQL-dialect (MySQL, NoSQL, gorm, ...)
type Controller struct {
	ClassRepository    model.ClassRepository
	MapRepository      model.MapRepository
	ClientRepository   model.ClientRepository
	PlayerRepository   model.PlayerRepository
	ActivityRepository model.ActivityRepository
	RoleRepository     model.RoleRepository
	StatusRepository   model.StatusRepository
	Rcon               RconCommands
}

// Reset clears all tables of their data. Done each week on Thursday 23:59:58 for a fresh
// series of people to reach the top ranks
func (c Controller) Reset() {
	c.PlayerRepository.ResetPlayer()
	c.ClassRepository.ResetClass()
	c.MapRepository.ResetMap()
	c.ActivityRepository.ResetActivity()
	c.RoleRepository.ResetRole()
	c.ClientRepository.ResetClient()
	os.Exit(1)
}

// Reset24 clears all tables of daily data (suffix '24'). Done each day on midnight
func (c Controller) Reset24() {
	c.PlayerRepository.ResetPlayer24()
}
