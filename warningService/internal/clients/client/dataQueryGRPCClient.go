package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	dataProto "warningService/internal/clients/protos"
	"warningService/internal/models"
)

// AddressForData 连接地址
const AddressForData string = "172.19.0.5:8010"

var grpcClient dataProto.QueryClient

// GetAggregateValue ss
func GetAggregateValue(Granularity string, AggregateFunction string) (models.AggregateValue, int32) {
	// 连接服务器
	conn, err := grpc.Dial(AddressForData, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	// 建立返回结构体
	targetAggregateValues := models.AggregateValue{}

	// 建立gRPC连接
	grpcClient = dataProto.NewQueryClient(conn)
	// 创建发送结构体
	req := dataProto.QueryReq{
		MetricsType:       "CPU_Rate",
		Granularity:       Granularity,
		AggregateFunction: AggregateFunction,
	}
	resp, err := grpcClient.AggregateQuery(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 将cpu的聚合数据赋值
	targetAggregateValues.HostIDs = resp.HostId
	targetAggregateValues.CpuRateAggregateValues = resp.AggregateValue
	req.MetricsType = "MEM_Rate"
	resp, err = grpcClient.AggregateQuery(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	targetAggregateValues.MemRateAggregateValues = resp.AggregateValue
	code := resp.Code
	return targetAggregateValues, code
}
