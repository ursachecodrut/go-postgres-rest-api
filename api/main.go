package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ursachecodrut/api/config"
)

func main() {
	port := config.LoadFromEnv("PORT")
	db_uri := config.LoadFromEnv("DB_URI")

	config.ConnectDB(db_uri)
	// handler := controller.New(db)

	router := gin.Default()

	router.Run(":" + port)
}
