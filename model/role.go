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

// Role is a container that combines all player performances on a specific class to gain
// an overview how a class overall is performing.
type Role struct {
	ID        uint
	Name      string `gorm:"type:varchar(100);unique"`
	Survived  int
	AliveTime int
	PlayCount int
}

// RoleRepository defines the CRUD operations required to be implemented
// to persist a Role
type RoleRepository interface {

	// StoreRole persists a role. Can be used to update attributes of an already persisted player.
	StoreRole(role *Role)

	// FindRoleByIndex returns a role by their index (see: model/Class.go)
	FindRoleByIndex(index int) (*Role, error)

	// FindRole prepares a list of all roles
	FindRole() []Role

	// ResetRole clears the table of all rows
	ResetRole()
}

// AddTimeAlive increases the total time this class has spent alive by timeAlive (if its not a zombie)
// The intent is to track how long this class survives on average
func (r *Role) AddTimeAlive(timeAlive int) {
	if !IsInfected((int)(r.ID - 1)) {
		r.AliveTime += timeAlive
	}
	r.AddPick()
}

// AddPick increases the play_count
func (r *Role) AddPick() {
	r.PlayCount++
}

// Survive is called when a player with that choosen role survives after the time limit is up.
// Increments the counter.
func (r *Role) Survive() {
	r.Survived++
}
