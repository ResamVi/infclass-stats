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

// Player contains overall performance statistics of a player
// To clearly identify the player, a Name has to be set.
type Player struct {
	ID             uint
	Name           string `gorm:"type:varchar(100);unique"`
	Time           int
	Score          int
	KillCount      int
	RoundsSurvived int

	// These variables will reset daily as opposed to weekly
	Time24           int
	Score24          int
	KillCount24      int
	RoundsSurvived24 int

	// Class specifc
	Classes []Class
}

// PlayerRepository defines the CRUD operations required to be implemented
// to persist a Player
type PlayerRepository interface {

	// StorePlayer persists a player. Can be used to update attributes of an already persisted player.
	StorePlayer(player *Player)

	// FindPlayerByName returns a player by his name
	FindPlayerByName(name string) (*Player, error)

	// FindPlayer prepares a list of all players
	FindPlayer() []Player

	// ResetPlayer clears the table of all rows
	ResetPlayer()

	// Reset24 clears all tables of daily data (suffix '24'). Done each day on midnight
	ResetPlayer24()
}

// AddTime increases the player's time-spent-in-server counter.
func (p *Player) AddTime() {
	p.Time++
	p.Time24++
}

// AddScore increases the player's total score
// Also increases the class score the player has chosen when he was not infected while scoring.
func (p *Player) AddScore(points int, class int, infected bool) {
	p.Score += points
	p.Score24 += points
	if !infected {
		p.Classes[class].Score += points
	}
}

// AddKill increases the player's kill count.
// Also increments the class-specifc 'kills' columns in table classes associated with the player when he was not infected while killing.
// Does nothing when player with name was not found
func (p *Player) AddKill(class int, infected bool) {
	p.KillCount++
	p.KillCount24++
	if !infected {
		p.Classes[class].Kills++
	}
}

// AddPick increases the player's pick count for a specific class
// Warning: this function will fail if class is a zombie class
func (p *Player) AddPick(class int) {
	if !IsInfected(class) {
		p.Classes[class].Picks++
	}
}

// Survive is called when a player survives after the time limit is up. Increments the counter.
func (p *Player) Survive() {
	p.RoundsSurvived++
	p.RoundsSurvived24++
}

// CheckRecord will compare the score gained on this round and whether they have been the highest
// of any rounds played on this class. If the highest score has been broken update the record
func (p *Player) CheckRecord(score int, class int) {
	if !IsInfected(class) && score > p.Classes[class].Record {
		p.Classes[class].Record = score
	}
}

// Summary returns all the general numbers of this player: killcount, daily killcount, rounds survived, daily rounds survived, score, daily score.
func (p *Player) Summary() (int, int, int, int, int, int) {
	return p.KillCount, p.KillCount24, p.RoundsSurvived, p.RoundsSurvived24, p.Score, p.Score24
}
