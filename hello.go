package main

import (
	factorialpacked "example/learn-go/factorial"
	"example/learn-go/recursive-functions"
	hello "example/learn-go/service"
	"example/learn-go/src/math"
	"fmt"
)

func add(x int, y int) int {

	return x + y
}
func multiply(x int, y int) int {

	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {

	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}

func Litres(time float64) int {
	output := int(time * 0.5)
	return output
}

// 10 --> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// 1 --> [1]

// output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func monkeyCount(n int) []int {
	output := []int{} // Пустой срез
	for i := 1; i < n+1; i++ {
		output = append(output, i)
	}
	return output
}

// It("should return \"Odd\" for odd positive numbers", func() {
//     Expect(EvenOrOdd(1)).To(Equal("Odd"))
//   })

//   It("should return \"Even\" for even positive numbers", func() {
//     Expect(EvenOrOdd(2)).To(Equal("Even"))
//   })

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {

	// action(10, 25, add)    // 35
	// action(5, 6, multiply) // 30

	// newArr := []int{1, 2, 3, 4, 5}
	// buba := []int{1, 1}
	// arr := make([]int, len(newArr))
	// copy(arr, newArr)
	// copy(arr, buba)

	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := s1

	// fmt.Println(s1, s2)

	// func Litres(time float64) int {
	// 	time * 0.5

	// }
	// Litres(3)
	// monkeyCount(10)
	EvenOrOdd(1)
	// fmt.Println(factorial(4))
	// fmt.Println(factorial(5))
	// fmt.Println(factorial(4))
	hello.Hello()
	math.StrSlice()
	numbers := []float64{10.0}
	math.Average(numbers)
	fmt.Println(factorialpacked.Factorial(4))
	fmt.Println(recursivefunctions.Factorial(4))
}
