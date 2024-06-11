package handlers

import (
	"net/http"
	"strconv"

	"myproject2/models"

	"github.com/gin-gonic/gin"
)

func GetProblems(c *gin.Context) {
	c.JSON(http.StatusOK, models.Problems)
}

func GetProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range models.Problems {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
}

func CreateProblem(c *gin.Context) {
	var problem models.Problem
	if err := c.BindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	problem.ID = len(models.Problems) + 1
	models.Problems = append(models.Problems, problem)
	c.JSON(http.StatusCreated, problem)
}

func UpdateProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.Problems {
		if item.ID == id {
			var problem models.Problem
			if err := c.BindJSON(&problem); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			problem.ID = id
			models.Problems[index] = problem
			c.JSON(http.StatusOK, problem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
}

func DeleteProblem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.Problems {
		if item.ID == id {
			models.Problems = append(models.Problems[:index], models.Problems[index+1:]...)
			c.JSON(http.StatusOK, models.Problems)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
}
