package service

import (
	"context"
	configs "dataAnalyzeService/internal/config"
	"dataAnalyzeService/internal/models"
	"dataAnalyzeService/internal/pkg"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

func QueryRecords(param *models.QueryParam) []models.TimeSeriesData {
	client := influxdb2.NewClient(configs.Server, configs.Token)
	defer client.Close()

	// Get query client
	queryAPI := client.QueryAPI(configs.Organization)

	// query
	var length = len(param.TagsKV)
	record := make([]models.TimeSeriesData, length)
	for i := 0; i < length; i++ {
		// query Host's record one by one: InfluxDB Go Client for bath searching is not finished yet.
		record[i].HostID = param.TagsKV[i].TagValue
		record[i].MetricsType = param.Fields

		query := pkg.NewQuery()
		queryScript := query.From(param.Bucket).TimeRange(param.StartFrom).Measurement(
			param.Measurement).Tags([]pkg.TagsKV{param.TagsKV[i]}).Fields(
			param.Fields).Window(param.Duration).AggregateFunc(param.AggregateFunc).Tail().Done()

		log.Println(queryScript)
		result, err := queryAPI.Query(context.Background(), queryScript)
		if err == nil {
			for result.Next() {
				record[i].MetricsValue = append(record[i].MetricsValue, result.Record().Value().(float64))
				record[i].TimeStamp = append(record[i].TimeStamp, result.Record().Time().Unix())
			}
			// check for an error
			if result.Err() != nil {
				fmt.Printf("query parsing error: %s\n", result.Err().Error())
			}
		}
	}

	return record
}
