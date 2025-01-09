package config

import (
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type TServerEnvsConfig struct {
	Port        string
	MongoUri    string
	MongoDbName string
	JwtSecret   string
}

var (
	ServerEnvsConfig *TServerEnvsConfig
	once             sync.Once
)

func ConfigEnvs() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}

		ServerEnvsConfig = &TServerEnvsConfig{Port: os.Getenv("PORT"), MongoUri: os.Getenv("MONGO_URI"), MongoDbName: os.Getenv("MONGO_DB_NAME"), JwtSecret: os.Getenv("JWT_SECRET")}
	})
}

func StartServer(engine *gin.Engine, port string) {
	log.Printf("Server listening on port %s", port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Error starting server", err)
	}
}
