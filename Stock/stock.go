package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Item represents an inventory item
type Item struct {
	ID    int
	Name  string
	Stock int
}

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db_ecom")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Perform a simple query to test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Create an inventory table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS inventory (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		stock INT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inventory table created")

	// Insert sample data
	_, err = db.Exec(`INSERT INTO inventory (name, stock) VALUES
		('Item A', 10),
		('Item B', 5)
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sample data inserted")

	// Query all items in the inventory
	rows, err := db.Query("SELECT id, name, stock FROM inventory")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Inventory items:")
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Stock)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Stock: %d\n", item.ID, item.Name, item.Stock)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
