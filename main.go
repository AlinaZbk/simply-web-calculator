package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlinaZbk/simply-web-calculator.git/handler"
)

func main() {
	http.HandleFunc("/calc", handler.CalculatorHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
