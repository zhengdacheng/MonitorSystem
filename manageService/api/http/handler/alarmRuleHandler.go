package handler

import (
	"github.com/gin-gonic/gin"
	"manageService/internal/models"
	"net/http"
)

// Insert a rule into mysql then return Code and msg
func Insert(ctx *gin.Context) {
	alarmRule := models.AlarmRule{}
	var result = models.HttpResult{}
	// bind?
	err := ctx.ShouldBind(&alarmRule)
	if err == nil {
		// success
		num := alarmRule.Insert()
		result.CODE = http.StatusOK
		result.MSG = "Insert success!"
		result.DATA = gin.H{
			"num": num,
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func FindALL(ctx *gin.Context) {
	alarmRule := models.AlarmRule{}
	alarmRules := alarmRule.FindAll()
	result := models.HttpResult{
		CODE: http.StatusOK,
		MSG: "Find success!",
		DATA: alarmRules,
	}
	ctx.JSON(http.StatusOK, result)
}