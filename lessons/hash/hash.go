package hash

import "fmt"

func HashInfo() {
	// m := make(map[string]int)
	// fmt.Println((m))

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	slice := []int{1, 2, 3}
	for key, value := range people {
		fmt.Println(key, value)
		slice = append(slice, value)
		// if value == 2 {
		// fmt.Println("Check!")
		// }
	}

	// slice = append(slice, )
	fmt.Println(slice)

	// fmt.Println(people)
	// if val, ok := people["Alice"]; ok {
	// 	fmt.Println(val)
	// }

	// people1 := map[string]int{
	// 	"Tom":   1,
	// 	"Bob":   2,
	// 	"Sam":   4,
	// 	"Alice": 8,
	// }

	// fmt.Println(people1)

	var people3 = map[string]int{"Tom": 1, "Bob": 2}
	people["Kate"] = 128
	fmt.Println(people3)

	var people4 = map[string]int{"Tom": 1, "Bob": 2, "Sam": 8}
	delete(people, "Bob")
	fmt.Println(people4)
}
