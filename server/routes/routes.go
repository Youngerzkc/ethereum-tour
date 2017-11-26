package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bitschain/ethereum-tour/server/service"
)

func InitRoutes(router *gin.Engine, es *service.EtherService) {
	router.GET("/admin/nodeinfo", func(c *gin.Context) {
		nodeInfo, _ := es.GetNodeInfo()
		c.JSON(http.StatusOK, nodeInfo)
	})

	router.GET("/net/version", func(c *gin.Context) {
		version , _ := es.GetVersion()
		c.JSON(http.StatusOK, version)
	})
}
