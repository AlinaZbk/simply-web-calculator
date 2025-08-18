package service

import "github.com/AlinaZbk/simply-web-calculator.git/model"

func Calculate(req model.CalculatorRequest) model.CalculatorResponse {
	var res float64

	switch req.Op {
	case "add":
		res = req.A + req.B
	case "sub":
		res = req.A - req.B
	case "mul":
		res = req.A * req.B
	case "div":
		if req.B == 0 {
			return model.CalculatorResponse{Error: "Деление на ноль запрещено!"}
		}
		res = req.A / req.B
	default:
		return model.CalculatorResponse{Error: "Неизвестная операция"}
	}

	return model.CalculatorResponse{Result: &res}
}
