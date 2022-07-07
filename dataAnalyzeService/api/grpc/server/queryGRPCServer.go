package server

import (
	"context"
	pb "dataAnalyzeService/api/grpc/proto"
	"dataAnalyzeService/internal/app"
	configs "dataAnalyzeService/internal/config"
	"dataAnalyzeService/internal/pkg"
	"fmt"
	"log"
)

type QueryServer struct {
}

func (s *QueryServer) AggregateQuery(ctx context.Context, req *pb.QueryReq) (*pb.QueryResp, error) {
	// Get query script
	query := pkg.NewQuery()
	queryScript := query.From(configs.Bucket).TimeRange(req.Granularity).Measurement(
		configs.Measurement).Fields(req.MetricsType).Window(req.Granularity).AggregateFunc(
		req.AggregateFunction).Tail().AggregateDone()
	log.Println(queryScript)
	//queryScript := "from(bucket: \"Monitor\")\n  |> range(start: -5m)\n  |> filter(fn: (r) => r[\"_measurement\"] == \"host\")\n  |> filter(fn: (r) => r[\"HostID\"] == \"Kevin-D20\")\n  |> filter(fn: (r) => r[\"_field\"] == \"CPU_Rate\")\n  |> mean()"
	// Execute query
	result, err := app.QueryAPI.Query(context.Background(), queryScript)
	log.Println("start query")
	// Result
	queryResp := pb.QueryResp{}
	var hostIDs []string
	var aggregateValues []float64
	if err == nil {
		for result.Next() {
			hostIDs = append(hostIDs, result.Record().ValueByKey("HostID").(string))
			aggregateValues = append(aggregateValues, result.Record().Value().(float64))
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
			queryResp.Code = 400
			queryResp.AggregateValue = []float32{}
			return &queryResp, result.Err()
		}
	} else {
		queryResp.Code = 400
		queryResp.AggregateValue = []float32{}
		return &queryResp, err
	}

	queryResp.Code = 200
	for i := 0; i < len(aggregateValues); i++ {
		queryResp.HostId = append(queryResp.HostId, hostIDs[i])
		queryResp.AggregateValue = append(queryResp.AggregateValue, float32(aggregateValues[i]))
	}
	return &queryResp, nil
}
