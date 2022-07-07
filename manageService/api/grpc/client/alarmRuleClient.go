package main

import (
	"context"
	"log"
	"manageService/api/grpc/proto"

	"google.golang.org/grpc"

)

// Address 连接地址
const Address string = "172.19.0.11:8030"

var grpcClient proto.AlarmRuleServiceClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = proto.NewAlarmRuleServiceClient(conn)
	getAlarmRule()
}

// getAlarmRule 调用服务端Route方法
func getAlarmRule() {
	// 创建发送结构体
	req := proto.AlarmRuleQueryReq{}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.AlarmRuleQuery(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}