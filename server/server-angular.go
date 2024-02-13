package serverangular

import (
	"fmt"
	"log"
	"net/http"

)

func Info() {
	staticFilesDir := "D:/git/learn-go/server/dist"

	fileServer := http.FileServer(http.Dir(staticFilesDir))

	http.Handle("/", http.StripPrefix("/", fileServer))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is a test endpoint")
	})

	fmt.Println("Сервер запущен и слушает порт 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
