package main

import (
	"fmt"
	"log"

	"example.com/m/postgres"
	"example.com/m/service"
	"example.com/m/stoarge"
)

func main() {
    db := postgres.ConnectToDatabase()
    defer db.Close()

    newUser := service.User{
        Username: "newuser",
        Email:    "newuser@example.com",
        Password: "password123",
    }

    err := stoarge.CreateUser(db, newUser)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    user, err := stoarge.GetUser(db, 1)
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }
    fmt.Printf("Retrieved User: %+v\n", user)

    user.Username = "updateduser"
    err = stoarge.UpdateUser(db, *user)
    if err != nil {
        log.Fatalf("Failed to update user: %v", err)
    }

    err = stoarge.DeleteUser(db, user.ID)
    if err != nil {
        log.Fatalf("Failed to delete user: %v", err)
    }

    newProduct := service.Product{
        Name:          "newproduct",
        Description:   "product description",
        Price:         99.99,
        StockQuantity: 100,
    }

    err = stoarge.CreateProduct(db, newProduct)
    if err != nil {
        log.Fatalf("Failed to create product: %v", err)
    }

    product, err := stoarge.GetProduct(db, 1)
    if err != nil {
        log.Fatalf("Failed to get product: %v", err)
    }
    fmt.Printf("Retrieved Product: %+v\n", product)

    product.Name = "updatedproduct"
    err = stoarge.UpdateProduct(db, *product)
    if err != nil {
        log.Fatalf("Failed to update product: %v", err)
    }

    err = stoarge.DeleteProduct(db, product.ID)
    if err != nil {
        log.Fatalf("Failed to delete product: %v", err)
    }
}
