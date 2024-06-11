package handlers

import (
	"net/http"
	"strconv"

	"myproject2/models"

	"github.com/gin-gonic/gin"
)

func GetSolvedProblems(c *gin.Context) {
	c.JSON(http.StatusOK, models.SolvedProblems)
}

func GetSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range models.SolvedProblems {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Solved problem not found"})
}

func CreateSolvedProblem(c *gin.Context) {
	var solvedProblem models.SolvedProblem
	if err := c.BindJSON(&solvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	solvedProblem.ID = len(models.SolvedProblems) + 1
	models.SolvedProblems = append(models.SolvedProblems, solvedProblem)
	c.JSON(http.StatusCreated, solvedProblem)
}

func UpdateSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.SolvedProblems {
		if item.ID == id {
			var solvedProblem models.SolvedProblem
			if err := c.BindJSON(&solvedProblem); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			solvedProblem.ID = id
			models.SolvedProblems[index] = solvedProblem
			c.JSON(http.StatusOK, solvedProblem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Solved problem not found"})
}

func DeleteSolvedProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.SolvedProblems {
		if item.ID == id {
			models.SolvedProblems = append(models.SolvedProblems[:index], models.SolvedProblems[index+1:]...)
			c.JSON(http.StatusOK, models.SolvedProblems)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Solved problem not found"})
}
