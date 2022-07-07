package app

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"manageService/internal/config"
)

func GetHostIDFromInfluxDB() []string {
	var hostIDs []string
	client := influxdb2.NewClient(config.Cfg.Server, config.Cfg.Token)
	defer client.Close()
	// Get query client
	queryAPI := client.QueryAPI(config.Cfg.Organization)
	query, err := queryAPI.Query(context.Background(), "from(bucket: \"Monitor\")\n  |> range(start: -5m)\n  |> filter(fn: (r) => r[\"_measurement\"] == \"host\")\n  |> filter(fn: (r) => r[\"HostID\"] == \"Kevin-D20\")\n  |> filter(fn: (r) => r[\"_field\"] == \"CPU_Rate\")\n  |> mean()")
	if err == nil {
		for query.Next() {
			hostIDs = append(hostIDs, query.Record().ValueByKey("HostID").(string))
		}
	} else {
		panic(err)
	}
	return hostIDs
}
