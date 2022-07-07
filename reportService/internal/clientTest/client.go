package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/resolver"
	"log"
	"math/rand"
	pb "reportService/api/proto"
	"reportService/internal/etcd"
	"time"

	"google.golang.org/grpc"
)


var (
	// EtcdEndpoints etcd地址
	EtcdEndpoints = []string{"172.19.0.20:2379"}
	// SerName 服务名称
	SerName    = "report_grpc"
	grpcClient pb.ReportClient
)

func main() {
	/*// 连接服务器
	conn, err := grpc.Dial(AddressRpc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()*/
	//初始化日志，标记位置
	log.SetFlags(log.Ldate|log.Llongfile)
	//新建服务发现
	discovery, err := etcd.NewServiceDiscovery(EtcdEndpoints)
	if err!=nil {
		panic(err)
	}
	//进行服务发现
	resolver.Register(discovery)
	// 连接服务器
	//"/monitor_grpcLB/report_grpc/localhost:8000"
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", discovery.Scheme(), SerName),
		grpc.WithBalancerName("round_robin"),//轮询
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewReportClient(conn)
	for i := 0; i < 50; i++ {
		do()
		time.Sleep(time.Second * 1)
	}

}

// do 调用服务端Route方法
func do() {
	// 创建发送结构体
	req := pb.ReportReq{
		HostID:    "kevin_D20",
		CpuRate:   rand.Float32(),
		MemRate:   rand.Float32(),
		TimeStamp: time.Now().Unix(),
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Report(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}


