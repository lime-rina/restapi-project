package main

import (
	"example.com/main/db"
	"example.com/main/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //lh:8080

}
