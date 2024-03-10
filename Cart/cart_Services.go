package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Cart struct {
	ID     int
	UserID string
	ItemID int
	Quantity int
}

type CartService struct {
	db *sql.DB
}

func NewCartService(db *sql.DB) *CartService {
	return &CartService{
		db: db,
	}
}

func (cs *CartService) AddItem(userID string, itemID, quantity int) error {
	_, err := cs.db.Exec("INSERT INTO carts (user_id, item_id, quantity) VALUES (?, ?, ?)", userID, itemID, quantity)
	return err
}

func (cs *CartService) RemoveItem(userID string, itemID int) error {
	_, err := cs.db.Exec("DELETE FROM carts WHERE user_id = ? AND item_id = ?", userID, itemID)
	return err
}

func (cs *CartService) UpdateItem(userID string, itemID, quantity int) error {
	_, err := cs.db.Exec("UPDATE carts SET quantity = ? WHERE user_id = ? AND item_id = ?", quantity, userID, itemID)
	return err
}

func (cs *CartService) GetTotal(userID string) (float64, error) {
	var total float64

	err := cs.db.QueryRow("SELECT SUM(quantity * price) FROM carts JOIN items ON carts.item_id = items.id WHERE user_id = ?", userID).Scan(&total)

	return total, err
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

	// Create carts table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS carts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		item_id INT NOT NULL,
		quantity INT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Carts table created")

	c