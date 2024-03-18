package inter

import (
	"fmt"
	"sync"
)

func Inter() {
	// numJobs := 10
	// numWorkers := 4

	// // Создаем каналы для передачи задач и результатов
	// jobs := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	// // Создаем WaitGroup для ожидания завершения горутин
	// var wg sync.WaitGroup

	// // Создаем несколько горутин-рабочих
	// for i := 1; i <= numWorkers; i++ {
	// 	wg.Add(1)
	// 	go worker(i, jobs, results, &wg)
	// }

	// // Подготовка задач и отправка их в канал jobs
	// for j := 1; j <= numJobs; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// go func() {
	// 	// Дожидаемся завершения всех горутин
	// 	wg.Wait()

	// 	// Закрываем канал results после завершения всех горутин
	// 	close(results)
	// }()

	// // Получение результатов работы из канала results
	// for r := range results {
	// 	fmt.Println("Received result:", r)
	// }

	// fmt.Println("All jobs are completed")
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()

		for val := range ch {
			fmt.Println("Получено значение:", val)
		}
	}()

	wg.Wait()
	fmt.Println("Основная горутина завершена.")
}

// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for j := range jobs {
// 		fmt.Printf("Worker #%d processing job %d\n", id, j)
// 		// Выполнение работы...

// 		// Отправка результата работы в канал results
// 		results <- j * 2
// 	}
// }

// func ticker() {
// 	fmt.Println("start")

// 	ticker := time.NewTicker(1 * time.Second)
// 	done := make(chan bool)

// 	// Горутина 1
// 	go func() {
// 		defer func() {
// 			done <- true
// 		}()

// 		for index := range ticker.C {
// 			fmt.Println("Горутина 1:", index)
// 		}
// 	}()

// 	// Горутина 2
// 	go func() {
// 		defer func() {
// 			done <- true
// 		}()

// 		for index := range ticker.C {
// 			fmt.Println("Горутина 2:", index)
// 		}
// 	}()

// 	// Ожидаем завершения обеих горутин
// 	<-done
// 	<-done
// 	ticker.Stop()

// 	fmt.Println("end")
// }

// func counter() {
// 	fmt.Println("start")

// 	// Создаем каналы для синхронизации
// 	done := make(chan bool)

// 	// Горутина
// 	go func() {
// 		defer func() {
// 			fmt.Println("END")
// 			done <- true
// 		}()

// 		for i := 1; i <= 100; i++ {
// 			fmt.Println("Counter Goroutine:", i)
// 		}
// 	}()

// 	// Цикл
// 	for i := 1; i <= 100; i++ {
// 		fmt.Println("Counter Loop:", i)
// 	}

// 	// Ожидаем завершения горутины
// 	<-done

// 	fmt.Println("Both counters have completed")
// 	fmt.Println("end")
// }

// func secondCounter() {
// 	fmt.Println("start")

// 	// Создаем каналы для синхронизации
// 	done := make(chan bool)
// 	counterChan := make(chan int, 100)

// 	// Горутина
// 	go func() {
// 		defer func() {
// 			fmt.Println("END")
// 			done <- true
// 		}()

// 		for i := 1; i <= 100; i++ {
// 			counterChan <- i
// 		}
// 		close(counterChan) // Закрываем канал после отправки всех значений
// 	}()

// 	// Цикл
// 	for i := range counterChan {
// 		fmt.Println("Counter Loop:", i)
// 	}

// 	// Ожидаем завершения горутины
// 	<-done

// 	fmt.Println("Both counters have completed")
// 	fmt.Println("end")
// }
