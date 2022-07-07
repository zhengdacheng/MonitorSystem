package router

import (
	"github.com/gin-gonic/gin"
	"manageService/api/http/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	ruleRouter := router.Group("/rule")
	{
		// request for all rules
		ruleRouter.GET("", handler.FindALL)
		// insert a rule
		ruleRouter.POST("", handler.Insert)
	}
	hostRouter := router.Group("/host")
	{
		// request for all rules
		hostRouter.GET("", handler.FindALLHost)
		// request for hosts which aren't managed in mysql yet.
		hostRouter.GET("/manage", handler.GetUniqueHost)
		// insert a rule
		hostRouter.POST("", handler.HostInsert)
		// update a host
		hostRouter.PUT("", handler.HostUpdate)
	}
	return router
}