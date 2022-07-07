package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	proto "manageService/api/grpc/proto"
	server "manageService/api/grpc/server"
	manageRouter "manageService/api/http/router"
	"manageService/internal/app"
	"manageService/internal/models"
	"net"
	"sync"
)

const (
	// Address 监听地址
	Address string = ":8030"
	// Network 网络通信协议
	Network string = "tcp"

	HTTPS string = ":8040"
)

func startHttpRouter() {
	router := manageRouter.SetupRouter()
	err := router.Run(HTTPS)
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
	proto.RegisterAlarmRuleServiceServer(grpcServer, &server.AlarmRuleServer{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func testFindLast() {
	alarmRule := models.AlarmRule{}
	code, rule :=alarmRule.FindLatest()
	if code == 200 {
		fmt.Println(rule.ContactEmail)
	}
}

func testInsert() {
	alarmRule := models.AlarmRule{
		CpuNoteworthyThreshold: 0.8,
		CpuSeriousThreshold:  0.9,
		CpuDeadlyThreshold: 0.95,

		MemNoteworthyThreshold: 0.8,
		MemSeriousThreshold:  0.9,
		MemDeadlyThreshold: 0.95,

		ContactEmail: "1305479162@qq.com",
	}
	alarmRule.Insert()
}

func testQuery(hostID string, newLocation string) {
	host := models.Host{}
	err := host.Update(hostID, newLocation)
	if err != nil {
		fmt.Println("更新失败")
	}
	fmt.Println("更新成功")
}

func main() {
	// Init config file
	/*data, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic("Read file error when config initialization...")
	}
	err = json.Unmarshal(data, &config.Cfg)
	if err != nil {
		panic("Config file unmarshal error when config initialization...")
	}*/
	// start mysql
	err := app.InitMysql()
	if err != nil {
		log.Println("start mysql error")
		return
	}
	err = app.DB.AutoMigrate(&models.AlarmRule{})
	if err != nil {
		panic(err)
	}
	err = app.DB.AutoMigrate(&models.Host{})
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go startHttpRouter()
	go startGRPCServer()
	wg.Wait()
	//testQuery("Kevin-D20", "777")
}
