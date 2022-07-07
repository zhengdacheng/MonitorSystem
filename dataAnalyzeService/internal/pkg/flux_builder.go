package pkg

import (
	"bytes"
	"fmt"
	"strings"
)


type FluxQuery interface {
	From(bucket string) FluxQuery
	TimeRange(duration string) FluxQuery
	Measurement(measurement string) FluxQuery
	Tags(tagsKV []TagsKV) FluxQuery
	Fields(fields ...string) FluxQuery
	Window(duration string) FluxQuery
	AggregateFunc(aggregateFunc string) FluxQuery
	Tail() FluxQuery
	Done() string
	AggregateDone() string
}

type TagsKV struct {
	TagKey   string
	TagValue string
}

type Query struct {
	bucket        string
	startAt       string
	measurement   string
	tags          []TagsKV
	fields        []string
	duration      string
	aggregateFunc string
}

func (q *Query) From(bucket string) FluxQuery {
	q.bucket = bucket
	return q
}

func (q *Query) TimeRange(startAt string) FluxQuery {
	q.startAt = startAt
	return q
}

func (q *Query) Measurement(measurement string) FluxQuery {
	q.measurement = measurement
	return q
}

func (q *Query) Tags(tagKV []TagsKV) FluxQuery {
	q.tags = tagKV
	return q
}

func (q *Query) Fields(field ...string) FluxQuery {
	q.fields = append(q.fields, field...)
	return q
}

func (q *Query) Window(duration string) FluxQuery {
	q.duration = duration
	return q
}

func (q *Query) AggregateFunc(aggregateFunc string) FluxQuery {
	q.aggregateFunc = aggregateFunc
	return q
}

func (q *Query) Tail() FluxQuery {
	return q
}

func NewQuery() FluxQuery {
	return &Query{}
}

func (q *Query) BuildBucket() string {
	if q.bucket == "" {
		panic("There must be something in Bucket!")
	}
	return fmt.Sprintf(`from(bucket: "%s")`, q.bucket)
}

func (q *Query) BuildTimeRange() string {
	return fmt.Sprintf(` |> range(start: -%s)`, q.startAt)
}

func (q *Query) BuildMeasurement() string {
	return fmt.Sprintf(` |> filter(fn: (r) => r["_measurement"] == "%s")`, q.measurement)
}

func (q *Query) BuildTags() string {
	var tagsBuffer bytes.Buffer
	beginner := " |> filter(fn: (r) => "
	tagsBuffer.WriteString(beginner)
	for _, tag := range q.tags{
		tagsBuffer.WriteString(fmt.Sprintf(`r["%s"] == "%s"`, tag.TagKey, tag.TagValue))
	}
	tagsBuffer.WriteString(")")
	return tagsBuffer.String()
}

func (q *Query) BuildFields() string {
	return fmt.Sprintf(` |> filter(fn: (r) => r["_field"] == "%s")`, q.fields[0])
}

func (q *Query) BuildWindow() string {
	return fmt.Sprintf(` |> aggregateWindow(every: %s, fn: %s, createEmpty: false)`, q.duration, q.aggregateFunc)
}

func (q *Query) GetTail() string {
	return fmt.Sprintf(` |> yield(name: "%s")`, q.aggregateFunc)
}

func (q *Query) Done() string {
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString(q.BuildBucket())
	queryBuffer.WriteString(q.BuildTimeRange())
	if len(q.tags) != 0 {
		queryBuffer.WriteString(q.BuildTags())
	}
	queryBuffer.WriteString(q.BuildFields())
	queryBuffer.WriteString(q.BuildWindow())
	queryBuffer.WriteString(q.GetTail())

	return strings.TrimSpace(queryBuffer.String())
}

func (q *Query) AggregateDone() string {
	var queryBuffer bytes.Buffer

	queryBuffer.WriteString(q.BuildBucket())
	queryBuffer.WriteString(q.BuildTimeRange())
	if len(q.tags) != 0 {
		queryBuffer.WriteString(q.BuildTags())
	}
	queryBuffer.WriteString(q.BuildFields())
	queryBuffer.WriteString(fmt.Sprintf(` |> %s()`, string(q.aggregateFunc)))

	return strings.TrimSpace(queryBuffer.String())
}