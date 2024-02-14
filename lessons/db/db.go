package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbMain() {

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
}
