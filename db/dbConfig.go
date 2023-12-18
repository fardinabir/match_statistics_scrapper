package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDb() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	password := viper.GetString("database.password")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=disable",
		host, port, password, user, dbname,
	)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error", err)
	}

	fmt.Println("Successfully Connected Database :D")

	//Db.AutoMigrate(models.MatchStatUrl{})
}
