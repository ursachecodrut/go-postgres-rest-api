package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ursachecodrut/api/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/env/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	db.Init(dbUrl)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)

}
