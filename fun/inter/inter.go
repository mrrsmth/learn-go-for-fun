package inter

import (
	"fmt"
	"time"
)

func Inter() {
	fmt.Println("start")

	// Создаем каналы для синхронизации
	done := make(chan bool)

	// Горутина
	go func() {
		defer func() {
			done <- true
		}()

		for i := 1; i <= 100; i++ {
			fmt.Println("Counter Goroutine:", i)
		}
	}()

	// Цикл
	for i := 1; i <= 100; i++ {
		fmt.Println("Counter Loop:", i)
	}

	// Ожидаем завершения горутины
	<-done

	fmt.Println("Both counters have completed")
	fmt.Println("end")
}

func ticker() {
	fmt.Println("start")

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	// Горутина 1
	go func() {
		defer func() {
			done <- true
		}()

		for index := range ticker.C {
			fmt.Println("Горутина 1:", index)
		}
	}()

	// Горутина 2
	go func() {
		defer func() {
			done <- true
		}()

		for index := range ticker.C {
			fmt.Println("Горутина 2:", index)
		}
	}()

	// Ожидаем завершения обеих горутин
	<-done
	<-done
	ticker.Stop()

	fmt.Println("end")
}
