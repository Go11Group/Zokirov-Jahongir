package main

import (
    "database/sql"
    "fmt"

    _"github.com/google/uuid"
    _"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "ertak"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Gender string
	Course string
}


func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		fmt.Println("Hello")
		panic(err)
	}

    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
    _, err = db.Exec(createUsersTable)
    if err != nil {
        log.Fatal(err)
    }

    createOrdersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        product_name TEXT NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`
    _, err = db.Exec(createOrdersTable)
    if err != nil {
        log.Fatal(err)
    }

    insertUsers := `
    INSERT INTO users (name) VALUES ('Sharapat'), ('Bobir');`
    _, err = db.Exec(insertUsers)
    if err != nil {
        log.Fatal(err)
    }

    insertOrders := `
    INSERT INTO orders (user_id, product_name) VALUES 
    (1, 'Laptop'), 
    (1, 'Smartphone'), 
    (2, 'Tablet');`
    _, err = db.Exec(insertOrders)
    if err != nil {
        log.Fatal(err)
    }

    query := `
    SELECT users.name, orders.product_name 
    FROM users 
    JOIN orders ON users.id = orders.user_id;`
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Users and their orders:")
    for rows.Next() {
        var userName string
        var productName string
        err := rows.Scan(&userName, &productName)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s ordered %s\n", userName, productName)
    }

    queryRow := `
    SELECT users.name, orders.product_name 
    FROM users 
    JOIN orders ON users.id = orders.user_id 
    WHERE orders.id = 1;`
    var userName string
    var productName string
    err = db.QueryRow(queryRow).Scan(&userName, &productName)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Order ID 1: %s ordered %s\n", userName, productName)
}