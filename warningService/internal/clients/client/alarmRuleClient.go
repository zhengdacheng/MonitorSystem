package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	proto "warningService/internal/clients/protos"
	"warningService/internal/models"
)

// Address 连接地址
const Address string = "172.19.0.11:8030"

var AlarmRuleGrpcClient proto.AlarmRuleServiceClient

// GetAlarmRule 调用服务端Route方法
func GetAlarmRule() (models.AlarmRule, int32) {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	AlarmRuleGrpcClient = proto.NewAlarmRuleServiceClient(conn)
	// 创建发送结构体
	req := proto.AlarmRuleQueryReq{}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	resp, err := AlarmRuleGrpcClient.AlarmRuleQuery(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印 & 返回

	if resp.Code == 200 {
		log.Printf("Latest alarm: %v", resp)
	} else if resp.Code == 400 {
		log.Println("There is not a rule is set.")
		log.Println(resp)
	}

	alarmRule := models.AlarmRule{
		CpuNoteworthyThreshold: resp.CpuNoteworthyThreshold,
		CpuSeriousThreshold: resp.CpuSeriousThreshold,
		CpuDeadlyThreshold: resp.CpuDeadlyThreshold,
		MemNoteworthyThreshold: resp.MemNoteworthyThreshold,
		MemSeriousThreshold: resp.MemSeriousThreshold,
		MemDeadlyThreshold: resp.MemDeadlyThreshold,
		Granularity: resp.Granularity,
		AggregateFunction: resp.AggregateFunction,
		ContactEmail: resp.ContactEmail,
	}
	return alarmRule, resp.Code
}