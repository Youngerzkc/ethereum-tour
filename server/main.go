package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bitschain/ethereum-tour/server/routes"
	"github.com/bitschain/ethereum-tour/server/service"
)

func main() {

	router := gin.Default()
	es := service.NewEtherService()
	routes.InitRoutes(router, es)
	router.Run(":8080")

}
