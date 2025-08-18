package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlinaZbk/simply-web-calculator.git/handler"
)

func main() {
	http.HandleFunc("/calc", handler.CalculatorHandler)

	fmt.Println("Server running at http://localhost:8080/test")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
