### 上报服务模块
## Note:
* 此模块为核心模块之一，负责接收agent上报的数据并存入kafka中；
* 此处需要增加冗余，多个服务节点注册到etcd后被agent客户端发现服务，并将agent客户端的服务请求负载均衡到各节点上；
