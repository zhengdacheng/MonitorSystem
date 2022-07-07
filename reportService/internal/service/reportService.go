package service

import (
	"context"
	"log"
	"math/rand"
	"reportService/api/proto"
	"reportService/internal/configs"
	tsModel "reportService/internal/models"
)

type ReportService struct {
}

func (s *ReportService) Report(ctx context.Context, req *proto.ReportReq) (*proto.ReportResp, error) {
	// 接收agent的上报数据，以传入kafka为准
	dataUnit := tsModel.MonitorData{
		HostID:    req.HostID,
		CPURate:   req.CpuRate,
		MemRate:   req.MemRate,
		TimeStamp: req.TimeStamp,
	}
	//log.Println(dataUnit)
	// 写入kafka
	kafka := KafkaMessageQueue{
		Brokers:   configs.Broker,
		Topic:     configs.Topic,
		Partition: int32(rand.Intn(3)),
	}

	// new a producer
	producer, err := kafka.NewAsyncProducer()
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	// prepare msg
	msg := kafka.PrepareMessage(kafka.Topic, dataUnit)
	// send message by async producer
	err = kafka.StoreMessage(&producer, msg)
	if err != nil {
		log.Println("failure to send message")
	}
	// 暂未考虑过多的边缘条件，暂时以最顺利的情况做考虑，后几天慢慢加
	log.Println(req)
	resp := proto.ReportResp{
		Code:  200,
		Value: "report success.",
	}

	return &resp, nil
}
