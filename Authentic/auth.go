package main
//Authentication Service: Takes care of user authentication (login, logout)
// and user authorization for accessing the different services 
//while fetching data from MySQL database `db_ecom`.
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
}

type AuthService struct {
	db *sql.DB
}

func NewAuthService() (*AuthService, error) {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(database-host:port)/db_ecom")
	if err != nil {
		return nil, err
	}

	return &AuthService{
		db: db,
	}, nil
}

func (auth *AuthService) RegisterUser(username, password string) error {
	// Check if the username already exists
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int
	err := auth.db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("Username already exists")
	}

	// Insert the new user into the database
	insertQuery := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = auth.db.Exec(insertQuery, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (auth *AuthService) Login(username, password string) error {
	// Check if the username and password are valid
	query := "SELECT COUNT(*) FROM users WHERE username = ? AND password = ?"
	var count int
	err := auth.db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Invalid credentials")
	}

	// Perform login logic (e.g., set session, update logged_in status, etc.)
	// ...

	return nil
}

func (auth *AuthService) Logout(username string) error {
	// Perform logout logic (e.g., destroy session, update logged_in status, etc.)
	// ...

	return nil
}

func main() {
	// Example usage:
	authService, err := NewAuthService()
	if err != nil {
		log.Fatal