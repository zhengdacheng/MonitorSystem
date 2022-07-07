package service

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	models2 "reportService/internal/models"
)

type MessageQueue interface {
	StoreMessage(data models2.MonitorData)
}

type KafkaMessageQueue struct {
	// Kafka 相关配置
	Brokers   []string
	Topic     string
	Partition int32
	//Partition .Producer.Partitioner

}

func (kafka *KafkaMessageQueue) NewAsyncProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	// config.Producer.Return.Successes = true
	client, err := sarama.NewClient(kafka.Brokers, config)
	if err != nil {
		fmt.Println("client initialize fail...")
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	return producer, err
}

func (kafka *KafkaMessageQueue) PrepareMessage(topic string, message models2.MonitorData) *sarama.ProducerMessage {
	// message -> json
	byteMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println("encode struct to json fail...")
	}
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: 0,
		Value:     sarama.ByteEncoder(byteMessage),
	}

	return msg
}

func (kafka *KafkaMessageQueue) StoreMessage(producer *sarama.AsyncProducer, message *sarama.ProducerMessage) error {
	(*producer).Input() <- message
	// wait response
	select {
	case err := <-(*producer).Errors():
		log.Println("Produced message failure: ", err)
		return err
	default:
		log.Println("success")
		return nil
	}
}
