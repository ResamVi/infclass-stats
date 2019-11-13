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

package model

// Map records human/zombie wins on each map and survival duration
// A map is identified by their name
type Map struct {
	ID       uint
	Name     string
	Human    int
	Zombie   int
	Eligible int
	Duration int
}

// MapRepository defines the CRUD operations required to be implemented
// to persist map statistics
type MapRepository interface {

	// StoreMap persists a map. Can be used to update attributes of an already persisted map.
	StoreMap(mapp *Map)

	// FindMapByName returns a map with the given name (e.g. "infc_hardcorepit")
	FindMapByName(name string) (*Map, error)

	// FindMap prepares a list of all maps played
	FindMap() []Map

	// ResetMap clears the table of all rows
	ResetMap()
}

// Win picks the winner of this round on this map and how long it took to complete
// depending on who won (string "humans" or zombies when anything else) increments the human/zombie column.
// If an entry with the given name of the map does not exist in the table it will be created
func (m *Map) Win(race string, duration int) {
	if race == "humans" {
		m.Human++
	} else {
		m.Zombie++
	}
	m.Duration += duration
	m.Eligible++
}
