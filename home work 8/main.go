package main

type Student struct {
    ID       string
    Name     string
    Age      int
    Gender   string
    CourseID string
}

type Course struct {
    ID   string
    Name string
}

package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/google/uuid"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    dbname   = "dars"
    password = "ertak"
)

func main() {
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
        host, port, user, dbname, password)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
}

func createStudent(db *sql.DB, student Student) error {
    query := `INSERT INTO student (id, name, age, gender, course_id) VALUES ($1, $2, $3, $4, $5)`
    _, err := db.Exec(query, uuid.NewString(), student.Name, student.Age, student.Gender, student.CourseID)
    return err
}

func createCourse(db *sql.DB, course Course) error {
    query := `INSERT INTO course (id, name) VALUES ($1, $2)`
    _, err := db.Exec(query, uuid.NewString(), course.Name)
    return err
}

unc getStudentByID(db *sql.DB, id string) (Student, error) {
    student := Student{}
    query := `SELECT id, name, age, gender, course_id FROM student WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&student.ID, &student.Name, &student.Age, &student.Gender, &student.CourseID)
    return student, err
}

func getCourseByID(db *sql.DB, id string) (Course, error) {
    course := Course{}
    query := `SELECT id, name FROM course WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&course.ID, &course.Name)
    return course, err
}

func updateStudent(db *sql.DB, student Student) error {
    query := `UPDATE student SET name = $1, age = $2, gender = $3, course_id = $4 WHERE id = $5`
    _, err := db.Exec(query, student.Name, student.Age, student.Gender, student.CourseID, student.ID)
    return err
}

func updateCourse(db *sql.DB, course Course) error {
    query := `UPDATE course SET name = $1 WHERE id = $2`
    _, err := db.Exec(query, course.Name, course.ID)
    return err
}

func deleteStudent(db *sql.DB, id string) error {
    query := `DELETE FROM student WHERE id = $1`
    _, err := db.Exec(query, id)
    return err
}

func deleteCourse(db *sql.DB, id string) error {
    query := `DELETE FROM course WHERE id = $1`
    _, err := db.Exec(query, id)
    return err
}
