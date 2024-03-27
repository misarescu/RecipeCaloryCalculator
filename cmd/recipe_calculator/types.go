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

type IngredientCaloriesResponse struct {
	Name    string  `json:"name"`
	KJoules float64 `json:"kJoules"`
	KCals   float64 `json:"kCals"`
}

type RecipeCaloriesResponse struct {
	Name        string                       `json:"recipe-name"`
	KJoules     float64                      `json:"total-kJoules"`
	KCals       float64                      `json:"total-kCals"`
	Ingredients []IngredientCaloriesResponse `json:"ingredients"`
}
