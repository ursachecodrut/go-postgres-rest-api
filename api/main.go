package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ursachecodrut/api/config"
)

func main() {
	router := gin.Default()

	port := config.LoadFromEnv("PORT")
	db_uri := config.LoadFromEnv("DB_URI")
	config.ConnectDB(db_uri)

	router.Run(":" + port)
}
