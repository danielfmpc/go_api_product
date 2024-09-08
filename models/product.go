package models

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}
