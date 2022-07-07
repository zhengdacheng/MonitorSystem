package configs
//
//import (
//	"encoding/json"
//	"io/ioutil"
//)
//
//var Cfg = &Config{}
//
//type KafkaConfig struct {
//	// kafka
//	//Broker = []string{"localhost:9092"}
//	//Topic = "MonitorData"
//	//Partition = int32(0)
//	Broker []string `json:"Broker"`
//	Topic string `json:"Topic"`
//
//}
//
//type MySQLConfig struct {
//	// mysql
//	//User     = "root"
//	//PassWord = "123456"
//	//Host     = "localhost"
//	//Port     = 3306
//	//DB       = "golang"
//	User     string `json:"User"`
//	PassWord string `json:"Password"`
//	Host     string `json:"Host"`
//	Port     int    `json:"Port"`
//	DB       string `json:"DB"`
//}
//
//type Config struct {
//	InfluxDBConfig
//	MySQLConfig
//}
//
//func InitConfigFromJson() {
//	data, err := ioutil.ReadFile("./config/config.json")
//	if err != nil {
//		panic("config initialization fail...")
//	}
//	err = json.Unmarshal(data, &Cfg)
//	if err != nil {
//		panic("config initialization fail...")
//	}
//}