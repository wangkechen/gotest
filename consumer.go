package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"time"
)


var (
	kafkaConsumer *cluster.Consumer
	//kafkaBrokers = []string{"127.0.0.1:9092"}
	kafkaBrokers = []string{"10.155.0.135:30947"}
	kafkaTopic   = "harix_patrol_exception"
	groupId = ""
	//groupId = "csdn_test_1"
)

func init11() {
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	kafkaConsumer, err = cluster.NewConsumer(kafkaBrokers, groupId, []string{kafkaTopic}, config)
	if err != nil {
		panic(err.Error())
	}
	if kafkaConsumer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v, topic: %v, group: %v}", kafkaBrokers, kafkaTopic, groupId))
	}
	fmt.Printf("kafka init success, consumer -> %v, topic -> %v, ", kafkaConsumer, kafkaTopic)
}

func main111() {
	for {
		select {
		case msg, ok := <-kafkaConsumer.Messages():
			if ok {
				fmt.Printf("kafka msg: %s \n", msg.Value)
				kafkaConsumer.MarkOffset(msg, "")
			} else {
				fmt.Printf("kafka 监听服务失败")
			}
		case err, ok := <-kafkaConsumer.Errors():
			if ok {
				fmt.Printf("consumer error: %v" , err)
			}
		case ntf, ok := <-kafkaConsumer.Notifications():
			if ok {
				fmt.Printf("consumer notification: %v" , ntf)
			}
		}
	}
}