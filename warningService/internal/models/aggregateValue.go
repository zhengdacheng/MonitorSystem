package models

type AggregateValue struct {
	HostIDs                []string
	CpuRateAggregateValues []float32
	MemRateAggregateValues []float32
}
