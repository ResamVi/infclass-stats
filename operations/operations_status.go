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

package operations

import (
	"infclass-stats/model"
)

// CreateStatus creates a new status entry
// When player with name already exists, toggle the "Connected" to true
func (c Controller) CreateStatus(name string, clientID int64) {
	status := c.StatusRepository.FindStatusByName(name)

	if status == nil {
		newStatus := model.Status{Name: name, ClientID: clientID, Connected: true, ClassPicked: model.Unknown, CurrentClass: model.Unknown}
		c.StatusRepository.StoreStatus(&newStatus)
	} else {
		status.Connected = true
		status.ClientID = clientID
	}
}

// AddTimeToAll increments the active counter of all currently connected players
// If the player class is unknown we can reasonably assume he is spectating.
// This operations should be invoked every second.
func (c Controller) AddTimeToAll() {
	active := make([]string, 0)

	for _, status := range c.StatusRepository.FindStatus() {
		if status.Connected && status.ClassPicked != model.Unknown {
			active = append(active, status.Name)
		}
	}

	c.AddTime(active)
}

// NewRound is used when a new round is started. It is safer to explicity reset all scores and classes
// to match the current game state.
func (c Controller) NewRound() {
	for _, status := range c.StatusRepository.FindStatus() {
		status.Reset()
	}
}

// ChangeMap changes the map which currently is played
func (c Controller) ChangeMap(name string) {
	c.StatusRepository.SetMap(name)
}
