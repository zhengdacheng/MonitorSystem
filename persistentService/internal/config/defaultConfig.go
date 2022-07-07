package configs

var (
	// kafka
	Broker = []string{"172.19.0.15:9092","172.19.0.16:9093","172.19.0.17:9094"}
	Topic = "MonitorData"
	// influxdb
	Server       = "http://172.19.0.18:8086"
	Token        = "8x9SIwvmkcUbbXTtT3JEKoz7iRHuLARWpiBC8b92DxKZjy3OwbpJlTcE_SeuOpaHFAb4uCjcglpJvhVIlwxf2A=="
	Organization = "gdut_4"
	Bucket       = "Monitor"
	Measurement  = "host"
)
