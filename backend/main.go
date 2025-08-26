package main

import (
	"github.com/gin-gonic/gin"
	"go_test/config"
	"go_test/db"
	"go_test/routes"
)

func main() {

	// load config
	conf := config.LoadConfig()

	db.InitDB(conf) // db initialization

	server := gin.Default() // initiate go gin server

	// routes register
	routes.RegisterRoutes(server)

	server.Run(conf.Port) // run go gin server

}
