package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jester/api/probes"
	"jester/api/v1"
	"jester/databases/db"
	"jester/databases/postgres"
	"jester/docs"
	"jester/logger"
)

// @title           Jester API
// @version         1.0
// @description     Store and reuse your CLI commands.
// @termsOfService  https://jesthub.sh/terms/

// @contact.name   Petar Nikolovski
// @contact.url    https://jesthub.sh/support
// @contact.email  petar.nikolovski@protonmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  JestHub Documentation
// @externalDocs.url          https://jesthub.sh/docs
func serve() {
	router := gin.Default()

	prometheus := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
		ginprom.Ignore(
			"/metrics",
			"/ready",
			"/healthz",
		),
	)
	router.Use(prometheus.Instrument())

	docs.SwaggerInfo.BasePath = "/api/v1"

	probes.Routes(router)
	v1.Routes(router)

	// Look into https://github.com/swaggo/swag/issues/1568
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func loadDotEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Log.Warn(err)
	}
}

func main() {
	loadDotEnvFile()
	logger.SetupLogger()

	database, err := postgres.Connect()
	db.SetDB(database)

	sqlDB, err := database.DB()
	if err != nil {
		logger.Log.Panic(err)
	}
	defer sqlDB.Close()

	postgres.InitDatabase(database)

	serve()
}
