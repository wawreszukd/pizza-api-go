package models

type Pizza struct {
	ID      int32   `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Topping string  `json:"topping"`
}
