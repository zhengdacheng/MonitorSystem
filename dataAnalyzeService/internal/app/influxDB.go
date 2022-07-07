package app

import (
	config "dataAnalyzeService/internal/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var InfluxClient influxdb2.Client
var QueryAPI api.QueryAPI

func InitInfluxDB() (err error) {
	InfluxClient = influxdb2.NewClient(config.Cfg.Server, config.Cfg.Token)
	// Get query client
	QueryAPI = InfluxClient.QueryAPI(config.Cfg.Organization)
	return nil
}
