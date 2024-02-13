package chanLesson

import (
	"fmt"
	"sync"
	"time"

)

// var counter int = 0

// func work(number int, ch chan bool) {
// 	counter = 0
// 	for k := 1; k <= 5; k++ {
// 		counter++
// 		fmt.Println("Goroutine", number, "-", counter)
// 	}
// 	ch <- true
// }

func MainChan() {
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	// вызываем горутины
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")

	// ch := make(chan bool) // канал
	// for i := 1; i < 5; i++ {
	// 	go work(i, ch)
	// }
	// // ожидаем завершения всех горутин
	// for i := 1; i < 5; i++ {
	// 	<-ch
	// }
	// fmt.Println("The End")

	// 	intCh := make(chan int)

	// 	go Factorial(5, intCh) // вызов горутины
	// 	fmt.Println(<-intCh)   // получение данных из канала
	// 	fmt.Println("The End")
	// }
	// func Factorial(n int, ch chan int) {

	// 	result := 1
	// 	for i := 1; i <= n; i++ {
	// 		result *= i
	// 	}
	// 	fmt.Println(n, "-", result)

	// 	ch <- result // отправка данных в канал

	// intCh := make(chan int)

	// go func() {
	// 	fmt.Println("Go routine starts")
	// 	intCh <- 5 // блокировка, пока данные не будут получены функцией main
	// }()
	// fmt.Println(<-intCh) // получение данных из канала
	// fmt.Println("The End")

	// 	intCh := make(chan int)

	// 	go factorial(5, intCh) // вызов горутины
	// 	fmt.Println(<-intCh)   // получение данных из канала
	// 	fmt.Println("The End")
	// }
	// func factorial(n int, ch chan int) {

	// 	result := 1
	// 	for i := 1; i <= n; i++ {
	// 		result *= i
	// 	}
	// 	fmt.Println(n, "-", result)

	// 	ch <- result // отправка данных в канал
	// intCh := make(chan int, 3)
	// intCh <- 10
	// // intCh <- 3
	// // intCh <- 24
	// // intCh <- 15 // блокировка - функция main ждет, когда освободится место в канале

	// fmt.Println(<-intCh)
	// fmt.Println("The End")

	// fmt.Println("Start")
	// // создание канала и получение из него данных
	// fmt.Println(<-createChan(5)) // 5
	// fmt.Println("End")
}

//	func createChan(n int) chan int {
//		ch := make(chan int) // создаем канал
//		go func() {
//			ch <- n // отправляем данные в канал
//		}() // запускаем горутину
//		return ch // возвращаем канал
//	}
