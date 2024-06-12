package main

import (
	"database/sql"
	"log"
	"myproject/functions"
	"myproject/strcts"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Filter struct {
	Age           string
	Gender        string
	Nation        string
	Field         string
	Limit, Offset string
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.1"})

	r.GET("/getAll", getAll)

	r.Run(":8080")
}

func getAll(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		filter []string
	)

	age := c.Query("age")
	gender := c.Query("gender")
	nation := c.Query("nation")
	field := c.Query("field")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	baseQuery := `SELECT id, first_name, last_name, age, gender, nation, field, parent_name, city FROM People`

	if gender != "" {
		params["gender"] = gender
		filter = append(filter, "gender = :gender")
	}

	if nation != "" {
		params["nation"] = nation
		filter = append(filter, "nation = :nation")
	}

	if field != "" {
		params["field"] = field
		filter = append(filter, "field = :field")
	}

	if age != "" {
		params["age"] = age
		filter = append(filter, "age = :age")
	}

	if limitStr != "" {
		params["limit"] = limitStr
		limit = " LIMIT :limit"
	}

	if offsetStr != "" {
		params["offset"] = offsetStr
		if limit != "" {
			limit += " OFFSET :offset"
		} else {
			limit = " OFFSET :offset"
		}
	}

	if len(filter) > 0 {
		baseQuery += " WHERE " + strings.Join(filter, " AND ")
	}

	query := baseQuery + limit

	query, arr = functions.ReplaceQueryParams(query, params)

	log.Printf("Executing query: %s with params: %v", query, arr)

	connStr := "user=postgres dbname=dars host=localhost password=ertak port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		c.JSON(500, gin.H{"error": "Error opening database"})
		return
	}
	defer db.Close()

	rows, err := db.Query(query, arr...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(500, gin.H{"error": "Error executing query"})
		return
	}
	defer rows.Close()

	var people []strcts.People
	for rows.Next() {
		var p strcts.People
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Gender, &p.Nation, &p.Field, &p.ParentName, &p.City); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(500, gin.H{"error": "Error scanning row"})
			return
		}
		people = append(people, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error in row iteration: %v", err)
		c.JSON(500, gin.H{"error": "Error in row iteration"})
		return
	}

	c.JSON(200, people)
}
