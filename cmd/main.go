package main

import (
	"github.com/alideweb/shop-center-go/config"
	"github.com/alideweb/shop-center-go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	serverEnvsConfig := config.ConfigEnvs()

	db.ConnectToMongo(serverEnvsConfig.MongoUri)

	config.SetupRoutes(r)

	config.StartServer(r, serverEnvsConfig.Port)
}
