package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AlinaZbk/simply-web-calculator.git/model"
	"github.com/AlinaZbk/simply-web-calculator.git/service"
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.CalculatorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	resp := service.Calculate(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
