package stoarge

import (
    "database/sql"
    "errors"

	"example.com/m/service"
)

func CreateProduct(db *sql.DB, product service.Product) error {
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

    _, err = tx.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.StockQuantity)
    return err
}

func GetProduct(db *sql.DB, id int) (*service.Product, error) {
    var product service.Product
    err := db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    return &product, nil
}

func UpdateProduct(db *sql.DB, product service.Product) error {
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

    _, err = tx.Exec("UPDATE products SET name=$1, description=$2, price=$3, stock_quantity=$4 WHERE id=$5", product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
    return err
}

func DeleteProduct(db *sql.DB, id int) error {
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

    _, err = tx.Exec("DELETE FROM products WHERE id=$1", id)
    return err
}
