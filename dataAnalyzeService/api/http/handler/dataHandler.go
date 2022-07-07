package handler

import (
	configs "dataAnalyzeService/internal/config"
	"dataAnalyzeService/internal/models"
	"dataAnalyzeService/internal/pkg"
	"dataAnalyzeService/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DataQuery(ctx *gin.Context) {
	// 加载查询相关参数
	queryParam := &models.QueryParam{
		Bucket:        configs.Bucket,
		StartFrom:     ctx.Query("StartFrom"),
		Measurement:   configs.Measurement,
		TagsKV:        []pkg.TagsKV{pkg.TagsKV{TagKey: "HostID", TagValue: ctx.Query("HostID")}},
		Fields:        ctx.Query("MetricsType"),
		Duration:      ctx.Query("Granularity"),
		AggregateFunc: ctx.Query("AggregateFunc"),
	}
	// 执行查询
	record := service.QueryRecords(queryParam)

	// 返回json
	ctx.JSON(http.StatusOK, record)
}

func DataQueryByLocation(ctx *gin.Context) {
	// 先找到Location对应的几台Host
	location := ctx.Query("Location")
	host := models.Host{}
	relativeHostIDs := host.FindAllHostIdByLocation(location)
	log.Printf("Hosts in %s: %v", location, relativeHostIDs)
	// pack up
	paramTagKVs := make([]pkg.TagsKV, len(relativeHostIDs))
	for i := 0; i < len(relativeHostIDs); i++ {
		paramTagKVs[i] = pkg.TagsKV{
			TagKey: "HostID",
			TagValue: relativeHostIDs[i],
		}
	}
	// query param
	queryParam := &models.QueryParam{
		Bucket:        configs.Bucket,
		StartFrom:     ctx.Query("StartFrom"),
		Measurement:   configs.Measurement,
		TagsKV:        paramTagKVs,
		Fields:        ctx.Query("MetricsType"),
		Duration:      ctx.Query("Granularity"),
		AggregateFunc: ctx.Query("AggregateFunc"),
	}

	// 执行查询
	record := service.QueryRecords(queryParam)

	// 返回json
	ctx.JSON(http.StatusOK, record)
}

