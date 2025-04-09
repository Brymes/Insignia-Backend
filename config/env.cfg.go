package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ()

func LoadEnv() {
	_ = godotenv.Load()

}

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalln("Kindly Pass the environment variable named: ", key)
	}
	return val
}

func InitDSN(name string) string {
	var (
		dbHost = GetEnv("DB_HOST")
		dbUser = GetEnv("DB_USERNAME")
		dbPass = GetEnv("DB_PASSWORD")
		dbPort = GetEnv("DB_PORT")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s ", dbHost, dbUser, dbPass, dbPort, name)

	return dsn
}
