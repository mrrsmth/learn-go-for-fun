package goroutines

import "fmt"

func MainGoroutines() {
	// 	var wg sync.WaitGroup // Создаем переменную WaitGroup

	// 	for i := 1; i < 7; i++ {
	// 		wg.Add(1) // Увеличиваем счетчик WaitGroup на 1
	// 		go Factorial(i, &wg)
	// 	}

	// 	wg.Wait() // Ждем завершения всех горутин
	// 	fmt.Println("The End")
	// }

	// func Factorial(n int, wg *sync.WaitGroup) {
	// 	defer wg.Done() // Уменьшаем счетчик WaitGroup на 1 при завершении горутины

	// 	if n < 1 {
	// 		fmt.Println("Invalid input number")
	// 		return
	// 	}

	// 	result := 1
	// 	for i := 1; i <= n; i++ {
	// 		result *= i
	// 	}

	// 	fmt.Println(n, "-", result)
	for i := 1; i < 7; i++ {
		go Factorial(i)
	}
	// fmt.Scanln() // ждем ввода пользователя
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

}

func Factorial(n int) {
	if n < 1 {
		fmt.Println("Unvalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}
