package pointers

import (
	"fmt"

)

func Poiners() {
	// fmt.Println("Pointers!")
	x := 10
	ptr := &x
	fmt.Println(ptr)
	fmt.Println(*ptr)

	n := 4
	p := &n
	*p = 25
	fmt.Println(n)

	var pf *float64
	if pf != nil {
		fmt.Println("Value:", *pf)
	}

	z := new(int)
	fmt.Println("Value:", *z) // Value: 0 - значение по умолчанию
	*z = 8                    // изменяем значение
	fmt.Println("Value:", *z)
}
