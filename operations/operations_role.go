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

import "fmt"

// AddTimeAlive increases the total time this class has spent alive by timeAlive
// The intent is to track how long this class survives on average
func (c Controller) AddTimeAlive(class int, time int, currentMap string) {

	// Skip infc_hardcorepit because one round is 1:00 instead of 5:00 and skews the average
	if currentMap == "infc_hardcorepit" {
		return
	}

	role, err := c.RoleRepository.FindRoleByIndex(class)
	if err != nil {
		fmt.Printf("could not add time to %d\n", class)
	}
	role.AddTimeAlive(time)
	c.RoleRepository.StoreRole(role)
}
