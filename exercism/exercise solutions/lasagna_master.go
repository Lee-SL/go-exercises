package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	totalNoodles := 0
	totalSauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			totalNoodles += 50
		} else if layer == "sauce" {
			totalSauce += 0.2
		}
	}

	return totalNoodles, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, yourList []string) {
	yourList[len(yourList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amount int) []float64 {
	factor := float64(amount) / 2

	var q []float64
	q = append(q, quantities...)

	for i, _ := range q {
		q[i] *= factor
	}

	return q
}
