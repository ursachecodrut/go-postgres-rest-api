package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/user")
}
