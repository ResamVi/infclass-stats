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

package operations

import (
	"fmt"
	"infclass-stats/model"
)

// CommandHelp prints all available commands
func (c Controller) CommandHelp(clientID int64) {
	c.Rcon.Whisper(clientID, fmt.Sprintf("Available Commands:"))
	c.Rcon.Whisper(clientID, fmt.Sprintf("!me, !kills, !active, !survivors, !best, !wins, !player <name>"))
	c.Rcon.Whisper(clientID, fmt.Sprintf("!<class>, !picks <class>, !record <class>, !kills <class>"))
	c.Rcon.Whisper(clientID, fmt.Sprintf("!today <kills/active/survivors>"))
}

// CommandMe prints your own stat summary
func (c Controller) CommandMe(clientID int64, playerName string) {
	p, err := c.PlayerRepository.FindPlayerByName(playerName)
	players := c.PlayerRepository.FindPlayer()

	if err != nil {
		c.Rcon.Whisper(clientID, fmt.Sprintf("Player not in database"))
		return
	}

	kills, kills24, survived, survived24, score, score24 := p.Summary()
	rank, rating := GetRank(playerName, players)

	classes := GetClassRank(playerName, players)
	first, second, third := classes[0], classes[1], classes[2]

	c.Rcon.Whisper(clientID, fmt.Sprintf("-- %s --", playerName))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Rank: %d, Rating: %.f", rank, rating))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Killcount today: %d, Killcount this week: %d", kills24, kills))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Survived today: %d, Survived this week: %d", survived24, survived))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Total scored today: %d, Scored this week: %d", score24, score))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Total scored today: %d, Scored this week: %d", score24, score))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Best classes: %s (%.f.) %s (%.f.) %s (%.f.)", first.Key, first.Value, second.Key, second.Value, third.Key, third.Value))
}

// CommandWins prints how many human wins vs zombie wins
func (c Controller) CommandWins(clientID int64) {
	maps := c.MapRepository.FindMap()
	c.Rcon.Whisper(clientID, fmt.Sprintf("Human Wins: %d", HumanWins(maps)))
	c.Rcon.Whisper(clientID, fmt.Sprintf("Zombie Wins: %d", ZombieWins(maps)))
}

func (c Controller) PrintRanking(clientID int64, list []Entry) {
	for i, entry := range list {
		if i < len(list) && i < 10 {
			log.Debugf("%d | %s -- %.0f", i+1, entry.Key, entry.Value)
			c.Rcon.Whisper(clientID, fmt.Sprintf("%d | %s -- %.0f", i+1, entry.Key, entry.Value))
		}
	}
}
func calculateAverage(list []Entry) float32 {
	result := float32(0)
	count := 0
	for _, entry := range list {
		if entry.Value > 0 {
			result += entry.Value
			count++
		}
	}

	if count == 0 {
		return float32(0)
	}

	return result / float32(count)

}

func getRatio(player model.Player, ratios []Entry) float32 {
	for _, ratio := range ratios {
		if ratio.Key == player.Name {
			return ratio.Value
		}
	}
	return float32(0)
}
