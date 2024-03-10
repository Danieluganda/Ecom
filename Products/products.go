package main

/*
 Takes care of all the product related information. 
 This includes the watch details like name, brand, images, etc.
*/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int
	Name  string
	Brand string
	Image string
}

func main() {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(hostname:port)/db_ecom")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure the connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Example operations with the Product Service

	// Insert a new product
	product := Product{Name: "Watch 1", Brand: "Brand 1", Image: "image1.jpg"}
	insertProduct(db, product)

	// Get a product by ID
	id := 1
	result, err := getProductByID(db, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product with ID %d: %+v\n", id, result)

	// Update a product
	product.Brand = "Updated Brand"
	err = updateProduct(db, product)
	if err != nil {
		log.Fatal(err)
	}

	// Delete a product by ID
	idToDelete := 1
	err = deleteProductByID(db, idToDelete)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product with ID %d deleted\n", idToDelete)
}

// Insert a new product into the database
func insertProduct(db *sql.DB, product Product) {
	query := "INSERT INTO products (name, brand, image) VALUES (?, ?, ?)"
	_, err := db.Exec(query, product.Name, product.Brand, product.Image)
	if err != nil {
		log.Fatal(err)
	}
}

// Get a product by its ID from the database
func getProductByID(db *sql.DB, id int) (Product, error) {
	var product Product
	query := "SELECT id, name, brand, image FROM products WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Brand, &product.Image)
	if err != nil {
		return product, err