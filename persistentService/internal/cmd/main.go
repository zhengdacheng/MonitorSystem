package main

import "persistentService/internal/service"

func main() {
	for  {
		service.WriteIntoInfluxDB()
	}
}