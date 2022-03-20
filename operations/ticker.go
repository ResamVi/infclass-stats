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

// Package gamestate provides real-time analysis of the current game state on the server. This entails
// metrics like player count, current score, player is zombie/human etc.
//
// Unlike the data saved in the database everything gamestate tracks is only applicable to the single server
// that it tracks. Meaning for example the list of connected players (of table connected) may contain players of multiple servers and may differ
// to the array of playersOnline inside here.
//
// Some functions may have the same name as functions residing in database.go but with less arguments. This means that for that specific database request to
// be fulfilled some knowledge of the current gamestate has to be injected first which is done so in the function of the same name.
package operations

import (
	"fmt"
	"time"
)

// Track repeatedly scans the server of how many players are online
// and keeps track of the current date to coordinate the weekly/daily reset
func (c Controller) Track() {
	c.Rcon.Status()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		now := time.Now()

		/*fmt.Printf("%16s %3s \t %5s \t %7s \t %10s \t %10s \t %8s \t %8s\n", "", "ID", "Score", "HumanSc", "Pick", "Current", "Infected", "Connected")
		for _, status := range c.StatusRepository.FindStatus() {
			fmt.Printf("%16s %3d \t %5d \t %7d \t %10s \t %10s \t %8t \t %8t\n",
				status.Name,
				status.ClientID,
				status.Score,
				status.HumanScore,
				model.Classes[status.ClassPicked],
				model.Classes[status.CurrentClass],
				status.IsInfected(),
				status.Connected)
		}
		fmt.Println()*/

		// Take a Snapshot every 20 minute interval
		if now.Second() == 0 && (now.Minute() == 0 || now.Minute() == 20 || now.Minute() == 40) {
			c.Snapshot()
		}

		// Rescan all current connected players
		if now.Second() == 0 {
			c.LeaveAll()
			c.Rcon.Status()
		}

		// Reset on Thursday, 23:59:57 (give it a buffer of 2s)
		if now.Weekday() == 4 && now.Hour() == 23 && now.Minute() == 59 && now.Second() >= 57 {
			c.Reset()
		}

		// Reset of __24 stats at Midnight
		if now.Hour() == 23 && now.Minute() == 59 && now.Second() >= 57 {
			c.Reset24()
		}

		// Count up for every player connected (but not spectating)
		c.AddTimeToAll()
	}
}

// Enter occurs when a player with enters the game and sets up all relevant structures for monitoring.
// Set ClientID to -1 when it cannot be inferred from the log statement
func (c Controller) Enter(name string, clientID int64) {
	c.CreatePlayer(name)
	c.ConnectClient(name)
	c.CreateStatus(name, clientID)
}

// Leave will remove the player from the server list and thus gains no more active points.
func (c Controller) Leave(name string) {
	if name == "(connecting)" {
		return
	}

	c.DisconnectClient(name)
	status := c.StatusRepository.FindStatusByName(name)
	status.Disconnect()
}

// AddScore increases the player's total score and monitored score of the past five rounds on current map
// Does nothing when player with name was not found
func (c Controller) AddScore(name string, points int) {
	p, err := c.PlayerRepository.FindPlayerByName(name)
	status := c.StatusRepository.FindStatusByName(name)

	if err != nil {
		fmt.Println(name + " missing in database")
		return
	}
	p.AddScore(points, status.CurrentClass, status.IsInfected())
	c.PlayerRepository.StorePlayer(p)

	status.AddScore(points)
}

// AddKill increases the player's total kills, kills on class, and kills of the past five rounds on this map
func (c Controller) AddKill(name string) {
	p, err := c.PlayerRepository.FindPlayerByName(name)
	status := c.StatusRepository.FindStatusByName(name)

	if err != nil {
		fmt.Println(name + " missing in database")
	}
	p.AddKill(status.CurrentClass, status.IsInfected())
	c.PlayerRepository.StorePlayer(p)

	status.AddKill()
}

