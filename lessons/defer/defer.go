package deferl

import "fmt"

func Main() {
	// defer Finish()
	// defer fmt.Println("Program has been started")
	// fmt.Println("Program is working")
	fmt.Println(Divide(15, 5))
	fmt.Println(Divide(15, 5))
	fmt.Println(Divide(4, 0))
	fmt.Println(Divide(15, 5))
	fmt.Println("Program has been finished")
}

func Finish() {
	fmt.Println("Program has been finished")
}

func Divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}
