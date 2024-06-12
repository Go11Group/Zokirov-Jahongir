package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"myproject2month/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CourseHandler struct {
	db *sql.DB
}

func NewCourseHandler(db *sql.DB) *CourseHandler {
	return &CourseHandler{db: db}
}

func (h *CourseHandler) GetCourses(c *gin.Context) {
	rows, err := h.db.Query("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM Courses WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		courses = append(courses, course)
	}

	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.CourseID = uuid.New()
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	_, err := h.db.Exec("INSERT INTO Courses (course_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		course.CourseID, course.Title, course.Description, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) GetCourse(c *gin.Context) {
	id := c.Param("id")
	row := h.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM Courses WHERE course_id = $1 AND deleted_at IS NULL", id)

	var course models.Course
	err := row.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.UpdatedAt = time.Now()
	_, err := h.db.Exec("UPDATE Courses SET title = $1, description = $2, updated_at = $3 WHERE course_id = $4 AND deleted_at IS NULL",
		course.Title, course.Description, course.UpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	deletedAt := time.Now()
	_, err := h.db.Exec("UPDATE Courses SET deleted_at = $1 WHERE course_id = $2 AND deleted_at IS NULL", deletedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
