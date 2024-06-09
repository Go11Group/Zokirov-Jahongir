package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myproject/models"

	"github.com/gorilla/mux"
)

func GetProblems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Problems)
}

func GetProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range models.Problems {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Problem not found", http.StatusNotFound)
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var problem models.Problem
	json.NewDecoder(r.Body).Decode(&problem)
	problem.ID = len(models.Problems) + 1
	models.Problems = append(models.Problems, problem)
	json.NewEncoder(w).Encode(problem)
}

func UpdateProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range models.Problems {
		if item.ID == id {
			var problem models.Problem
			json.NewDecoder(r.Body).Decode(&problem)
			problem.ID = id
			models.Problems[index] = problem
			json.NewEncoder(w).Encode(problem)
			return
		}
	}
	http.Error(w, "Problem not found", http.StatusNotFound)
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range models.Problems {
		if item.ID == id {
			models.Problems = append(models.Problems[:index], models.Problems[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Problems)
}
