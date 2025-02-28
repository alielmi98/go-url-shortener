package migrations

import (
	"log"

	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/data/db"
	"github.com/alielmi98/go-url-shortener/data/models"

	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDb()

	createTables(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// Basic
	tables = addNewTable(database, models.ShortURL{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Migration, err.Error())
	}
	log.Printf("Caller:%s Level:%s Msg:%s", constants.Postgres, constants.Migration, "tables created")
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func Down_1() {

}
