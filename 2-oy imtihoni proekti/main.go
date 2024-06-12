package main

import (
	"database/sql"
	"log"

	"myproject2month/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=postgres dbname=dars host=localhost password=ertak port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	initDB()
	defer db.Close()

	userHandler := handlers.NewUserHandler(db)
	courseHandler := handlers.NewCourseHandler(db)

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/courses", courseHandler.GetCourses)
	r.POST("/courses", courseHandler.CreateCourse)
	r.GET("/courses/:id", courseHandler.GetCourse)
	r.PUT("/courses/:id", courseHandler.UpdateCourse)
	r.DELETE("/courses/:id", courseHandler.DeleteCourse)

	r.Run(":8080")
}
