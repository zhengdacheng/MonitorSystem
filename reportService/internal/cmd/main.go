package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"reportService/api/proto"
	"reportService/internal/etcd"
	"reportService/internal/service"
)

var (
	// Address 监听地址
	Address *string = flag.String("addr","172.19.0.6:8080","the server addr")
	// Network 网络通信协议
	Network string = "tcp"
	// SerName 服务名称
	SerName string = "report_grpc"
)

// EtcdEndpoints etcd地址
var EtcdEndpoints = []string{"172.19.0.20:2379"}

func main() {
	//解析flag
	flag.Parse()
	//初始化日志，标记位置
	log.SetFlags(log.Ldate|log.Llongfile)
	// 监听本地端口
	listener, err := net.Listen(Network, *Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(*Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	proto.RegisterReportServer(grpcServer, &service.ReportService{})
	//把服务注册到etcd
	ser, err := etcd.NewServiceRegister(EtcdEndpoints, SerName, *Address, 5)
	if err != nil {
		log.Fatalf("register service err: %v", err)
	}
	defer func(ser *etcd.ServiceRegister) {
		err := ser.Close()
		if err != nil {
			log.Fatalf("ServiceRegister close err:%v\n",err)
		}
	}(ser)
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
