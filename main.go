package main

import (
	"github.com/gin-gonic/gin"
	"go_test/db"
	"go_test/routes"
)

func main() {

	db.InitDB() // db initialization

	server := gin.Default() // initiate go gin server

	// routes register
	routes.RegisterRoutes(server)

	server.Run(":8080") // run go gin server

}
