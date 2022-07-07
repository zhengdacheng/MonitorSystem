package models

import "dataAnalyzeService/internal/pkg"

type QueryParam struct {
	Bucket        string
	StartFrom     string
	Measurement   string
	TagsKV        []pkg.TagsKV
	Fields        string
	Duration      string
	AggregateFunc string
}
