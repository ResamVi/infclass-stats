package infrastructure

import (
	"infclass-stats/model"
	"sort"
)

// ServerState is the state of the server.
//
// This is a implementation of StatusRepository that does not use a database.
// So State contains information that are not shared between multiple instances of infclass-stats.
// Reason for this is because information change in real-time and would cause too many queries to the database)
type ServerState struct {
	players    *[]*model.Status
	currentMap *string
}

// StoreStatus persists a status.
func (s ServerState) StoreStatus(status *model.Status) {
	if !s.hasEntered(status.Name) {
		*s.players = append(*s.players, status)
	}
}

// Sort and return all current players by score
func (s ServerState) Sort() []*model.Status {
	result := *s.players

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}

// FindStatusByName returns a status of the associated name
func (s ServerState) FindStatusByName(name string) *model.Status {
	for _, player := range *s.players {
		if player.Name == name {
			return player
		}
	}
	return nil
}

func (s ServerState) FindStatus() []*model.Status {
	return *s.players
}

// ResetStatus clears the table of all rows
func (s ServerState) ResetStatus() {
	players := make([]*model.Status, 0)
	*s.players = players
}

func (s ServerState) SetMap(newMap string) {
	*s.currentMap = newMap
}

func (s ServerState) GetMap() string {
	return *s.currentMap
}

// hasEntered determines whether or not the player has entered the server before in the past 5 rounds of the map
func (s *ServerState) hasEntered(name string) bool {
	for _, player := range *s.players {
		if player.Name == name {
			return true
		}
	}
	return false
}

// NewServerState creates a new server State persisting statuses of all players
// and the current map played
func NewServerState() ServerState {
	players := make([]*model.Status, 0)
	current := "UNKNOWN"
	return ServerState{players: &players, currentMap: &current}
}
