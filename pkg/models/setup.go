package models
import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {

	dbUser 		:= viper.GetString("database.username")
	dbPassword 	:= viper.GetString("database.password")
	dbName 		:= viper.GetString("database.database")
	dbHost 		:= viper.GetString("database.addr")
	dbPort 		:= viper.GetString("database.port")

	connString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	database, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}