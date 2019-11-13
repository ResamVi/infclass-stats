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
	"fmt"
	"infclass-stats/model"
)

// ConnectClient records the player that has joined the server into the active player list
// Does nothing if entry with name already exists.
func (c Controller) ConnectClient(name string) {
	_, err := c.ClientRepository.FindClientByName(name)

	if err != nil {
		// TODO: change StorePlayer so no converting is done in usecases
		newClient := model.Client{Name: fmt.Sprintf("%+q", name)}
		c.ClientRepository.StoreClient(&newClient)
	}
}

// DisconnectClient removes the player from the active player list.
// The player's time-spent-active counter will stop increasing.
func (c Controller) DisconnectClient(name string) {
	c.ClientRepository.DeleteClientByName(name)
}

// LeaveAll clears the table containing all players currently online
func (c Controller) LeaveAll() {
	c.ClientRepository.ResetClient()
}
