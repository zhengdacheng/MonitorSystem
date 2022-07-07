package models

type AlarmRule struct {
	// cpu relative threshold
	CpuNoteworthyThreshold float32
	CpuSeriousThreshold float32
	CpuDeadlyThreshold float32

	// mem relative threshold
	MemNoteworthyThreshold float32
	MemSeriousThreshold float32
	MemDeadlyThreshold float32

	Granularity string
	AggregateFunction string
	ContactEmail string
}