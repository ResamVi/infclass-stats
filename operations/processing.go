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

// Process.go contains any logic to calculate, prepare and bundle the data to serve over the web.
import (
	"infclass-stats/model"
	"math"
	"sort"
	"time"
)

// Data contains all the data that is polled every 5 seconds and sent as JSON object
// to the website via a websocket connection
type Data struct {
	CurrentMap      string
	Online          []model.Client
	Activities      []model.Activity
	Maps            []model.Map
	HumanWins       int
	ZombieWins      int
	ClassSurvivals  []int
	ClassAliveTime  []float32
	Mvps            Dict
	DailyActive     Dict
	DailyKills      Dict
	DailySurvivals  Dict
	WeeklyKills     Dict
	WeeklyActive    Dict
	WeeklySurvivals Dict
	ClassPicks      ClassBundle
	ClassKills      ClassBundle
	ClassRatio      ClassBundle
	ClassRecords    ClassBundle
	ClassScores     ClassBundle
	ClassBest       ClassBundle
}

// Entry is a single element of a Dictionary
type Entry struct {
	Key   string
	Value float32
}

// Dict is intended to be used to display lists of (player, data) entries
type Dict []Entry

// ClassBundle gives information on a single stat (e.g. kill counts) of all players on all classes
type ClassBundle struct {
	Engineer  Dict
	Soldier   Dict
	Scientist Dict
	Biologist Dict
	Medic     Dict
	Hero      Dict
	Ninja     Dict
	Mercenary Dict
	Sniper    Dict
	Looper    Dict
}

// GetData processes and returns an object containing all relevant information that will displayed to the webpage
func (c Controller) GetData() Data {
	start := time.Now()

	// All raw data is queried once and then passed on to the functions
	online := c.ClientRepository.FindClient()
	activities := c.ActivityRepository.FindActivity()
	maps := c.MapRepository.FindMap()
	players := c.PlayerRepository.FindPlayer()
	classes := c.RoleRepository.FindRole()

	picks := ClassBundle{
		Mercenary: ClassPicks(players, 0),
		Medic:     ClassPicks(players, 1),
		Hero:      ClassPicks(players, 2),
		Engineer:  ClassPicks(players, 3),
		Soldier:   ClassPicks(players, 4),
		Ninja:     ClassPicks(players, 5),
		Sniper:    ClassPicks(players, 6),
		Scientist: ClassPicks(players, 7),
		Biologist: ClassPicks(players, 8),
		Looper:    ClassPicks(players, 9)}

	kills := ClassBundle{
		Mercenary: ClassKills(players, 0),
		Medic:     ClassKills(players, 1),
		Hero:      ClassKills(players, 2),
		Engineer:  ClassKills(players, 3),
		Soldier:   ClassKills(players, 4),
		Ninja:     ClassKills(players, 5),
		Sniper:    ClassKills(players, 6),
		Scientist: ClassKills(players, 7),
		Biologist: ClassKills(players, 8),
		Looper:    ClassKills(players, 9)}

	ratio := ClassBundle{
		Mercenary: ClassRatio(players, 0),
		Medic:     ClassRatio(players, 1),
		Hero:      ClassRatio(players, 2),
		Engineer:  ClassRatio(players, 3),
		Soldier:   ClassRatio(players, 4),
		Ninja:     ClassRatio(players, 5),
		Sniper:    ClassRatio(players, 6),
		Scientist: ClassRatio(players, 7),
		Biologist: ClassRatio(players, 8),
		Looper:    ClassRatio(players, 9)}

	records := ClassBundle{
		Mercenary: ClassRecords(players, 0),
		Medic:     ClassRecords(players, 1),
		Hero:      ClassRecords(players, 2),
		Engineer:  ClassRecords(players, 3),
		Soldier:   ClassRecords(players, 4),
		Ninja:     ClassRecords(players, 5),
		Sniper:    ClassRecords(players, 6),
		Scientist: ClassRecords(players, 7),
		Biologist: ClassRecords(players, 8),
		Looper:    ClassRecords(players, 9)}

	scores := ClassBundle{
		Mercenary: ClassScores(players, 0),
		Medic:     ClassScores(players, 1),
		Hero:      ClassScores(players, 2),
		Engineer:  ClassScores(players, 3),
		Soldier:   ClassScores(players, 4),
		Ninja:     ClassScores(players, 5),
		Sniper:    ClassScores(players, 6),
		Scientist: ClassScores(players, 7),
		Biologist: ClassScores(players, 8),
		Looper:    ClassScores(players, 9)}

	best := ClassBundle{}

	if len(BestClass(players, 0)) >= 4 {

	best = ClassBundle{
		Mercenary: BestClass(players, 0)[:5],
		Medic:     BestClass(players, 1)[:5],
		Hero:      BestClass(players, 2)[:5],
		Engineer:  BestClass(players, 3)[:5],
		Soldier:   BestClass(players, 4)[:5],
		Ninja:     BestClass(players, 5)[:5],
		Sniper:    BestClass(players, 6)[:5],
		Scientist: BestClass(players, 7)[:5],
		Biologist: BestClass(players, 8)[:5],
		Looper:    BestClass(players, 9)[:5]}
	}

	data := Data{
		Online:          online,
		Activities:      activities,
		Maps:            maps,
		ClassSurvivals:  ClassSurvivals(classes),
		ClassAliveTime:  ClassAverageSurvival(classes),
		CurrentMap:      c.StatusRepository.GetMap(),
		Mvps:            BestPlayers(players),
		DailyActive:     DailyActiveTime(players),
		DailyKills:      DailyKills(players),
		DailySurvivals:  DailySurvivals(players),
		WeeklyKills:     WeeklyKills(players),
		WeeklyActive:    WeeklyActiveTime(players),
		WeeklySurvivals: WeeklySurvivals(players),
		HumanWins:       HumanWins(maps),
		ZombieWins:      ZombieWins(maps),
		ClassPicks:      picks,
		ClassKills:      kills,
		ClassRatio:      ratio,
		ClassRecords:    records,
		ClassScores:     scores,
		ClassBest:       best}

	end := time.Now()
	log.Debugf("Time loaded = %v\n", end.Sub(start))

	return data
}

