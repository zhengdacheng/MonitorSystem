package models

import (
	"gorm.io/gorm"
	"log"
	"manageService/internal/app"
)

type Host struct {
	gorm.Model
	// unique identifier
	HostID string
	Location string
}

func (host Host) Insert() int {
	create := app.DB.Create(&host)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (host Host) FindAll() []Host {
	var hosts []Host
	app.DB.Find(&hosts)
	return hosts
}

func (host Host) Update(hostID string, newLocation string) error {
	update := app.DB.Model(&host).Where("host_id = ?", hostID).Update("location", newLocation)
	log.Println("update function")
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func GetHostIDFromMysql() []string {
	var hostIDFromMysql []string
	host := Host{}
	result := host.FindAll()
	for _, h := range result{
		hostIDFromMysql = append(hostIDFromMysql, h.HostID)

	}
	return hostIDFromMysql
}