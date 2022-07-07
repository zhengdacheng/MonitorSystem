package main

import (
	"context"
	pb "dataAnalyzeService/api/grpc/proto"
	"log"

	"google.golang.org/grpc"
)

// Address 连接地址
const Address string = "172.19.0.5:8010"

var grpcClient pb.QueryClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewQueryClient(conn)
	getAlarmRule()
}

// getAlarmRule 调用服务端Route方法
func getAlarmRule() {
	// 创建发送结构体
	req := pb.QueryReq{
		MetricsType:       "CPU_Rate",
		Granularity:       "5m",
		AggregateFunction: "mean",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.AggregateQuery(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
