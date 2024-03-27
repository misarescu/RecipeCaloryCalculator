package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"recipe-calculator.com/internal/server"
)

func GetRecipeValue(w http.ResponseWriter, r *http.Request) error {
	recipe := Recipe{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&recipe); err != nil {
		return fmt.Errorf("error unmarshalling json from body: %s", err.Error())
	}

	resIngredients := []IngredientCaloriesResponse{}

	for _, ing := range recipe.Ingredients {
		resIngredients = append(resIngredients, IngredientCaloriesResponse{Name: ing.Name, KJoules: 2409.98, KCals: 576})
	}

	res := RecipeCaloriesResponse{Name: recipe.Name, KCals: 1134, KJoules: 4744.656, Ingredients: resIngredients}

	return server.WriteJSON(w, http.StatusOK, res)
}
