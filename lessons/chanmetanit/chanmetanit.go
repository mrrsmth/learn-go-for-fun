package chanmetanit

import "fmt"

func Chan() {
	// 1
	// intCh := make(chan int)

	// go factorial(5, intCh)                  // вызов горутины
	// fmt.Println("Данные получены", <-intCh) // получение данных из канала
	// fmt.Println("The End")

	// 2

	intCh := make(chan int)

	go square(intCh)                   // square ожидает получения через канал
	intCh <- 4                         // отправляем в канал число
	fmt.Println("result := ", <-intCh) // получаем из канала результат
	fmt.Println("The End")
}

func factorial(n int, ch chan int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)

	ch <- result // отправка данных в канал
}

func square(ch chan int) {

	num := <-ch // получаем из канала число
	fmt.Println("num := ", num)
	ch <- num * num // обратно отправляем квадрат числа
}
