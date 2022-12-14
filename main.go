package main

import (
	"github.com/dawson1096/cointracker-project/src"
	"github.com/gin-gonic/gin"
)

func main() {
	orc := &src.Orchestrator{}
	orc.Init()
	router := gin.Default()
	router.POST("/api/:address/add", orc.AddAddress)
	router.DELETE("/api/:address/remove", orc.RemoveAddress)
	router.POST("/api/:address/sync", orc.SyncAddress)
	router.GET("/api/listAddresses", orc.ListAddresses)
	router.GET("/api/:address", orc.GetAddress)

	router.Run("localhost:8080")
}
