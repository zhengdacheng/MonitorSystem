package pkg

import (
	"agentService/internal/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"time"
)

type MemStatus struct {
	All            uint64  //总的内存容量
	Available      uint64  //内存已使用部分
	Used           uint64  //内存空闲
	MemUsedPercent float64 //内存利用率
}

type CPUStatus struct {
	CPUUsedPercent float64 //cpu利用率
}

func MemStat() MemStatus {
	MemStatus := MemStatus{}
	memory, err := mem.VirtualMemory() //获取内存信息
	if err != nil {
		log.Fatalf("The error when get memory percentage:%v\n", err)
	}
	MemStatus.All = memory.Total           //内存总容量
	MemStatus.Available = memory.Available //内存可用容量
	MemStatus.Used = memory.Used           //内存已用容量
	MemStatus.MemUsedPercent = memory.UsedPercent

	return MemStatus
}

func CPUStat() CPUStatus {
	cpuStatus := CPUStatus{}
	percent, err := cpu.Percent(time.Second, false) //按秒获取cpu利用率
	if err != nil {
		log.Fatalf("The error when get cpu percentage:%v\n", err)
	}
	cpuStatus.CPUUsedPercent = percent[0]
	
	return cpuStatus
}

func GetMetrics() *models.MonitorData {

	// cpu
	cpuRate := CPUStat().CPUUsedPercent
	// mem
	memRate := MemStat().MemUsedPercent
	// hostname
	hostname, err := os.Hostname()
	if err != nil {
		return nil
	}
	// timestamp
	timeStamp := time.Now().Unix()
	var dataUnit = models.MonitorData{
		HostID:    hostname,
		CPURate:   cpuRate,
		MemRate:   memRate,
		TimeStamp: timeStamp,
	}
	return &dataUnit
}