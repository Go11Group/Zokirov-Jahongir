package stoarge

import (
	"database/sql"
	"errors"

	"example.com/m/service"
)

func CreateUser(db *sql.DB, user service.User) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    _, err = tx.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
    return err
}

func GetUser(db *sql.DB, id int) (*service.User, error) {
    var user service.User
    err := db.QueryRow("SELECT id, username, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

func UpdateUser(db *sql.DB, user service.User) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()

    _, err = tx.Exec("UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4", user.Username, user.Email, user.Password, user.ID)
    return err
}

func DeleteUser(db *sql.DB, id int) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()

    _, err = tx.Exec("DELETE FROM users WHERE id=$1", id)
    return err
}
