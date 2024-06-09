package main

import (
	"log"
	"net/http"

	"myproject/handlers"
	"myproject/models"

	"github.com/gorilla/mux"
)

func main() {
	models.Problems = []models.Problem{
		{ID: 1, Title: "Two Sum", Description: "Find two numbers that add up to a target number.", Difficulty: "Easy"},
		{ID: 2, Title: "Reverse String", Description: "Reverse the given string.", Difficulty: "Easy"},
	}
	models.Users = []models.User{
		{ID: 1, Username: "johndoe", Email: "johndoe@example.com"},
		{ID: 2, Username: "janedoe", Email: "janedoe@example.com"},
	}
	models.SolvedProblems = []models.SolvedProblem{
		{ID: 1, UserID: 1, ProblemID: 1},
		{ID: 2, UserID: 2, ProblemID: 2},
	}

	r := mux.NewRouter()

	r.HandleFunc("/problems", handlers.GetProblems).Methods("GET")
	r.HandleFunc("/problems/{id}", handlers.GetProblem).Methods("GET")
	r.HandleFunc("/problems", handlers.CreateProblem).Methods("POST")
	r.HandleFunc("/problems/{id}", handlers.UpdateProblem).Methods("PUT")
	r.HandleFunc("/problems/{id}", handlers.DeleteProblem).Methods("DELETE")

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/solved_problems", handlers.GetSolvedProblems).Methods("GET")
	r.HandleFunc("/solved_problems/{id}", handlers.GetSolvedProblem).Methods("GET")
	r.HandleFunc("/solved_problems", handlers.CreateSolvedProblem).Methods("POST")
	r.HandleFunc("/solved_problems/{id}", handlers.UpdateSolvedProblem).Methods("PUT")
	r.HandleFunc("/solved_problems/{id}", handlers.DeleteSolvedProblem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
