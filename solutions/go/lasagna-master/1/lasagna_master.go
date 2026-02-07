package lasagna

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

func Quantities(layers []string) (noodleAmount int, sauceAmount float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauceAmount += 0.2
		}
		if layers[i] == "noodles" {
			noodleAmount += 50
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(amountForTwoPortions []float64, numberOfPortionsToCook int) []float64 {
	result := make([]float64, len(amountForTwoPortions))
	for i := 0; i < len(amountForTwoPortions); i++ {
		result[i] = (amountForTwoPortions[i] / 2) * float64(numberOfPortionsToCook)
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
