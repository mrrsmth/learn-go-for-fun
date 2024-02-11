package recursivefunctions

func Factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
