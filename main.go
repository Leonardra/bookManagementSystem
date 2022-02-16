package main

import (
	"bookManagementSystem/configs"
	"bookManagementSystem/routes"
	"github.com/gin-gonic/gin"
)
func main() {

	router := gin.Default()
	configs.ConnectDB()

	routes.AuthorRouter(router)
	routes.BookRouter(router)
	err := router.Run()

	if err != nil {
		return
	}
}
