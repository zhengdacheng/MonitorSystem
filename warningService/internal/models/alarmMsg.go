package models

type AlarmMsg struct {
	CpuInDangerHostIDs []string
	CpuInDangerValues  []float32
	MemInDangerHostIDs []string
	MemInDangerValues  []float32
}
