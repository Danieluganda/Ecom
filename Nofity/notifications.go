package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Notification struct {
	ID           int
	UserID       int
	Message      string
	SentAt       time.Time
	IsDelivered  bool
	DeliveryType string
}

type NotificationService struct {
	db *sql.DB
}

func NewNotificationService(db *sql.DB) *NotificationService {
	return &NotificationService{
		db: db,
	}
}

func (ns *NotificationService) SendNotification(userID int, message, deliveryType string) (int64, error) {
	result, err := ns.db.Exec("INSERT INTO notifications (user_id, message, sent_at, is_delivered, delivery_type) VALUES (?, ?, ?, ?, ?)", userID, message, time.Now(), false, deliveryType)
	if err != nil {
		return 0, err
	}

	notificationID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return notificationID, nil
}

func (ns *NotificationService) GetNotification(notificationID int) (*Notification, error) {
	notification := &Notification{}
	err := ns.db.QueryRow("SELECT id, user_id, message, sent_at, is_delivered, delivery_type FROM notifications WHERE id = ?", notificationID).Scan(&notification.ID, &notification.UserID, &notification.Message, &notification.SentAt, &notification.IsDelivered, &notification.DeliveryType)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db_ecom")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Create notifications table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS notifications (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		message TEXT NOT NULL,
		sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		is_delivered BOOLEAN;
		delivery_type VARCHAR(255) NOT NULL
		)`
		if err != nil {
			log.Fatal(err)
		}