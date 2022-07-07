package service

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	configs "persistentService/internal/config"
	"persistentService/internal/models"
	"time"
)

func WriteIntoInfluxDB() {
	// 1. consume from kafka
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(configs.Broker, config)
	if err != nil {
		return
	}

	// 2. write into influxdb
	client := influxdb2.NewClientWithOptions(configs.Server, configs.Token, influxdb2.DefaultOptions().SetBatchSize(20))

	// Get non-blocking write client
	writeAPI := client.WriteAPI(configs.Organization, configs.Bucket)

	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(configs.Topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		consumePartition, err := consumer.ConsumePartition(configs.Topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			return
		}
		for message := range consumePartition.Messages() {
			// unmarshal json to struct
			data := models.MonitorData{}
			err := json.Unmarshal(message.Value, &data)
			if err != nil {
				return
			}
			// create point
			p := influxdb2.NewPoint(
				configs.Measurement,
				map[string]string{
					"HostID": data.HostID,
				},
				map[string]interface{}{
					"CPU_Rate":  data.CPURate,
					"MEM_Rate": data.MemRate,
				},
				time.Unix(data.TimeStamp, 0))
			// write asynchronously
			writeAPI.WritePoint(p)
			fmt.Printf("[Consumer] data in Kafka's partitionid: %d; offset:%d, value: %s is saved\n", message.Partition, message.Offset, string(message.Value))
		}

	}

	// Flush writes
	writeAPI.Flush()
	client.Close()
	fmt.Println("finish")
}
