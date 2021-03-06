package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Set up the consumer config
func newConfig() *kafka.ConfigMap {
	config := &kafka.ConfigMap{
		"bootstrap.servers": getEnv("SERVERS", "localhost"),
		"group.id":          getEnv("GROUP_ID", "consumer-group"),
		"auto.offset.reset": "latest",
	}
	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Configure the consumer
	config := newConfig()

	// Create the consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatal("Could not create consumer")
	}

	// Subscribe to topic and consume messages
	topic := getEnv("TOPIC", "test-topic")
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
}
