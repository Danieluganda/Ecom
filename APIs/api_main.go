package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type API struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

type APIService struct {
	db *sql.DB
}

func NewAPIService(db *sql.DB) *APIService {
	return &APIService{
		db: db,
	}
}

func (as *APIService) CreateAPI(name, description string) (int64, error) {
	result, err := as.db.Exec("INSERT INTO apis (name, description, created_at) VALUES (?, ?, ?)", name, description, time.Now())
	if err != nil {
		return 0, err
	}

	apiID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return apiID, nil
}

func (as *APIService) GetAPI(apiID int) (*API, error) {
	api := &API{}
	err := as.db.QueryRow("SELECT id, name, description, created_at FROM apis WHERE id = ?", apiID).Scan(&api.ID, &api.Name, &api.Description, &api.CreatedAt)
	if err != nil {
		return nil, err
	}

	return api, nil
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

	// Create apis table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS apis (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("APIs table created")

	apiService := NewAPIService(db)

	// Create an example API
	apiID, err := apiService.CreateAPI("ecommerce-api", "API for managing ecommerce operations")
	if err !=