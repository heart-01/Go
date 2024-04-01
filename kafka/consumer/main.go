package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

func main() {
	group := os.Getenv("CONSUMER_GROUP")
	if group == "" {
		fmt.Println("CONSUMER_GROUP is required")
		os.Exit(1)
	}

	kConf := kafkaReader("message-to-say", group)
	reader := kafka.NewReader(kConf)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("error while receiving message: %s", err.Error())
			continue
		}

		if err != nil {
			fmt.Printf("error while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("[topic:partition:offset %v:%v:%v:] %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
	}
}

func kafkaReader(topic, group string) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:         []string{"127.0.0.1:9093"},
		GroupID:         group,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}
}
