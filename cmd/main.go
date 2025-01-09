package main

import (
	"github.com/alideweb/shop-center-go/config"
	"github.com/alideweb/shop-center-go/db"
	"github.com/alideweb/shop-center-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConfigEnvs()

	db.ConnectToMongo(config.ServerEnvsConfig.MongoUri)

	routes.SetupRoutes(r)

	config.StartServer(r, config.ServerEnvsConfig.Port)
}
