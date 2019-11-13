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

// Class contains stats associated with a player that has chosen
// this class is in the past rounds
type Class struct {
	ID       uint
	PlayerID uint
	Name     string
	Kills    int
	Picks    int
	Record   int
	Score    int
}

// ClassRepository defines the CRUD operations required to be implemented
// to persist a Class
type ClassRepository interface {

	// StoreClasses persists all classes
	StoreClasses(classes []Class)

	// ResetClass clears all data
	ResetClass()
}

// Classes contains all current human classes
//
// Each class is attributed a number. Keep in mind that this is not 1:1 with classes.h and instead
// all classes were shifted by -2 for easier array indexing. So mercenary is class 0 and looper class 9
// "Unknown" counts as zombie.
var Classes = map[int]string{
	0:  "Mercenary",
	1:  "Medic",
	2:  "Hero",
	3:  "Engineer",
	4:  "Soldier",
	5:  "Ninja",
	6:  "Sniper",
	7:  "Scientist",
	8:  "Biologist",
	9:  "Looper",
	10: "Unknown",
	12: "Smoker",
	13: "Boomer",
	14: "Hunter",
	15: "Bat",
	16: "Ghost",
	17: "Spider",
	18: "Ghoul",
	19: "Slug",
	20: "Voodoo",
	21: "Witch",
	22: "Undead"}

// Unknown is used when the class can not yet be determined.
// This occures e.g. when an infclass-stats instance is started in the middle of a round
// Classes are determined first by a player choosing a class
var Unknown = 10

// HumanLimit means everything class x < HumanLimit is a human class
var HumanLimit = 10
