package main

import (
	"Insignia-Backend/api"
	"Insignia-Backend/config"
	"Insignia-Backend/models"
)

func init() {
	config.LoadEnv()
	config.InitDB()
	models.MigrateModels()
}

func main() {
	api.Server()
}
