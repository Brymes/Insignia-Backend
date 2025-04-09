package models

import (
	"Insignia-Backend/config"
	"log"
)

func MigrateModels() {
	// Create custom type for Transactions
	uuidQuery := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	config.DBClient.Exec(uuidQuery)

	// Temporary Migration for dev
	err := config.DBClient.AutoMigrate(&BoilerBooking{})
	if err != nil {
		log.Fatalln(err)
	}
}
