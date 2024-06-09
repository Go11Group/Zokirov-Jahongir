package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myproject/models"

	"github.com/gorilla/mux"
)

func GetSolvedProblems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SolvedProblems)
}

func GetSolvedProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range models.SolvedProblems {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Solved problem not found", http.StatusNotFound)
}

func CreateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var solvedProblem models.SolvedProblem
	json.NewDecoder(r.Body).Decode(&solvedProblem)
	solvedProblem.ID = len(models.SolvedProblems) + 1
	models.SolvedProblems = append(models.SolvedProblems, solvedProblem)
	json.NewEncoder(w).Encode(solvedProblem)
}

func UpdateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range models.SolvedProblems {
		if item.ID == id {
			var solvedProblem models.SolvedProblem
			json.NewDecoder(r.Body).Decode(&solvedProblem)
			solvedProblem.ID = id
			models.SolvedProblems[index] = solvedProblem
			json.NewEncoder(w).Encode(solvedProblem)
			return
		}
	}
	http.Error(w, "Solved problem not found", http.StatusNotFound)
}

func DeleteSolvedProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range models.SolvedProblems {
		if item.ID == id {
			models.SolvedProblems = append(models.SolvedProblems[:index], models.SolvedProblems[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.SolvedProblems)
}
