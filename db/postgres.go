package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"message-processor/models"
)

var db *sql.DB

// Инициализация подключения к базе данных PostgreSQL
func InitDB() {
	var err error
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_DBNAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы сообщений, если она не существует
	createTableQuery := `CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		processed BOOLEAN DEFAULT FALSE
	)`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized and table created if not exists")
}

// Сохранение нового сообщения в базе данных
func SaveMessage(msg *models.Message) {
	log.Printf("Saving message: %s", msg.Content)
	_, err := db.Exec("INSERT INTO messages (content) VALUES ($1)", msg.Content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message saved successfully")
}

// Пометка сообщения как обработанного
func MarkMessageProcessed(msg *models.Message) {
	log.Printf("Marking message as processed: %s", msg.Content)
	_, err := db.Exec("UPDATE messages SET processed = TRUE WHERE content = $1", msg.Content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Message marked as processed")
}

// Получение статистики по обработанным сообщениям
func GetMessageStats() map[string]int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = TRUE").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return map[string]int{"processed_messages": count}
}
