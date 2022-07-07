package etcd

import (
    "context"
    "github.com/coreos/etcd/mvcc/mvccpb"
    "go.etcd.io/etcd/clientv3"
    "google.golang.org/grpc/resolver"
    "log"
    "sync"
    "time"
)

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
    Cli        *clientv3.Client //etcd client
    CConn         resolver.ClientConn
    ServerList map[string]resolver.Address //服务列表
    Lock       sync.Mutex
}

func NewServiceDiscovery(endpoints []string) (resolver.Builder, error) {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: time.Second * 5, //超时时间
    })
    if err!=nil {
        log.Fatalf("client.New err:%v\n",err)
        return nil,err
    }
    ser:=&ServiceDiscovery{
        Cli: client,
    }
    return ser,nil
}

//Build 为给定目标创建一个新的`resolver`，当调用`grpc.Dial()`时执行
func (s *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
    //ClientConn进行存储
    s.CConn = cc
    //创建map
    s.ServerList = make(map[string]resolver.Address)
    //"/monitor_grpcLB/report_grpc/localhost:8000"
    prefix := "/" + target.Scheme + "/" + target.Endpoint + "/"
    log.Println("prefix:"+prefix)
    //根据前缀获取现有的key
    resp, err := s.Cli.Get(context.Background(), prefix, clientv3.WithPrefix())
    if err != nil {
        log.Fatalf("s.Cli.Get err:%v\n",err)
        return nil, err
    }
    //遍历，保存注册列表
    for _, ev := range resp.Kvs {
        s.SetServiceList(string(ev.Key), string(ev.Value))
    }
    log.Println(s.ServerList)
    s.CConn.NewAddress(s.getServices())
    //监视前缀，修改变更的server
    go s.watcher(prefix)
    return s, nil
}

// ResolveNow 监视目标更新
func (s *ServiceDiscovery) ResolveNow(rn resolver.ResolveNowOptions) {
    log.Println("ResolveNow")
}

//Scheme return schema
func (s *ServiceDiscovery) Scheme() string {
    return Schema
}

//Close 关闭
func (s *ServiceDiscovery) Close() {
    log.Println("Close")
    s.Cli.Close()
}

//watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
    //监听租约绑定的key
    rch := s.Cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
    log.Printf("watching prefix:%s now...", prefix)
    for wresp := range rch {
        for _, ev := range wresp.Events {
            switch ev.Type {
            case mvccpb.PUT: //新增或修改
                s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
            case mvccpb.DELETE: //删除
                s.DelServiceList(string(ev.Kv.Key))
            }
        }
    }
}

//SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
    s.Lock.Lock()
    defer s.Lock.Unlock()
    s.ServerList[key] = resolver.Address{Addr: val}
    s.CConn.NewAddress(s.getServices())
    log.Println("put key :", key, "val:", val)
}

//DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
    s.Lock.Lock()
    defer s.Lock.Unlock()
    delete(s.ServerList, key)
    s.CConn.NewAddress(s.getServices())
    log.Println("del key:", key)
}

//GetServices 获取服务地址
func (s *ServiceDiscovery) getServices() []resolver.Address {
    addrs := make([]resolver.Address, 0, len(s.ServerList))
    for _, v := range s.ServerList {
        addrs = append(addrs, v)
    }
    return addrs
}