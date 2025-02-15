package main

import (
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/db"
	"C/Users/anura/OneDrive/Documents/GitHub/MyProject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()      // this func already return pointer
	routes.RegisterRoute(server) // now sending tis pointer to our routes
	server.Run(":8081")          // localhost:8080

}
