package models

type TimeSeriesData struct {
	HostID       string    `json:"host_id"`
	MetricsType  string    `json:"metrics_type"`
	MetricsValue []float64 `json:"metrics_value"`
	TimeStamp    []int64   `json:"time_stamp"`
}
