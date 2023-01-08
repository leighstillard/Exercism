package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.00

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(newRecipe []string, ingredients []string) {
	//cut last item
	ingredients = ingredients[:len(ingredients)-1]

	//
	newRecipe = newRecipe[len(newRecipe)-1:]
	_ = append(ingredients, newRecipe...)

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(ingredients []float64, portions int) []float64 {
	scale := float64(portions) / 2.0
	output := make([]float64, len(ingredients))

	for i := 0; i < len(output); i++ {
		output[i] = ingredients[i] * scale
	}

	return output
}
