package etcd

import (
    "context"
    "go.etcd.io/etcd/clientv3"
    "log"
    "time"
)

const Schema = "monitor_grpcLB"

// ServiceRegister 创建租约注册服务
type ServiceRegister struct {
    Cli *clientv3.Client //etcd client
    LeaseId clientv3.LeaseID //etcd 租约ID
    KeepAliveChan <- chan *clientv3.LeaseKeepAliveResponse //健康检查chan
    Key string //etcd key
    Value string //etcd value
}

// NewServiceRegister 新建服务注册，把地址注册到etcd中
func NewServiceRegister(endpoints []string, serName, addr string, leaseTime int64) (*ServiceRegister, error) {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: time.Second * 5, //设置超时时间5s
    })
    if err!=nil {
        log.Fatalf("The process of clientv3.New:%v\n",err)
    }
    //创建注册服务
    ser:=&ServiceRegister{
        Cli: client,
        //"/monitor_grpcLB/report_grpc/localhost:8000"
        Key:"/"+Schema+"/"+serName+"/"+addr,
        Value: addr,
    }
    //设置租约，将租约和key、value进行绑定
    if err=ser.putKeyWithLease(leaseTime);err!= nil {
      return nil,err
    }
    return ser,nil
}

//设置租约
func (s * ServiceRegister) putKeyWithLease(leaseTime int64) error {
    //设置租约、租约时间
    grantResponse, err := s.Cli.Grant(context.Background(), leaseTime)
    if err!=nil {
        log.Fatalf("grant fail:%v\n",err)
        return err
    }
    //将key、value与租约绑定
    _, err = s.Cli.Put(context.Background(), s.Key, s.Value, clientv3.WithLease(grantResponse.ID))
    if err!=nil {
        log.Fatalf("the fail of putting key with grant:%v\n",err)
        return err
    }
    //保持租约存活
    alive, err := s.Cli.KeepAlive(context.Background(), grantResponse.ID)
    if err!=nil {
        log.Fatalf("the fail of keepAlive with grant:%v\n",err)
        return err
    }
    //把lease Id 和存活凭证存到struct中
    s.KeepAliveChan=alive
    s.LeaseId=grantResponse.ID
    log.Printf("Put key:%s  val:%s  success!", s.Key, s.Value)
    return nil
}

//ListenLeaseRespChan 监听 续租情况
func (s *ServiceRegister) ListenLeaseRespChan() {
    for leaseKeepResp := range s.KeepAliveChan {
        log.Println("the success of grant:", leaseKeepResp)
    }
    log.Println("the close of grant")
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
    //撤销租约
    if _, err := s.Cli.Revoke(context.Background(), s.LeaseId); err != nil {
        return err
    }
    log.Println("revoke grant")
    return s.Cli.Close()
}