// BestPlayers create a sorted  list of (name, rating) pairs that dictates the overall best score
// in comparison to the average
//
// The formular is compares your kill-count, survival-count to the average kill-count and survival-count
// and maps the ratio to a log2 function
func BestPlayers(players []model.Player) Dict {
	result := make(Dict, 0)

	avgKills, avgSurvived := averageKills(players), averageSurvivals(players)

	if avgKills == 0 || avgSurvived == 0 {
		for _, player := range players {
			result = append(result, Entry{Key: player.Name, Value: float32(0)})
		}
	} else {
		for _, player := range players {
			kills, survived := player.KillCount, player.RoundsSurvived

			x, y := float64(kills)/float64(avgKills), float64(survived)/float64(avgSurvived)
			score := math.Floor(1000*math.Log2(x+1) + 1000*math.Log2(y+1))

			result = append(result, Entry{Key: player.Name, Value: float32(score)})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// BestClass prints the ratings of the top 10 players on that class
func BestClass(players []model.Player, class int) Dict {
	ratios := ClassRatio(players, class)

	averagePicks := calculateAverage(ClassPicks(players, class))
	averageKills := calculateAverage(ClassKills(players, class))
	averageRecord := calculateAverage(ClassRecords(players, class))
	averageRatio := calculateAverage(ratios)
	averageScore := calculateAverage(ClassScores(players, class))

	result := make(Dict, 0)

	for _, player := range players {

		kills := player.Classes[class].Kills
		picks := player.Classes[class].Picks
		record := player.Classes[class].Record
		score := player.Classes[class].Score
		ratio := getRatio(player, ratios)

		var a, b, c, d, e float64

		if averageKills > 0 {
			a = float64(kills) / float64(averageKills)
		} else {
			a = 1
		}

		if averagePicks > 0 {
			b = float64(picks) / float64(averagePicks)
		} else {
			b = 1
		}

		if averageRecord > 0 {
			c = float64(record) / float64(averageRecord)
		} else {
			c = 1
		}

		if averageRatio > 0 {
			d = float64(ratio) / float64(averageRatio)
		} else {
			d = 1
		}

		if averageScore > 0 {
			e = float64(score) / float64(averageScore)
		} else {
			e = 1
		}

		rating := math.Floor(100*math.Log2(a+1) + 100*math.Log2(b+1) + 100*math.Log2(c+1) + 100*math.Log2(d+1) + 400*math.Log2(e+1))

		result = append(result, Entry{Key: player.Name, Value: float32(rating)})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// GetRank returns the index in the list of BestPlayers and the rating
func GetRank(name string, players []model.Player) (int, float32) {
	dict := BestPlayers(players)

	for i, entry := range dict {
		if entry.Key == name {
			return (i + 1), entry.Value
		}
	}

	return -1, 0
}

// GetClassRank returns a list of (classname, rankposition)
func GetClassRank(name string, players []model.Player) Dict {

	result := make(Dict, 0)

	// For each class
	for class := 0; class < model.HumanLimit; class++ {
		// Search for player in class ranking
		dict := BestClass(players, class)
		for position, entry := range dict {
			if entry.Key == name {
				result = append(result, Entry{Key: model.Classes[class], Value: float32(position + 1)})
				break
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	return result
}

// WeeklySurvivals creates a sorted list of (name, amount) pairs where amount is the total survivals in a week interval
func WeeklySurvivals(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].RoundsSurvived > players[j].RoundsSurvived
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.RoundsSurvived)})
	}

	return result
}

// DailySurvivals creates a sorted list of (name, amount) pairs where amount is the total survivals today
func DailySurvivals(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].RoundsSurvived24 > players[j].RoundsSurvived24
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.RoundsSurvived24)})
	}

	return result
}

