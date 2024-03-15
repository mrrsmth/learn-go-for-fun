package goroutines

import (
	"fmt"
	"time"
)

func MainGoroutines() {
	fmt.Println("init")
	for i := 1; i < 7; i++ {
		go factorial(i)
	}
	time.Sleep(1 * time.Second)
	// fmt.Scanln()
	fmt.Println("The End")

	for i := 1; i < 7; i++ {

		go func(n int) {
			result := 1
			for j := 1; j <= n; j++ {
				result *= j
			}
			fmt.Println(n, "-", result)
		}(i)
	}

	fmt.Scanln()
	fmt.Println("The End")

	for i := 1; i < 7; i++ {
		num := i
		go func() {
			fmt.Println(num)
		}()
	}
	time.Sleep(1 * time.Second)

}

func factorial(n int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}
