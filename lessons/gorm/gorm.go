package gorm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID      int
	Model   string
	Company string
	Price   int
}

func Up(db gorm.DB) {
	updatedProduct := Product{
		Model:   "Ноутбук Pro",
		Company: "ASUS",
		Price:   2999,
	}

	err := UpdateProduct(db, 11, updatedProduct)
	if err != nil {
		log.Fatal(err)
	}
}

func Gorm() {
	// Открываем соединение с базой данных
	dsn := "root:123456@tcp(localhost:3306)/productdb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Получаем объект sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Создаем таблицу "products", если она еще не существует
	// err = db.AutoMigrate(&Product{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Создаем продукт по ID
	createdProduct := Product{ID: 1, Model: "Ноутбук", Company: "ASUS", Price: 1999}
	err = db.Create(&createdProduct).Error
	if err != nil {
		log.Fatal(err)
	}

	// Получаем продукт по ID
	var fetchedProduct Product
	err = db.First(&fetchedProduct, 1).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Полученный продукт: %+v\n", fetchedProduct)

	// Обновляем продукт по ID
	updatedProduct := Product{ID: 1, Model: "Обновленный ноутбук", Company: "ASUS", Price: 2499}
	err = db.Save(&updatedProduct).Error
	if err != nil {
		log.Fatal(err)
	}

	// Удаляем продукт по ID
	err = DeleteProduct(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Вызываем функцию Up для обновления продукта по ID
	Up(*db)
}

func DeleteProduct(db *gorm.DB, id int) error {
	product := Product{}
	result := db.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProduct(db gorm.DB, id int, updatedProduct Product) error {
	product := Product{}
	result := db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return result.Error
	}

	product.Model = updatedProduct.Model
	product.Company = updatedProduct.Company
	product.Price = updatedProduct.Price

	result = db.Save(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GormStart() {
	Gorm()
}
