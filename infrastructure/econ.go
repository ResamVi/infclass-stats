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

// Package econ sets up an econ connection by authenticating and opening a new tcp socket.
// All incoming messages will be written to a file sent to the parser for evaluation
package infrastructure

import (
	"bufio"
	"bytes"
	"fmt"
	"infclass-stats/config"
	"infclass-stats/operations"
	"net"
	"os"
	"strings"
)

var reader *bufio.Reader

// Authenticate will initialize an econ connection and authorize with the parameters given
// in the config package. A goroutine will write the incoming messages into file and relay
// the message to the parser.
func Authenticate() (*bufio.Writer, Rcon) {
	conn, err := net.Dial("tcp", config.SERVER_IP+":"+config.ECON_PORT)
	if err != nil {
		log.Fatal("\nCheck if ec_bindaddr, ec_port and ec_password have been set in your autoexec.cfg and the server is running.\n" + err.Error())
	}

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error(err)
	}

	reader = bufio.NewReader(conn)
	writer := bufio.NewWriter(f)
	prompt, _ := reader.ReadString('\n')

	// Authenticate
	if strings.Contains(prompt, "Enter password") {
		fmt.Fprintln(conn, config.ECON_PASSWORD)
	} else {
		log.Fatal("No password was demanded - could not authenticate.")
	}

	response, _ := reader.ReadString('\n')
	if strings.Contains(response, "Authentication successful") {
		log.Info("Authentication succesful.")
	} else {
		log.Fatal("Authentication unsuccesful.")
	}

	return writer, Rcon{conn: conn}
}

func Listen(writer *bufio.Writer, c operations.Controller) {
	go func() {
		for {
			response, _ := reader.ReadString('\n')
			if len(response) != 1 {
				_, err := writer.Write(bytes.Trim([]byte(response), "\x00"))
				if err != nil {
					log.Error(err)
				}
				writer.Flush()
				c.Parse(response)
			}
		}
	}()
}
