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

func CreateProduct(db *gorm.DB, newProduct Product) error {
	err := db.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, id int, updatedProduct Product) error {
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

func DeleteProduct(db *gorm.DB, id int) error {
	product := Product{}
	result := db.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetProduct(db *gorm.DB, id int) (Product, error) {
	product := Product{}
	result := db.First(&product, id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

func GetNextProductID(db *gorm.DB) (int, error) {
	var product Product
	result := db.Order("id desc").First(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID + 1, nil
}

func Gorm() {
	dsn := "root:123456@tcp(localhost:3306)/productdb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// createdProduct := Product{ID: 1, Model: "Ноутбук", Company: "ASUS", Price: 1999}
	// err = CreateProduct(db, createdProduct)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	id, err := GetNextProductID(db)
	if err != nil {
		log.Fatal(err)
	}

	newProduct := Product{ID: id, Model: "PupoBook", Company: "PupoPro", Price: 1555999}
	err = CreateProduct(db, newProduct)
	if err != nil {
		log.Fatal(err)
	}

	fetchedProduct, err := GetProduct(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Полученный продукт: %+v\n", fetchedProduct)

	updatedProduct := Product{Model: "Обновленный ноутбук", Company: "ASUS", Price: 2499}
	err = UpdateProduct(db, 1, updatedProduct)
	if err != nil {
		log.Fatal(err)
	}

	// err = DeleteProduct(db, 10)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var products []Product
	err = db.Find(&products).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Итоговая база данных:")
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}

func GormStart() {
	Gorm()
}
