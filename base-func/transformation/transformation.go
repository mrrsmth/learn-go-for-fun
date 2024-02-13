package transformation

import (
	"fmt"
	"reflect"
	"strconv"
)

func TransformationMain() {
	// radius := 5
	// circumference := 2 * math.Pi * float64(radius)
	// fmt.Println(math.Pi)
	// fmt.Println("Радиус:", radius)
	// fmt.Println("Длина окружности:", circumference)
	// fmt.Println(math.Sqrt(16)) // 4

	// var myInt int = 42
	// var myFloat float64 = float64(myInt)

	// var myFloat float64 = 3.14
	// var myInt int = int(myFloat)

	var myInt int = 42
	var myString string = strconv.Itoa(myInt)

	fmt.Println(myInt)
	fmt.Println(myString)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
	fmt.Println(err)
	fmt.Println("Тип переменной number:", reflect.TypeOf(f))
	fmt.Println("Тип переменной text:", reflect.TypeOf(err))

}
