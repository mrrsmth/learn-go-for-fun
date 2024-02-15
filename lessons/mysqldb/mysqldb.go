package mysqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

// Добавление данных
func AddItem() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)",
		"iPhone X", "Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

// Получение данных

func GetItems() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from productdb.Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}

func Reload() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	// обновляем строку с id=1
	result, err := db.Exec("update productdb.Products set price = ? where id = ?", 11000, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func Del() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productdb")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("delete from productdb.Products where id = 10")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
