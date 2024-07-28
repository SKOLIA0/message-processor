package kafka

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

var consumer sarama.Consumer

// Initialize Kafka consumer
func InitConsumer() {
	var err error
	brokers := []string{os.Getenv("KAFKA_BROKER")}
	consumer, err = sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
}

// Consume messages from Kafka
func ConsumeMessages() {
	if consumer == nil {
		log.Fatal("Kafka consumer is not initialized")
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	topic := os.Getenv("KAFKA_TOPIC")
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error starting Kafka consumer for partition: %v", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Consumer started. Waiting for messages...")

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s", string(msg.Value))
			// Process the received message
		case <-signals:
			log.Println("Terminating consumer...")
			return
		}
	}
}
