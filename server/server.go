package server

import (
	"fmt"
	"net/http"
)

type Msg string

func (m Msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func Server() {
	msgHandler := Msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", msgHandler)
}
