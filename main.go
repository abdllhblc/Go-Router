package main

import (
	"module/server"
	"net/http"
)

func main() {
	router := server.NewServer()

	router.Get("/",helloHandler)
	router.RunServer(":8000",nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
