package main

import (
    "fmt"

    "github.com/Shopify/sarama"
    "time"
)

//消息写入kafka
func main() {
    //初始化配置
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true
    //生产者
    client, err := sarama.NewSyncProducer([]string{"222.200.105.85:9092","222.200.105.85:9093","222.200.105.85:9094"}, config)
    if err != nil {
        fmt.Println("producer close,err:", err)
        return
    }

    defer client.Close()
    var n int=0

    for n<20{
        n++
        //创建消息
        msg := &sarama.ProducerMessage{}
        msg.Topic = "test"
        msg.Value = sarama.StringEncoder("this is a good test,hello chaoge!!")
        //发送消息
        pid, offset, err := client.SendMessage(msg)
        if err != nil {
            fmt.Println("send message failed,", err)
            return
        }
        fmt.Printf("pid:%v offset:%v\n,", pid, offset)
        time.Sleep(10 * time.Millisecond)

    }

}
