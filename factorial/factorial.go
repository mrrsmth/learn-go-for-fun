package factorial

func Factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
func main() {

	// fmt.Println(factorial(4)) // 24
	// fmt.Println(factorial(5)) // 120
	// fmt.Println(factorial(6)) // 720
}
