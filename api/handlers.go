package api

import (
	"encoding/json"
	"log"
	"net/http"

	"message-processor/db"
	"message-processor/kafka"
	"message-processor/models"
)

// Обработчик для создания нового сообщения
func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println("Error decoding message:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Message received:", msg.Content)

	// Сохранение сообщения в базе данных
	db.SaveMessage(&msg)
	// Отправка сообщения в Kafka
	kafka.SendMessage(&msg)

	log.Println("Message created successfully")
	w.WriteHeader(http.StatusCreated)
}

// Обработчик для получения статистики по обработанным сообщениям
func GetMessageStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := db.GetMessageStats()
	json.NewEncoder(w).Encode(stats)
}
