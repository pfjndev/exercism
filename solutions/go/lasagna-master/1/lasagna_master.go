package lasagnamaster

import "strings"

// PreparationTime calculates the total prep time based on the number of layers.
func PreparationTime(layers []string, avgPrepTime int) (int) {
    if avgPrepTime <= 0 { avgPrepTime = 2 }
    return len(layers) * avgPrepTime
}
// Quantities calculates the amounts of noodles and sauce needed.
func Quantities(layers []string) (int, float64) {
    noodlesQnt, sauceQnt := 0, 0.0
    for _, v := range layers {
        switch {
            case strings.EqualFold(v, "noodles"): noodlesQnt += 50
            case strings.EqualFold(v, "sauce"): sauceQnt += 0.2
        }
    }
    return noodlesQnt, sauceQnt
}
// AddSecretIngredient replaces the last item in your list with the last from your friend's.
func AddSecretIngredient(list1, list2 []string) {
    list2[len(list2)-1] = list1[len(list1)-1]
}
// ScaleRecipe returns the amounts needed for a different number of portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))
	// Convert portions to float64 first to avoid integer division
	factor := float64(portions) / 2.0

	for key := range amounts {
		scaled[key] = amounts[key] * factor
	}
	return scaled
}
