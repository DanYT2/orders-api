package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler : http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on port 3000")
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}