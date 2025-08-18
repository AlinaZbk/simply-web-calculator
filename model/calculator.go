package model

type CalculatorRequest struct {
	A  float64 `json:"a"`
	B  float64 `json:"b"`
	Op string  `json:"op"` //add, sub, mul, div
}

type CalculatorResponse struct {
	Result *float64 `jsone:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}