// ClassSurvivals creates a sorted list of (name, amount) pairs where amount is the total survivals players have
// achieved playing this class
func ClassSurvivals(classes []model.Role) []int {
	result := make([]int, 0)

	for i, class := range classes {
		if i < model.HumanLimit {
			result = append(result, class.Survived)
		}
	}

	return result
}

// ClassAverageSurvival creates a sorted list of (name, avg) pairs where avg is the time alive on that class divided by the times picked
func ClassAverageSurvival(classes []model.Role) []float32 {
	result := make([]float32, 0)

	for i, class := range classes {
		if i < model.HumanLimit {
			if class.PlayCount > 0 {
				result = append(result, float32(class.AliveTime)/float32(class.PlayCount))
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}

// DailyActiveTime creates a sorted list of (name, amount) pairs where amount is the time spent playing on the server today
func DailyActiveTime(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Time24 > players[j].Time24
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Time24)})
	}

	return result
}

// WeeklyActiveTime creates a sorted list of (name, amount) pairs where amount is the time spent playing on the server
func WeeklyActiveTime(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Time > players[j].Time
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Time)})
	}

	return result
}

// DailyKills creates a sorted list of (name, amount) pairs where amount is the total kills today
func DailyKills(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].KillCount24 > players[j].KillCount24
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.KillCount24)})
	}

	return result
}

// WeeklyKills creates a sorted list of (name, amount) pairs where amount is the total kills
func WeeklyKills(players []model.Player) Dict {
	sort.Slice(players, func(i, j int) bool {
		return players[i].KillCount > players[j].KillCount
	})

	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.KillCount)})
	}

	return result
}

// ClassPicks creates a sorted list of (name, amount) pairs where amount is the total amount this class has been picked
func ClassPicks(players []model.Player, class int) Dict {
	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Classes[class].Picks)})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// ClassKills creates a sorted list of (name, amount) pairs where amount is the total amount of kills on this class
func ClassKills(players []model.Player, class int) Dict {
	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Classes[class].Kills)})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// ClassRatio creates a sorted list of (name, ratio) pairs where ratio is the total score divided by total picks on that class
func ClassRatio(players []model.Player, class int) Dict {
	result := make(Dict, 0)

	for _, player := range players {

		score, pick := player.Classes[class].Score, player.Classes[class].Picks

		if pick < 2 {
			result = append(result, Entry{Key: player.Name, Value: 0})
		} else {
			result = append(result, Entry{Key: player.Name, Value: float32(score) / float32(pick)})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// ClassRecords creates a sorted list of (name, score) pairs where score is the highest score gained in a round on that class
func ClassRecords(players []model.Player, class int) Dict {
	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Classes[class].Record)})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// ClassScores creates a sorted list of (name, score) pairs where score is the total score gained during all rounds on this class
func ClassScores(players []model.Player, class int) Dict {
	result := make(Dict, 0)

	for _, player := range players {
		result = append(result, Entry{Key: player.Name, Value: float32(player.Classes[class].Score)})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

// HumanWins returns the amount of rounds humans survived
func HumanWins(maps []model.Map) int {
	result := 0
	for _, m := range maps {
		result += m.Human
	}
	return result
}

// ZombieWins returns the amount of rounds all humans were killed
func ZombieWins(maps []model.Map) int {
	// calculate sum
	result := 0
	for _, m := range maps {
		result += m.Zombie
	}

	return result
}

// helper methods -------
func averageKills(players []model.Player) float32 {
	sum := 0
	for _, player := range players {
		sum += player.KillCount
	}

	if len(players) > 0 {
		return float32(sum) / float32(len(players))
	}

	return 0
}

func averageSurvivals(players []model.Player) float32 {
	sum := 0
	for _, player := range players {
		sum += player.RoundsSurvived
	}

	if len(players) != 0 {
		return float32(sum) / float32(len(players))
	}

	return 0
}
