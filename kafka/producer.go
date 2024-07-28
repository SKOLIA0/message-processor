package kafka

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"message-processor/db"
	"message-processor/models"
)

var producer sarama.SyncProducer

// Инициализация Kafka продюсера
func InitProducer() {
	var err error
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{os.Getenv("KAFKA_BROKER")}
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}
}

// Отправка сообщения в Kafka
func SendMessage(msg *models.Message) {
	message := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Value: sarama.StringEncoder(msg.Content),
	}

	_, _, err := producer.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message sent to Kafka successfully")

	// Пометка сообщения как обработанного
	db.MarkMessageProcessed(msg)
}
