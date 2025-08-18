package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CalculatorRequest struct {
	A  float64 `json:"a"`
	B  float64 `json:"b"`
	Op string  `json:"op"` //add, sub, mul, div
}

type CalculatorResponse struct {
	Result float64 `jsone:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	// эндпоинт /calc - универсальный калькулятор
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CalculatorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid Json", http.StatusBadRequest)
			return
		}

		var result float64
		var errorMsg string

		switch req.Op {
		case "add": // сложение
			result = req.A + req.B
		case "sub": // вычитание
			result = req.A - req.B
		case "mul": // умножение
			result = req.A * req.B
		case "div": // деление
			if req.B == 0 {
				errorMsg = "Деление на 0 запрещено"
			} else {
				result = req.A / req.B
			}
		default:
			errorMsg = "Операции не существует"
		}

		resp := CalculatorResponse{Result: result, Error: errorMsg}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// эндпоинт /test - передаем в html
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("index.html")
		if err != nil {
			http.Error(w, "file not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	fmt.Println("Server running at http://localhost:8080/test")
	http.ListenAndServe(":8080", nil)
}
