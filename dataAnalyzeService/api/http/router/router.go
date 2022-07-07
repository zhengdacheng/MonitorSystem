package router

import (
	"dataAnalyzeService/api/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	index := router.Group("")
	{
		index.GET("/data", handler.DataQuery)
		index.GET("/location", handler.DataQueryByLocation)
	}
	return router
}
