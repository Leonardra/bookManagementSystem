package routes

import (
	"bookManagementSystem/controllers"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine){
	router.POST("/", controllers.CreateBook())
}