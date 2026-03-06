package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    /*a, b := 1, 20
	return a + rand.Intn(b-a+1) // a ≤ n ≤ b
    */
    return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	//return 11.9 * rand.Float64()
    return 12 * rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    animals := []string{ "ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	//words := strings.Fields("ant beaver cat dog elephant fox giraffe hedgehog")
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
    return animals
}
