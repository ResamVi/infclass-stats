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

package operations

import (
	"infclass-stats/model"
	"regexp"
	"strconv"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

/*
 * Parse econ log statements.
 * Contains logic to evaluate/act upon any incoming logs coming from the econ.
 *
 * Abstract:	"[TYPE]: INFORMATION"
 * Into:			result[1]: "CATEGORY"
 *						result[2]: "CONTENT"
 *
 * Example:	  "[register]: ERROR: the master server..."
 * Into:			result[1]: "register",
 *						result[2]: "ERROR: the master server..."
 */
func (c Controller) Parse(input string) {
	reg := regexp.MustCompile(`\[(\w+)\]: (.+)`)
	result := reg.FindStringSubmatch(input)
	category, content := result[1], result[2]

	if strings.HasPrefix(content, "team_join") {
		// Sample: "team_join player='0:nameless tee' team=0"
		reg := regexp.MustCompile(`team_join player='\d+:(.+)' m?_?[tT]eam=(-?\d)`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Player: " + result[1])
		log.Debug("Team: " + result[2])

		c.Enter(result[1], -1)

	} else if strings.HasPrefix(content, "(") {
		// Sample: "(#00) Player        : [antispoof=1] [login=0] [level=2] [ip=192.168.0.1:32596]"
		reg := regexp.MustCompile(`\(#(\d+)\) (.+): (.+)`)
		result := reg.FindStringSubmatch(content)
		log.Debug(result)
		log.Debug("ClientID: " + result[1])
		log.Debug("Player: " + strings.TrimSpace(result[2]))

		clientID, err := strconv.ParseInt(result[1], 10, 0)
		if err != nil {
			log.Error(err)
		}

		c.Enter(strings.TrimSpace(result[2]), clientID)

	} else if strings.HasPrefix(content, "leave player") {
		// Sample: "leave player player='0:nameless tee'
		reg := regexp.MustCompile(`leave player='\d+:(.+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Left: " + result[1])

		c.Leave(result[1])

	} else if strings.HasPrefix(content, "change_name") {
		// Sample: "change_name previous='Armadillop' now='Armadillo'"
		reg := regexp.MustCompile(`change_name previous='(.+)' now='(.+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Previous name: " + result[1])
		log.Debug("Current name: " + result[2])

		c.Leave(result[1])
		c.Enter(result[2], -1)

	} else if strings.HasPrefix(content, "kill") {
		// Sample: "kill killer='Armadillo' victim='nameless tee' weapon=2"
		reg := regexp.MustCompile(`kill killer='(.+)' victim='(.+)' weapon=(-?\d+)`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Killer: " + result[1])
		log.Debug("Killed: " + result[2])
		log.Debug("Weapon: " + result[3])

		// Do not count kill by leaving
		if result[3] != "-3" {
			c.AddKill(result[1])
		}
	} else if strings.HasPrefix(content, "infected") {
		// Sample: "infected victim='Armadillo' duration='10'"
		reg := regexp.MustCompile(`infected victim='(.+)' duration='(\d+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Victim: " + result[1])
		log.Debug("Duration: " + result[2])

		duration, err := strconv.ParseInt(result[2], 10, 0)
		if err != nil {
			log.Error(err)
		}

		c.Infect(result[1], int(duration))

	} else if strings.HasPrefix(content, "round_end too few players") {
		// Sample: "round_end too few players round='4 of 8'"
		return

	} else if strings.HasPrefix(content, "round_end") {

		// Extract:		"round_end winner='human' survivors='2' duration='300' round='3 if 5'"
		reg := regexp.MustCompile(`round_end winner='(\w+)' survivors='(\d+)' duration='(\d+)' round='(\d+) of (\d+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Winner: " + result[1])
		log.Debug("Survivors: " + result[2])
		log.Debug("Duration: " + result[3])
		log.Debug("Current Round: " + result[4])
		log.Debug("Final Round: " + result[5])

		duration, err := strconv.ParseInt(result[3], 10, 0)
		if err != nil {
			log.Error(err)
		}

		// Increment human/zombie wins and update highest score records
		c.Winner(result[1], int(duration))
		c.CheckRecords()

		// After final round print summary and print rank
		if result[4] == result[5] {
			c.Summary()
		}

	} else if strings.HasPrefix(content, "score") {
		// Extract:		"score player='Money' amount='-10'"
		reg := regexp.MustCompile(`score player='(.+)' amount='(-?\d+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Player: " + result[1])
		log.Debug("Score: " + result[2])

		score, err := strconv.ParseInt(result[2], 10, 0)
		if err != nil {
			log.Error(err)
		}

		c.AddScore(result[1], int(score/10))

	} else if strings.HasPrefix(content, "choose_class") {
		// Sample: "choose_class player='nameless tee' class='2'"
		reg := regexp.MustCompile(`choose_class player='(.+)' class='(\d+)'`)
		result := reg.FindStringSubmatch(content)

		class, err := strconv.Atoi(result[2])
		if err != nil {
			log.Debug("Conversion failed")
		}

		log.Debug("Player: " + result[1])
		log.Debugf("Class: %d", (class - 2))

		// Class 0 means this class was reset
		if class == 0 {
			c.SetClass(result[1], model.Unknown)
		} else {
			c.SetClass(result[1], class-2) // Mercenary is class 2 but we start at 0
		}

	} else if strings.HasPrefix(content, "infc_") {
		// Sample: "infc_hardcorepit"
		log.Debug("New Map: " + content)

		c.ChangeMap(content)
		time.AfterFunc(5*time.Second, c.Rcon.Status)
		time.AfterFunc(10*time.Second, c.Rank)

	} else if strings.HasPrefix(content, "rotating map to") {
		// Sample: "rotating map to infc_newdust"
		reg := regexp.MustCompile(`rotating map to (.+)`)
		result := reg.FindStringSubmatch(content)

		log.Debug("New Map: " + result[1])

		c.ChangeMap(result[1])
		time.AfterFunc(5*time.Second, c.Rcon.Status)
		time.AfterFunc(10*time.Second, c.Rank)

	} else if strings.HasPrefix(content, "survived") {
		// Sample: "survived player='(1)ArmadilIo'"
		reg := regexp.MustCompile(`survived player='(.+)'`)
		result := reg.FindStringSubmatch(content)

		log.Debug("Survivor: " + result[1])

		c.SurviveRound(result[1])
	} else if strings.HasPrefix(content, "start round") {
		log.Debug("New round: Checking records and resetting all scores")
		c.NewRound()
	} else if (category == "chat" || category == "teamchat") && strings.ContainsAny(content, "!") { // chat commands

		// Sample: "1:-2:(1)ArmadilIo: !best"
		reg := regexp.MustCompile(`^(\d+):-?\d+:(.+): !([a-z]+)\s?(.+)?$`)
		result := reg.FindStringSubmatch(content)

		if len(result) < 3 {
			log.Warning("Unintended use of '!': " + content)
			return
		}

		clientID, err := strconv.ParseInt(result[1], 10, 0)
		if err != nil {
			log.Error(err)
		}

		if len(result) == 4 {
			log.Debug("ClientID: " + result[1])
			log.Debug("Sender: " + result[2])
			log.Debug("Command: " + result[3])

			c.Command(clientID, result[2], result[3], "")
		}

		if len(result) == 5 {
			log.Debug("ClientID: " + result[1])
			log.Debug("Sender: " + result[2])
			log.Debug("Command: " + result[3])
			log.Debug("Argument: " + result[4])

			c.Command(clientID, result[2], result[3], result[4])
		}

	} else {
		log.Debug("Could not decide: " + content)
		log.Debug(category)
	}

}

// Command is executed when a !command was typed in the chat. Parses the command and decides what logic to execute
func (c Controller) Command(clientID int64, sender string, cmd string, arg string) {

	if cmd == "picks" {
		switch strings.ToLower(arg) {
		case "ninja":
			fallthrough
		case "engineer":
			fallthrough
		case "soldier":
			fallthrough
		case "scientist":
			fallthrough
		case "biologist":
			fallthrough
		case "medic":
			fallthrough
		case "hero":
			fallthrough
		case "mercenary":
			fallthrough
		case "sniper":
			fallthrough
		case "looper":
			for k, className := range model.Classes {
				if strings.ToLower(className) == strings.ToLower(arg) {
					players := c.PlayerRepository.FindPlayer()
					c.PrintRanking(clientID, ClassPicks(players, k))
				}
			}
		}
	}

	if cmd == "kills" {
		switch strings.ToLower(arg) {
		case "ninja":
			fallthrough
		case "engineer":
			fallthrough
		case "soldier":
			fallthrough
		case "scientist":
			fallthrough
		case "biologist":
			fallthrough
		case "medic":
			fallthrough
		case "hero":
			fallthrough
		case "mercenary":
			fallthrough
		case "sniper":
			fallthrough
		case "looper":
			for k, className := range model.Classes {
				if strings.ToLower(className) == strings.ToLower(arg) {
					players := c.PlayerRepository.FindPlayer()
					c.PrintRanking(clientID, ClassKills(players, k))
				}
			}
		}
	}

	if cmd == "record" {
		switch strings.ToLower(arg) {
		case "ninja":
			fallthrough
		case "engineer":
			fallthrough
		case "soldier":
			fallthrough
		case "scientist":
			fallthrough
		case "biologist":
			fallthrough
		case "medic":
			fallthrough
		case "hero":
			fallthrough
		case "mercenary":
			fallthrough
		case "sniper":
			fallthrough
		case "looper":
			for k, className := range model.Classes {
				if strings.ToLower(className) == strings.ToLower(arg) {
					players := c.PlayerRepository.FindPlayer()
					c.PrintRanking(clientID, ClassRecords(players, k))
				}
			}
		}
	}

	if cmd == "today" {
		switch arg {
		case "kills":
			c.PrintRanking(clientID, DailyKills(c.PlayerRepository.FindPlayer()))
		case "active":
			c.PrintRanking(clientID, DailyActiveTime(c.PlayerRepository.FindPlayer()))
		case "survivors":
			c.PrintRanking(clientID, DailySurvivals(c.PlayerRepository.FindPlayer()))
		}
	}

	switch strings.ToLower(cmd) {
	case "help":
		c.CommandHelp(clientID)
	case "kills":
		c.PrintRanking(clientID, WeeklyKills(c.PlayerRepository.FindPlayer()))
	case "active":
		c.PrintRanking(clientID, WeeklyActiveTime(c.PlayerRepository.FindPlayer()))
	case "survivors":
		c.PrintRanking(clientID, WeeklySurvivals(c.PlayerRepository.FindPlayer()))
	case "best":
		c.PrintRanking(clientID, BestPlayers(c.PlayerRepository.FindPlayer()))
	case "player":
		c.CommandMe(clientID, arg)
	case "me":
		c.CommandMe(clientID, sender)
	case "wins":
		c.CommandWins(clientID)
	case "ninja":
		fallthrough
	case "engineer":
		fallthrough
	case "soldier":
		fallthrough
	case "scientist":
		fallthrough
	case "biologist":
		fallthrough
	case "medic":
		fallthrough
	case "hero":
		fallthrough
	case "mercenary":
		fallthrough
	case "sniper":
		fallthrough
	case "looper":
		for k, className := range model.Classes {
			if strings.ToLower(className) == strings.ToLower(cmd) {
				c.PrintRanking(clientID, BestClass(c.PlayerRepository.FindPlayer(), k))
			}
		}
	}
}
