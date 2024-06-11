package handlers

import (
	"net/http"
	"strconv"

	"myproject2/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range models.Users {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = len(models.Users) + 1
	models.Users = append(models.Users, user)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.Users {
		if item.ID == id {
			var user models.User
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user.ID = id
			models.Users[index] = user
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range models.Users {
		if item.ID == id {
			models.Users = append(models.Users[:index], models.Users[index+1:]...)
			c.JSON(http.StatusOK, models.Users)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
