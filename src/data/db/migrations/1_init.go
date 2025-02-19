package migrations

import (
	"log"

	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/data/db"
	"github.com/alielmi98/go-url-shortener/data/models"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDb()

	createTables(database)
	createUserDefaultInformation(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// User
	tables = addNewTable(database, models.User{}, tables)

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

func createUserDefaultInformation(database *gorm.DB) {

	u := models.User{Username: constants.DefaultUserName, Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u)

}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
	}
}

func Down_1() {

}
