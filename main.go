package main

import (
	"project/config"
	"project/migrate"
	"project/routes"
)

func main() {
	config.InitDB()
	migrate.AutoMigrate()

	e := routes.New()
	e.Start(":8800")
}
