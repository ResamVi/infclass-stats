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

import "infclass-stats/model"

// CreateMap creates a new map if has not been seen play and was not yet
// found in the database
func (c Controller) CreateMap(name string) *model.Map {
	mapp := model.Map{Name: name}
	c.MapRepository.StoreMap(&mapp)
	return &mapp
}

// Winner picks the winner of this round on this map and how long it took to complete
// depending on who won (string "humans" or zombies when anything else) increments the human/zombie column.
// If an entry with the given name of the map does not exist in the table it will be created
func (c Controller) Winner(race string, duration int) {
	mapp, err := c.MapRepository.FindMapByName(c.StatusRepository.GetMap())

	if err != nil {
		mapp = c.CreateMap(c.StatusRepository.GetMap())
	}

	mapp.Win(race, duration)
	c.MapRepository.StoreMap(mapp)
}
