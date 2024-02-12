package basefunc

import "fmt"

func Base() {
	numbers := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)

	filtered := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			filtered = append(filtered, num)
		}
	}
	fmt.Println(filtered)

	includes := false
	for _, num := range numbers {
		if num == 3 {
			includes = true
			break
		}
	}
	fmt.Println(includes)
}

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

func monkeyCount(n int) []int {
	output := []int{} // Пустой срез
	for i := 1; i < n+1; i++ {
		output = append(output, i)
	}
	return output
}

func Litres(time float64) int {
	output := int(time * 0.5)
	return output
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
