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

// A Client is a connected player on the server
type Client struct {
	ID   uint
	Name string `gorm:"type:varchar(100);unique"`
}

// ClientRepository defines the CRUD operations required to be implemented
// to persist a Client
type ClientRepository interface {

	// StoreClient persists the activity
	StoreClient(client *Client)

	// FindClientByName returns a currently connected client with the given name
	FindClientByName(name string) (*Client, error)

	// FindClient prepares a list of players currently on server
	FindClient() []Client

	// DeleteClientByName removes the row with the name.
	// Does nothing if row does not exist
	DeleteClientByName(name string)

	// ResetClient clears the table of all rows
	ResetClient()

	// Get the amount of rows in the table of clients connected
	Count() int
}
