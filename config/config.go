package config

import (
	"log"
	"os"

	auth "github.com/alideweb/shop-center-go/modules/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ServerEnvsConfig struct {
	Port        string
	MongoUri    string
	MongoDbName string
}

func ConfigEnvs() *ServerEnvsConfig {
	// --> Load Envs File
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// --> Load Envs value
	serverEnvsConfig := ServerEnvsConfig{Port: os.Getenv("PORT"), MongoUri: os.Getenv("MONGO_URI"), MongoDbName: os.Getenv("MONGO_DB_NAME")}

	return &serverEnvsConfig
}

func SetupRoutes(engine *gin.Engine) {
	auth.Routes(engine)
}

func StartServer(engine *gin.Engine, port string) {
	log.Printf("Server listening on port %s", port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Error starting server", err)
	}
}
