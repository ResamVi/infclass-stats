package infrastructure

import (
	"fmt"
	"infclass-stats/config"
	"infclass-stats/model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

// GormRepo is a Controller implementation using the Golang ORM "gorm"
type GormRepo struct {
	db *gorm.DB
}

/*
 * Activity
 */

// StoreActivity persists the activity
func (repo GormRepo) StoreActivity(activity *model.Activity) {
	repo.db.Create(activity)
}

// FindActivity prepares a list of all activities ordered by time
func (repo GormRepo) FindActivity() []model.Activity {
	result := []model.Activity{}

	err := repo.db.Order("timestamp desc").Find(&result).Error
	if err != nil {
		log.Error(err)
	}

	return result
}

// ResetActivity clears all data
func (repo GormRepo) ResetActivity() {
	repo.db.Delete(model.Activity{})
}

/*
 * Class
 */

// StoreClasses persists all classes
func (repo GormRepo) StoreClasses(classes []model.Class) {
	for _, class := range classes {
		repo.db.Create(&class)
	}
}

// ResetClass clears all data
func (repo GormRepo) ResetClass() {
	repo.db.Delete(model.Class{})
}

/*
 * Client
 */

// StoreClient persists the activity
func (repo GormRepo) StoreClient(client *model.Client) {
	repo.db.Create(client)
}

// FindClientByName returns a currently connected client with the given name
// Returns nil if name was not found
func (repo GormRepo) FindClientByName(name string) (*model.Client, error) {
	client := model.Client{}

	// Convert normal name into escaped unicode string as it is stored this way in the database
	err := repo.db.First(&client, "name = ?", fmt.Sprintf("%+q", name)).Error

	return &client, err
}

// FindClient prepares a list of players currently on server
func (repo GormRepo) FindClient() []model.Client {
	clients := []model.Client{}

	err := repo.db.Find(&clients).Error

	if err != nil {
		log.Error(err)
	}

	for i := range clients {
		// Convert escaped unicode back to readable names
		clients[i].Name, err = strconv.Unquote(clients[i].Name)
		if err != nil {
			log.Error(err)
		}
	}

	return clients
}

// DeleteClientByName removes the row with the name.
// Does nothing if row does not exist
func (repo GormRepo) DeleteClientByName(name string) {
	repo.db.Where("name = ?", fmt.Sprintf("%+q", name)).Delete(&model.Client{})
}

// ResetClient clears the table of all rows
func (repo GormRepo) ResetClient() {
	repo.db.Delete(&model.Client{})
}

// Count gets the amount of rows in the table
func (repo GormRepo) Count() int {
	return len(repo.FindClient())
}

/*
 * Map
 */

// StoreMap persists a map. Can be used to update attributes of an already persisted map.
func (repo GormRepo) StoreMap(mapp *model.Map) {
	if repo.db.NewRecord(*mapp) {
		repo.db.Create(mapp)
	} else {
		repo.db.Save(mapp) // modifies already existing row
	}
}

// FindMapByName returns a map with the given name (e.g. "infc_hardcorepit")
func (repo GormRepo) FindMapByName(name string) (*model.Map, error) {
	mapp := model.Map{}
	err := repo.db.First(&mapp, "name = ?", name).Error
	return &mapp, err
}

// FindMap prepares a list of all maps played
func (repo GormRepo) FindMap() []model.Map {
	maps := []model.Map{}

	err := repo.db.Find(&maps).Error
	if err != nil {
		log.Error(err)
	}

	return maps
}

// ResetMap clears all data
func (repo GormRepo) ResetMap() {
	repo.db.Delete(model.Map{})
}

/*
 * Player
 */

// StorePlayer persists a player. Can be used to update attributes of an already persisted player.
func (repo GormRepo) StorePlayer(player *model.Player) {
	if repo.db.NewRecord(*player) {
		repo.db.Create(player)
	} else {
		repo.db.Save(player) // modifies already existing row
	}
}

// FindPlayerByName returns a player by his name
// Returns nil if name was not found
func (repo GormRepo) FindPlayerByName(name string) (*model.Player, error) {
	p := model.Player{}

	// Converting names to escaped unicode strings.
	err := repo.db.Preload("Classes").First(&p, "name = ?", fmt.Sprintf("%+q", name)).Error

	return &p, err
}

// FindPlayer prepares a list of all players
func (repo GormRepo) FindPlayer() []model.Player {
	players := []model.Player{}

	err := repo.db.Find(&players).Error
	if err != nil {
		log.Error(err)
	}

	for i := range players {
		// Convert escaped unicode to readable uncicode
		players[i].Name, err = strconv.Unquote(players[i].Name)
		if err != nil {
			log.Error(err)
		}

		repo.db.Model(players[i]).Related(&players[i].Classes)
	}

	return players
}

// ResetPlayer clears all data
func (repo GormRepo) ResetPlayer() {
	repo.db.Delete(model.Player{})
}

// ResetPlayer24 clears all tables of daily data (suffix '24'). Done each day on midnight
func (repo GormRepo) ResetPlayer24() {
	repo.db.Model(&model.Player{}).Update("time24", 0)
	repo.db.Model(&model.Player{}).Update("score24", 0)
	repo.db.Model(&model.Player{}).Update("kill_count24", 0)
	repo.db.Model(&model.Player{}).Update("rounds_survived24", 0)
}

/*
 * Role
 */

// StoreRole persists a role. Can be used to update attributes of an already persisted player.
func (repo GormRepo) StoreRole(role *model.Role) {
	if repo.db.NewRecord(*role) {
		repo.db.Create(role)
	} else {
		repo.db.Save(role) // modifies already existing row
	}
}

// FindRoleByIndex returns a role by their index (see: model/Class.go)
func (repo GormRepo) FindRoleByIndex(index int) (*model.Role, error) {
	roles := model.Role{}
	err := repo.db.Where("name = ?", model.Classes[index]).First(&roles).Error

	return &roles, err
}

// FindRole prepares a list of all roles
func (repo GormRepo) FindRole() []model.Role {
	roles := []model.Role{}

	err := repo.db.Find(&roles).Error
	if err != nil {
		log.Error(err)
	}

	return roles
}

// ResetRole clears the table of all rows
func (repo GormRepo) ResetRole() {
	repo.db.Delete(model.Role{})
}

/*
 * modelgo
 */

// Reset clears all tables of their data. Done each week on Thursday 23:59:58 for a fresh
// series of people to reach the top ranks
func (repo GormRepo) Reset() {
	repo.db.Delete(&model.Player{})
	repo.db.Delete(&model.Class{})
	repo.db.Delete(&model.Map{})
	repo.db.Delete(&model.Activity{})
	repo.db.Delete(&model.Role{})
	repo.db.Delete(&model.Client{})
}

// NewGormDB creates a new connection to a database implemented via ORM library "gorm"
// Creates the tables if they do not exist yet and seeds the database
// with an initial set of class objects to track overall class-specific performance.
func NewGormDB() GormRepo {
	args := config.MYSQL_USER + ":" + config.MYSQL_PASSWORD + "@/" + config.MYSQL_DB
	db, err := gorm.Open("mysql", args)
	if err != nil {
		log.Panicf("Could not connect to database: %s\n%s\n", args, err.Error())
	}

	db.AutoMigrate(&model.Activity{}, &model.Player{}, &model.Class{}, &model.Client{}, &model.Role{}, &model.Map{})

	seedRoles(db)

	return GormRepo{db: db}
}

// seedRoles creates entry for all classes (Engineer, Scientist, ... Smoker, ...) if they do not exist yet
func seedRoles(db *gorm.DB) {
	role := []model.Role{}
	err := db.Find(&role).Error
	if err != nil {
		log.Error(err)
	}

	if len(role) < model.HumanLimit+1 {
		db.Delete(model.Role{})
		for i := 0; i < len(model.Classes); i++ {
			db.Create(&model.Role{ID: uint(i + 1), Name: model.Classes[i]})
		}
	}
}
