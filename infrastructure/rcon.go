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

package infrastructure

import (
	"fmt"
	"net"
)

// Rcon is used to execute rcon commands
type Rcon struct {
	conn net.Conn
}

// Say uses remote console command 'say' to broadcast to all players
func (r Rcon) Say(message string) {
	fmt.Fprintln(r.conn, "say "+message)
}

// Whisper uses remote console command 'whisper' to broadcast to a single player
func (r Rcon) Whisper(clientID int64, message string) {
	fmt.Fprintf(r.conn, "whisper %d %s\n", clientID, message)
}

// Status uses remote console command 'status' to display all players currently on the server
func (r Rcon) Status() {
	fmt.Fprintf(r.conn, "status\n")
}
