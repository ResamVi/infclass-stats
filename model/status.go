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

// Status holds current state of the last 5 rounds (called a single map)
//
// State contains information that are not shared between multiple instances of infclass-stats
// because clients with the same name can be on multiple servers
type Status struct {
	// Name of player
	Name string

	// Used to whisper to him
	ClientID int64

	// Total kills on this map
	Kills int

	// Total score on this map
	Score int

	// Player has played this map, but may not be connected currently
	Connected bool

	// Score gained while being human on this single round
	HumanScore int

	// Which human class was initially chosen
	ClassPicked int

	// Either ClassPicked or a zombie class if player is infected.
	// You can deduce whether player is infected or not via this attribute
	CurrentClass int
}

// StatusRepository defines the CRUD operations required to be implemented
// to persist a Status
type StatusRepository interface {

	// StoreStatus persists a status.
	StoreStatus(status *Status)

	// FindStatusByName returns a status of the associated name
	FindStatusByName(name string) *Status

	FindStatus() []*Status

	// Sort and return all current players by score
	Sort() []*Status

	// ResetStatus clears the table of all rows
	ResetStatus()

	SetMap(newMap string)

	GetMap() string
}

// AddKill adds one kill to the total that is tracked for 5 rounds or one map
func (s *Status) AddKill() {
	s.Kills++
}

// SetClass changes the class of this player. If the player has picked a class
// and turns zombie the class is adjusted but his past choice is saved as well.
func (s *Status) SetClass(class int) {
	if IsInfected(class) {
		s.CurrentClass = class
	} else {
		s.ClassPicked = class
		s.CurrentClass = class
	}
}

// Disconnect tags the player as not connected, he is then finally removed when the map changes after the fifth round
func (s *Status) Disconnect() {
	s.Connected = false
}

// AddScore increases the score. The score is reset after the map changes though.
func (s *Status) AddScore(points int) {
	if s.Connected {
		s.Score += points

		if !s.IsInfected() {
			s.HumanScore += points
		}
	}
}

// Reset sets the score counter and class back to 0 for a new round that has started
func (s *Status) Reset() {
	s.CurrentClass = Unknown
	s.ClassPicked = Unknown
	s.HumanScore = 0
}

// IsInfected checks if the class is in the range of all human classes.
//
// Each class is attributed a number. Keep in mind that this is not 1:1 with classes.h and instead
// all classes were shifted by -2 for easier array indexing. So mercenary is class 0 and looper class 9 (see models/Player.go)
// "Unknown" counts as zombie.
func (s *Status) IsInfected() bool {
	return s.CurrentClass >= HumanLimit
}

// IsInfected is a helper class same as Status.IsInfected
func IsInfected(class int) bool {
	return class >= HumanLimit
}
