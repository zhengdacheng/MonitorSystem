package main

import (
	"dataAnalyzeService/api/grpc/proto"
	"dataAnalyzeService/api/grpc/server"
	router2 "dataAnalyzeService/api/http/router"
	"dataAnalyzeService/internal/app"
	"dataAnalyzeService/internal/models"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

const (
	// Address 监听地址
	Address string = "172.19.0.5:8010"
	// Network 网络通信协议
	Network string = "tcp"

	HTTPPort string = "172.19.0.5:8020"
)

func startHttpRouter() {
	router := router2.SetupRouter()
	err := router.Run(HTTPPort)
	if err != nil {
		return
	}
}

func startGRPCServer() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	proto.RegisterQueryServer(grpcServer, &server.QueryServer{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func main() {
	// Init config file
	/*data, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic("Read file error when config initialization...")
	}
	err = json.Unmarshal(data, &configs.Cfg)
	if err != nil {
		panic("Config file unmarshal error when config initialization...")
	}*/

	// Init MySQL
	err := app.InitMysql()
	if err != nil {
		panic("Start MySQL error")
	}
	err = app.DB.AutoMigrate(&models.Host{})
	if err != nil {
		panic(err)
	}

	// Init InfluxDB
	//err = app.InitInfluxDB()
	//if err != nil {
	//	return
	//}

	// Start two goroutine for gRPC and gin router
	wg := sync.WaitGroup{}
	wg.Add(2)
	go startHttpRouter()
	go startGRPCServer()
	wg.Wait()
}