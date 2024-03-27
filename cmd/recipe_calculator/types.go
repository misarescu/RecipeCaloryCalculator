package main

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Recipe struct {
	Name        string       `json:"recipe-name"`
	Ingredients []Ingredient `json:"ingredients"`
}
