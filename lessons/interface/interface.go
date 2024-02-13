package interfaceLesson

import "fmt"

type Vehicle interface {
	move()
}

// структура "Автомобиль"
type Car struct {
	brand   string
	model   string
	color   string
	mileage int
	year    int
}

type Stream interface {
	read() string
	write(string)
	close()
}

// структура "Самолет"
type Aircraft struct {
}

func (c Car) move() {
	fmt.Println(c.brand)
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func printCars(cars []Car) {
	for _, c := range cars {
		fmt.Println("Марка:", c.brand)
		fmt.Println("Модель:", c.model)
		fmt.Println("Цвет:", c.color)
		fmt.Println("Пробег:", c.mileage)
		fmt.Println("Год выпуска:", c.year)
		fmt.Println("---------------------")
	}
	fmt.Println("Самолет летит")
}

func InterfaceLesson() {
	// var car Vehicle = Car{
	// 	brand:   "Toyota",
	// 	model:   "Camry",
	// 	color:   "Blue",
	// 	mileage: 10000,
	// 	year:    2020,
	// }
	// var boing Vehicle = Aircraft{}
	// car.move()
	// boing.move()

	// cars := []Car{
	// 	{brand: "Toyota", model: "Camry", color: "Blue", mileage: 10000, year: 2020},
	// 	{brand: "Honda", model: "Accord", color: "Red", mileage: 15000, year: 2019},
	// 	{brand: "Ford", model: "Mustang", color: "Yellow", mileage: 5000, year: 2021},
	// }

	// printCars(cars)
}
