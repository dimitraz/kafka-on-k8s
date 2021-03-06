package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Set up the producer config
func newConfig() *kafka.ConfigMap {
	config := &kafka.ConfigMap{
		"bootstrap.servers": getEnv("SERVERS", "localhost"),
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
	// Configure the producer
	config := newConfig()

	// Create the producer
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatal("Could not create producer")
	}

	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			default:
				log.Printf("Unexpected event: %v\n", ev)
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := getEnv("TOPIC", "test-topic")
	for {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("Something Cool on %s", time.Now().Format("Mon Jan 2 15:04:05"))),
		}
		producer.Produce(msg, nil)
		time.Sleep(10 * time.Second)
	}
}
