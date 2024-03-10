package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Payment struct {
	ID        int
	OrderID   int
	Amount    float64
	CreatedAt time.Time
}

type PaymentService struct {
	db *sql.DB
}

func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{
		db: db,
	}
}

func (ps *PaymentService) ProcessPayment(orderID int, amount float64) (int64, error) {
	result, err := ps.db.Exec("INSERT INTO payments (order_id, amount, created_at) VALUES (?, ?, ?)", orderID, amount, time.Now())
	if err != nil {
		return 0, err
	}

	paymentID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return paymentID, nil
}

func (ps *PaymentService) GetPayment(paymentID int) (*Payment, error) {
	payment := &Payment{}
	err := ps.db.QueryRow("SELECT id, order_id, amount, created_at FROM payments WHERE id = ?", paymentID).Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return payment, nil
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

	// Create payments table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payments (
		id INT AUTO_INCREMENT PRIMARY KEY,
		order_id INT NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Payments table created")

	paymentService := NewPaymentService(db)

	// Process an example payment
	paymentID,