package models

type Problem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SolvedProblem struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProblemID int `json:"problem_id"`
}

var Problems []Problem
var Users []User
var SolvedProblems []SolvedProblem
