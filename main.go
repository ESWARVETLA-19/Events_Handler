package main

import (
	
	"rest_api/project/db"
	"rest_api/project/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	db.InitDB()
	server := gin.Default()


	routes.RegisterRoutes(server)
	server.Run(":8080") 

}


