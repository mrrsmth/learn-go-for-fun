package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID      int
	Model   string
	Company string
	Price   int
}

func InfoMain() {
	deletedProduct, err := DeleteProduct(2)
	if err != nil {
		panic(err)
	}
	fmt.Println((deletedProduct))
}

func Add() {
	fmt.Println("init")
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO productdb.products (model, company, price) VALUES (?, ?, ?)",
		"iPhone X", "Apple", 72000)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("ID добавленного объекта:", lastInsertID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Количество затронутых строк:", rowsAffected)

	// Вывод данных из таблицы
	rows, err := db.Query("SELECT * FROM productdb.products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var model string
		var company string
		var price int
		err := rows.Scan(&id, &model, &company, &price)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Модель:", model)
		fmt.Println("Компания:", company)
		fmt.Println("Цена:", price)
		fmt.Println("--------------------")
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

}

func Scan() {
	fmt.Println("init")
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM productdb.products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var model string
		var company string
		var price int

		err := rows.Scan(&id, &model, &company, &price)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Модель:", model)
		fmt.Println("Компания:", company)
		fmt.Println("Цена:", price)
		fmt.Println("--------------------")
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func DeleteProduct(id int) (Product, error) {
	var deletedProduct Product

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")
	if err != nil {
		return deletedProduct, err
	}
	defer db.Close()

	// Получаем данные удаленного продукта перед удалением
	err = db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&deletedProduct.ID, &deletedProduct.Model,
		&deletedProduct.Company, &deletedProduct.Price)
	if err != nil {
		return deletedProduct, err
	}

	// Удаляем продукт из базы данных
	_, err = db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return deletedProduct, err
	}

	return deletedProduct, nil
}
