package main

import (
	"voicerrs-cache-layer/api"
	"voicerrs-cache-layer/api/config"
	"voicerrs-cache-layer/api/server"
	"voicerrs-cache-layer/db"

	"github.com/joho/godotenv"
)

func init() {
    godotenv.Load()
}

func main() {
	config.Init()

	db.Create()
	api.AutoMigrate()

	server.Run()
}
