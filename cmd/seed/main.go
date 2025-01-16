package main

import (
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/infra/database"
)

const SeedersFilePath = "data/seeders/"
const SeedersDevPath = SeedersFilePath + "dev/"
const SeedersProdPath = SeedersFilePath + "prod/"

func main() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	//var path string
	//if env.AppEnv.AppEnv == "production" {
	//	path = SeedersProdPath
	//} else {
	//	path = SeedersDevPath
	//}
	//
	//validator := validator.Validator
	//uuid := uuid.UUID
	//bcrypt := bcrypt.Bcrypt

	//seedUsers()
}

func seedUsers() {
	// implement seeds
}
