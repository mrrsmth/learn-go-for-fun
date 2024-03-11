package fun

import (
	"log"
	"net/http"
	"time"
)

// Middleware для логирования запросов
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[%s] %s %s %s", r.Method, r.RemoteAddr, r.URL.Path, duration)
	})
}

// Обработчик для маршрута /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func Main() {
	// Создание маршрутизатора
	mux := http.NewServeMux()

	// Подключение middleware Logger к маршрутизатору
	mux.Handle("/hello", Logger(http.HandlerFunc(helloHandler)))

	// Запуск сервера
	http.ListenAndServe(":8080", mux)
}
