package main

import (
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code     string
	Price    uint
	ComeDate time.Time
	Type     string
	Calories float64
}

func main() {
	dbUrl := "postgres://postgres:ertak@localhost:5432/dars?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	log.Println("Database connected and migrated successfully")

	products := []Product{
		{Code: "A1", Price: 100, ComeDate: time.Now(), Type: "Food", Calories: 250.5},
		{Code: "A2", Price: 200, ComeDate: time.Now(), Type: "Drink", Calories: 150.0},
		{Code: "A3", Price: 300, ComeDate: time.Now(), Type: "Snack", Calories: 100.0},
		{Code: "A4", Price: 400, ComeDate: time.Now(), Type: "Dessert", Calories: 350.0},
		{Code: "A5", Price: 500, ComeDate: time.Now(), Type: "Meal", Calories: 450.0},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			panic(err)
		}
	}

	log.Println("Inserted 5 products successfully")
}
