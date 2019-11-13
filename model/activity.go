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

// An Activity is a snapshot of the current state of activity on the server.
// Used to keep record on how many players have played InfClass at specific
// timestamps.
type Activity struct {
	ID     uint
	Timestamp   int64 // Unix time (easier for apex)
	Amount int
}

// ActivityRepository defines the CRUD operations required to be implemented
// to persist an Activity
type ActivityRepository interface {

	// Store persists the activity
	StoreActivity(activity *Activity)

	// Find prepares a list of all activities ordered by time
	FindActivity() []Activity

	// ResetActivity clears all data
	ResetActivity()
}
