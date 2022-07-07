package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manageService/internal/app"
	"manageService/internal/models"
	"manageService/internal/pkg"
	"net/http"
)

func HostInsert(ctx *gin.Context) {
	newHost := models.Host{}
	var result = models.HttpResult{}
	// bind?
	err := ctx.ShouldBind(&newHost)
	if err == nil {
		num := newHost.Insert()
		result.CODE = http.StatusOK
		result.MSG = "Insert success!"
		result.DATA = gin.H{
			"num": num,
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func FindALLHost(ctx *gin.Context) {
	newHost := models.Host{}
	newHosts := newHost.FindAll()
	result := models.HttpResult{
		CODE: http.StatusOK,
		MSG:  "Find success!",
		DATA: newHosts,
	}
	ctx.JSON(http.StatusOK, result)
}

func HostUpdate(ctx *gin.Context) {
	host := models.Host{}
	err := ctx.Bind(&host)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusBadGateway,
			"msg":  "参数错误",
			"data": err,
		})
		return
	}

	fmt.Println(host)
	updates := app.DB.Where("host_id=?", host.HostID).Updates(&host)
	if updates.RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "更新成功",
			"data": "ok",
		})
		return
	} else {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusOK,
			"msg":  "更新失败",
			"data": updates.Error,
		})
		return
	}
}

func GetUniqueHost(ctx *gin.Context) {
	hostIDs := pkg.GetHostOnlyInInfluxDB()
	result := models.HttpResult{
		CODE: http.StatusOK,
		MSG:  "Find success!",
		DATA: hostIDs,
	}
	ctx.JSON(http.StatusOK, result)
}
