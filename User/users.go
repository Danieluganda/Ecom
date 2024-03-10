// User Service: Handle user data and provide functions for registering, updating, retrieving and deleting users.
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	Username  string
	Email     string
	FirstName string
	LastName  string
}

type UserService struct {
	db *sql.DB
}

func NewUserService() (*UserService, error) {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		return nil, err
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &UserService{db: db}, nil
}

func (us *UserService) RegisterUser(user User) error {
	// Prepare the SQL statement
	stmt, err := us.db.Prepare("INSERT INTO users(username, email, first_name, last_name) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(user.Username, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	fmt.Println("User registered successfully!")
	return nil
}

func (us *UserService) UpdateUser(user User) error {
	// Prepare the SQL statement
	stmt, err := us.db.Prepare("UPDATE users SET username=?, email=?, first_name=?, last_name=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(user.Username, user.Email, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}

	fmt.Println("User updated successfully!")
	return nil
}

func (us *UserService) RetrieveUser(id int) (User, error) {
	var user User

	// Prepare the SQL statement
	stmt, err := us.db.Prepare("SELECT id, username, email, first_name, last_name FROM users WHERE id=?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	
	// Execute the SQL statement.