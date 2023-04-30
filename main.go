package main

import (
	"github.com/gin-gonic/gin"
	"jester/api/probes"
	"jester/api/v1"
	"jester/databases/postgres"
	"log"
)

func serve() {
	router := gin.Default()

	probes.Routes(router)
	v1.Routes(router)

	router.Run()
}

func main() {
	db, err := postgres.Connect()

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	defer sqlDB.Close()

	postgres.InitDatabase(db)

	serve()
}
