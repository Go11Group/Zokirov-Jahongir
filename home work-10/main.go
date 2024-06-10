package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type Product struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	StockQuantity int
}

func main() {
	connStr := "user=postgres dbname=dars sslmode=disable password=ertak host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := User{Username: "john_doe", Email: "john@example.com", Password: "password123"}
	products := []Product{
		{Name: "Product 1", Description: "Description 1", Price: 9.99, StockQuantity: 100},
		{Name: "Product 2", Description: "Description 2", Price: 19.99, StockQuantity: 200},
	}

	if err := createUserWithProducts(db, user, products); err != nil {
		log.Fatal(err)
	}

	fmt.Println("User and products created successfully")
}

func createUserWithProducts(db *sql.DB, user User, products []Product) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var userID int
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	if err := tx.QueryRow(query, user.Username, user.Email, user.Password).Scan(&userID); err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range products {
		var productID int
		query := `INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4) RETURNING id`
		if err := tx.QueryRow(query, product.Name, product.Description, product.Price, product.StockQuantity).Scan(&productID); err != nil {
			tx.Rollback()
			return err
		}

		query = `INSERT INTO user_products (user_id, product_id) VALUES ($1, $2)`
		if _, err := tx.Exec(query, userID, productID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func updateUserWithProducts(db *sql.DB, user User, products []Product) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	if _, err := tx.Exec(query, user.Username, user.Email, user.Password, user.ID); err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range products {
		query := `UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5`
		if _, err := tx.Exec(query, product.Name, product.Description, product.Price, product.StockQuantity, product.ID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func deleteUserWithProducts(db *sql.DB, userID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM user_products WHERE user_id = $1`
	if _, err := tx.Exec(query, userID); err != nil {
		tx.Rollback()
		return err
	}

	query = `DELETE FROM users WHERE id = $1`
	if _, err := tx.Exec(query, userID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
