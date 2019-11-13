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

package infrastructure

import (
	"encoding/json"
	"infclass-stats/config"
	"infclass-stats/operations"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebserviceHandler struct {
	Controller operations.Controller
}

// Serve will initialize the web server to serve websocket connections
// TODO: Dont make Serve create the handler himself
func Serve(controller operations.Controller) {
	webserviceHandler := WebserviceHandler{Controller: controller}
	
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.getUpdates(res, req)
	})

	go func() {
		if config.USE_SSL {
			err := http.ListenAndServeTLS(":8000", config.CERT_FILE, config.KEY_FILE, nil)
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := http.ListenAndServe(":8000", nil)
			if err != nil {
				log.Panic(err)
			}
		}
	}()
}

func (handler WebserviceHandler) getUpdates(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		myJson, err := json.Marshal(handler.Controller.GetData())
		log.Debugf("Current size = %d\n", len(myJson))
		if err != nil {
			log.Error(err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, myJson)
		if err != nil {
			log.Error(err)
			break
		}

		time.Sleep(5 * time.Second)
	}
	conn.Close()
}
