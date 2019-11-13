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
	"time"
)

// Snapshot creates a new Activity entry.
// Counts the players currently on the server and makes an entry with a timestamp in the database
func (c Controller) Snapshot() {
	activity := model.Activity{Timestamp: time.Now().Unix(), Amount: c.ClientRepository.Count()}
	c.ActivityRepository.StoreActivity(&activity)
}