// SetClass set the current class played of the player with the given name and increments the counter how much this class has been picked
func (c Controller) SetClass(name string, class int) {

	if name == "(connecting)" {
		return
	}

	status := c.StatusRepository.FindStatusByName(name)
	status.SetClass(class)

	p, err := c.PlayerRepository.FindPlayerByName(name)
	if err != nil {
		fmt.Println("Could not set class")
	}
	p.AddPick(class)
	c.PlayerRepository.StorePlayer(p)

	role, err := c.RoleRepository.FindRoleByIndex(class)
	if err != nil {
		fmt.Printf("could not add play_count to %d\n", class)
	}
	role.AddPick()
	c.RoleRepository.StoreRole(role)

}

// Infect is called when a human turned a zombie or surviving the round to the end.
// Tracks how long the class which the player used survived.
func (c Controller) Infect(name string, aliveTime int) {

	// It may occur that (connecting) will be infected, i.e. a client that has just joined and is not yet fully connected
	if name == "(connecting)" {
		return
	}

	status := c.StatusRepository.FindStatusByName(name)
	c.AddTimeAlive(status.ClassPicked, aliveTime, c.StatusRepository.GetMap())
	//c.AddPick(status.CurrentClass)

}

// SurviveRound is called when a player survives when the time limit is up
// Increases his "amount survived" count and the "amount survived" of his chosen class
func (c Controller) SurviveRound(name string) {
	player, err := c.PlayerRepository.FindPlayerByName(name)
	status := c.StatusRepository.FindStatusByName(name)

	if err != nil {
		fmt.Println(name + "cannot survive")
	}
	player.Survive()
	c.PlayerRepository.StorePlayer(player)

	role, err := c.RoleRepository.FindRoleByIndex(status.ClassPicked)
	role.Survive()
	c.RoleRepository.StoreRole(role)
}

// CheckRecords will compare the score gained on this round and whether they have been the highest
// of any rounds played on this class. If the highest score has been broken update the record
func (c Controller) CheckRecords() {
	for _, status := range c.StatusRepository.FindStatus() {

		// Sadly the round ends before the +5 survival bonus is added, so we have to add it manually
		if !status.IsInfected() && status.Connected {
			status.HumanScore += 5
		}

		p, err := c.PlayerRepository.FindPlayerByName(status.Name)
		if err != nil {
			fmt.Println("Could not get player")
		}
		p.CheckRecord(status.HumanScore, status.ClassPicked)

		c.PlayerRepository.StorePlayer(p)
	}
}

// Summary broadcasts the 5 best players in score in the final round of the map
func (c Controller) Summary() {
	if c.ClientRepository.Count() < 3 {
		return
	}

	sorted := c.StatusRepository.Sort()

	c.Rcon.Say(fmt.Sprintf("--- HIGHEST SCORER ---\n"))
	for i, status := range sorted {
		if c.ClientRepository.Count() >= i && i < 6 {
			c.Rcon.Say(fmt.Sprintf("%d | %s: %d", i+1, status.Name, status.Score))
		}
	}
	c.Rcon.Say(fmt.Sprintf("Visit stats.resamvi.io for more stats"))

	time.AfterFunc(4*time.Second, c.StatusRepository.ResetStatus)
	time.AfterFunc(5*time.Second, c.Rcon.Status)
}

// Rank whispers to all players at what ranking they are
func (c Controller) Rank() {
	state := c.StatusRepository.FindStatus()
	players := c.PlayerRepository.FindPlayer()

	for _, player := range state {
		rank, rating := GetRank(player.Name, players)
		if player.ClientID != -1 {
			c.Rcon.Whisper(player.ClientID, fmt.Sprintf("-- STATS --"))
			c.Rcon.Whisper(player.ClientID, fmt.Sprintf("Your current rank: %d, Rating: %.f", rank, rating))
			c.Rcon.Whisper(player.ClientID, fmt.Sprintf("Command !help or http://stats.resamvi.io for more info"))
		}
	}
}
