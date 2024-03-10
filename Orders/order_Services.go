package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	ID        int
	UserID    int
	Amount    float64
	CreatedAt time.Time
}

type OrderService struct {
	db *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (os *OrderService) CreateOrder(userID int, amount float64) (int64, error) {
	result, err := os.db.Exec("INSERT INTO orders (user_id, amount, created_at) VALUES (?, ?, ?)", userID, amount, time.Now())
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (os *OrderService) GetOrder(orderID int) (*Order, error) {
	order := &Order{}
	err := os.db.QueryRow("SELECT id, user_id, amount, created_at FROM orders WHERE id = ?", orderID).Scan(&order.ID, &order.UserID, &order.Amount, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	return order, nil
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

	// Create orders table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Orders table created")

	orderService := NewOrderService(db)

	// Create an example order
	orderID, err := orderService.CreateOrder(1