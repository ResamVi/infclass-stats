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

// CreatePlayer creates a new Player entry in the database.
// Does nothing when entry with name already exists.
func (c Controller) CreatePlayer(name string) {
	_, err := c.PlayerRepository.FindPlayerByName(name)

	if err != nil {

		// The name will be stored as a chain of escaped unicode characters.
		// TODO: change StorePlayer so no converting is done in usecases
		newPlayer := model.Player{Name: fmt.Sprintf("%+q", name)}
		c.PlayerRepository.StorePlayer(&newPlayer)

		// Along with a player the set of class objects associated with the player's performance will be created.
		newClasses := []model.Class{}
		for i := 0; i < model.HumanLimit+1; i++ { // All human classes + unknown are created
			newClasses = append(newClasses, model.Class{
				ID:       (newPlayer.ID-1)*11 + uint(i+1), // if playerID is e.g. 2 then classId is 22-33
				PlayerID: newPlayer.ID,
				Name:     model.Classes[i]})
		}

		c.ClassRepository.StoreClasses(newClasses)
	}
}

// AddTime increases the player's time-spent-in-server counter.
// Does nothing when player with name was not found
func (c Controller) AddTime(names []string) {
	for _, name := range names {
		player, err := c.PlayerRepository.FindPlayerByName(name)
		if err != nil {
			fmt.Println(name + " missing in database")
			continue
		}
		player.AddTime()
		c.PlayerRepository.StorePlayer(player)
	}
}

// CheckRecord will compare the score gained on this round and whether they have been the highest
// of any rounds played on this class. If the highest score has been broken update the record
func (c Controller) CheckRecord(name string, score int, class int) {
	player, err := c.PlayerRepository.FindPlayerByName(name)
	if err != nil {
		fmt.Println(name + "cannot survive")
	}
	player.CheckRecord(score, class)
	c.PlayerRepository.StorePlayer(player)
}
