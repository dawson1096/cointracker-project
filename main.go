package main

import (
	"cointracker/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	orc := &internal.Orchestrator{}
	orc.Init()
	router := gin.Default()
	router.POST("/api/addAddress", orc.AddAddress)
	router.DELETE("/api/:address/remove", orc.RemoveAddress)
	router.POST("/api/:address/sync", orc.SyncAddress)
	router.GET("/api/:address", orc.GetAddress)
	router.GET("/api/listAddresses", orc.ListAddresses)

	router.Run("localhost:8080")
}
