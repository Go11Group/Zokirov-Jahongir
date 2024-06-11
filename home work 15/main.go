package main

import (
	"log"
	"net/http"

	"myproject2/handlers"
	"myproject2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.Problems = []models.Problem{
		{ID: 1, Title: "Two Sum", Description: "Find two numbers that add up to a target number.", Difficulty: "Easy"},
		{ID: 2, Title: "Reverse String", Description: "Reverse the given string.", Difficulty: "Easy"},
	}
	models.Users = []models.User{
		{ID: 1, Username: "johndoe", Email: "johndoe@gmail.com"},
		{ID: 2, Username: "janedoe", Email: "janedoe@gmail.com"},
	}
	models.SolvedProblems = []models.SolvedProblem{
		{ID: 1, UserID: 1, ProblemID: 1},
		{ID: 2, UserID: 2, ProblemID: 2},
	}

	r := gin.Default()

	r.GET("/problems", handlers.GetProblems)
	r.GET("/problems/:id", handlers.GetProblem)
	r.POST("/problems", handlers.CreateProblem)
	r.PUT("/problems/:id", handlers.UpdateProblem)
	r.DELETE("/problems/:id", handlers.DeleteProblem)

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/solved_problems", handlers.GetSolvedProblems)
	r.GET("/solved_problems/:id", handlers.GetSolvedProblem)
	r.POST("/solved_problems", handlers.CreateSolvedProblem)
	r.PUT("/solved_problems/:id", handlers.UpdateSolvedProblem)
	r.DELETE("/solved_problems/:id", handlers.DeleteSolvedProblem)

	log.Fatal(http.ListenAndServe(":8000", r))
}
