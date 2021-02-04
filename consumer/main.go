package consumer

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitConsumer() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka Consumer has been started")

	for {
		message, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(message.Value))
		}
	}
}
