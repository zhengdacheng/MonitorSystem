package config

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg = &Config{}

type InfluxDBConfig struct {
	// influxdb
	//Server       = "http://localhost:8086"
	//Token        = "8x9SIwvmkcUbbXTtT3JEKoz7iRHuLARWpiBC8b92DxKZjy3OwbpJlTcE_SeuOpaHFAb4uCjcglpJvhVIlwxf2A=="
	//Organization = "gdut_4"
	//Bucket       = "Monitor"
	//Measurement  = "host"
	Server string `json:"Server"`
	Token  string `json:"Token"`
	Organization string `json:"Organization"`
	Bucket       string `json:"Bucket"`
	Measurement  string `json:"Measurement"`
}

type MySQLConfig struct {
	// mysql
	//User     = "root"
	//PassWord = "123456"
	//Host     = "localhost"
	//Port     = 3306
	//DB       = "golang"
	User     string `json:"User"`
	PassWord string `json:"Password"`
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	DB       string `json:"DB"`
}

type Config struct {
	InfluxDBConfig
	MySQLConfig
}

func InitConfigFromJson() {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic("config initialization fail...")
	}
	err = json.Unmarshal(data, &Cfg)
	if err != nil {
		panic("config initialization fail...")
	}
}