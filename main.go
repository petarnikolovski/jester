package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jester/api/probes"
	"jester/api/v1"
	"jester/databases/postgres"
	"jester/docs"
	"log"
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

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func serve() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	probes.Routes(router)
	v1.Routes(router)

	// Look into https://github.com/swaggo/swag/issues/1568
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
