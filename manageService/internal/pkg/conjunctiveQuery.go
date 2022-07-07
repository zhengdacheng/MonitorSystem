package pkg

import (
	"manageService/internal/app"
	"manageService/internal/models"
)

func GetHostOnlyInInfluxDB() []string {
	hostIDs := make([]string, 1)
	err := app.DB.AutoMigrate(&models.Host{})
	if err != nil {
		panic(err)
	}
	// from mysql
	hostIDFromMysql := models.GetHostIDFromMysql()
	// from InfluxDB
	hostIDFromInfluxDB := app.GetHostIDFromInfluxDB()
	// records from mysql fill in a set
	hostIDSet := make(map[string]bool)
	for _, hid := range hostIDFromMysql {
		//fmt.Println("mysql:" + hid)
		hostIDSet[hid] = true
	}
	for _, hidInfluxdb := range hostIDFromInfluxDB {
		//fmt.Println("influxdb:" + hidInfluxdb)
		if !hostIDSet[hidInfluxdb] {
			//fmt.Println(hidInfluxdb)
			hostIDs = append(hostIDs, hidInfluxdb)
		}
	}
	return hostIDs[1:]
}
