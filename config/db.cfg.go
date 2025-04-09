package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DBClient *gorm.DB
)

func InitDB() {
	DBClient = InitPostgres(GetEnv("DB_DATABASE"))
}

func InitPostgres(db string) *gorm.DB {
	var err error

	DbConn, err := gorm.Open(postgres.Open(InitDSN(db)), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Println("Error Connecting to Database. Kindly set accurate Database environment variables")
		log.Fatal(err)
	}

	return DbConn
}
