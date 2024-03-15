package chanmetanit

import (
	"fmt"
	"sync"
	"time"
)

func Chan() {
	// Небуфферизированные каналы

	// 1
	// intCh := make(chan int)

	// go factorial(5, intCh)                  // вызов горутины
	// fmt.Println("Данные получены", <-intCh) // получение данных из канала
	// fmt.Println("The End")

	// 2

	// intCh := make(chan int)

	// go square(intCh)                   // square ожидает получения через канал
	// intCh <- 4                         // отправляем в канал число
	// fmt.Println("result := ", <-intCh) // получаем из канала результат
	// fmt.Println("The End")

	// Буферизированные каналы

	// intCh := make(chan int, 3)
	// intCh <- 10
	// intCh <- 3
	// intCh <- 24
	// fmt.Println(<-intCh) // 10
	// fmt.Println(<-intCh) // 3
	// fmt.Println(<-intCh) //24
	// fmt.Println(intCh)

	// Однонаправленные каналы

	// Определение канала только для отправки данных:
	// var inCh chan<- int
	// Определение канала только для получения данных:
	// var outCh <-chan int

	// intCh := make(chan int, 2)
	// go factorialSingle(5, intCh)
	// fmt.Println(<-intCh)
	// fmt.Println("The End")

	// Возвращение канала

	// fmt.Println("Start")
	// // создание канала и получение из него данных
	// fmt.Println(<-createChan(5)) // разблокировка
	// fmt.Println("End")

	// Закрытие канала

	// intCh := make(chan int, 3)
	// intCh <- 10
	// intCh <- 3
	// close(intCh) // канал закрыт

	// for i := 0; i < cap(intCh); i++ {
	// 	if val, opened := <-intCh; opened {
	// 		fmt.Println(val)
	// 	} else {
	// 		fmt.Println("Channel closed!")
	// 	}
	// }

	// Синхронизация
	// 1
	// intCh := make(chan int)
	// go factorial(5, intCh)
	// fmt.Println(<-intCh)

	// results := make(map[int]int)
	// structCh := make(chan struct{})

	// go factorialS(5, structCh, results)

	// <-structCh // ожидаем закрытия канала structCh

	// for i, v := range results {
	// 	fmt.Println(i, " - ", v)
	// }

	// WaitGroup
	// Создаем WaitGroup счетчик
	var wg sync.WaitGroup
	wg.Add(2) // Устанавливаем счетчик равным 2, так как у нас будет 2 горутины

	// Функция, которая будет выполнять работу в горутине
	work := func(id int) {
		defer wg.Done() // Уменьшаем счетчик на 1 при завершении работы горутины
		fmt.Printf("Горутина %d начала выполнение\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение\n", id)
	}

	// Запускаем первую горутину
	go work(1)
	// Запускаем вторую горутину
	go work(2)

	// Ожидаем завершения обеих горутин
	wg.Wait()
	fmt.Println("Горутины завершили выполнение")

}

func factorialS(n int, ch chan struct{}, results map[int]int) {
	defer close(ch) // закрываем канал после завершения горутины
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}

func factorial(n int, ch chan int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)

	ch <- result // отправка данных в канал
}

func createChan(n int) chan int {
	ch := make(chan int) // создаем канал

	go func() { // запускаем новую горутину
		ch <- n // отправляем данные в канал
	}()

	return ch // возвращаем канал
}

func createChanOne(n int) chan int {
	ch := make(chan int, 1) // создаем канал с буфером емкостью 1
	ch <- n                 // отправляем данные в канал
	return ch               // возвращаем канал
}

func factorialSingle(n int, ch chan<- int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func square(ch chan int) {

	num := <-ch // получаем из канала число
	fmt.Println("num := ", num)
	ch <- num * num // обратно отправляем квадрат числа
}
