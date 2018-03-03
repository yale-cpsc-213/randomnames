package randomnames

import (
	"math/rand"
	"time"
)

var globalGenerator *rand.Rand

func init() {
	source := rand.NewSource(time.Now().Unix())
	globalGenerator = rand.New(source)
}

// NatureFromGenerator returns a nature name
func NatureFromGenerator(generator *rand.Rand) (string, string) {
	i1 := generator.Intn(len(natureAdjectives))
	i2 := generator.Intn(len(natureNouns))
	return natureAdjectives[i1], natureNouns[i2]
}

// Nature returns a nature name and uses the package-scoped
// random number generator.
func Nature() (string, string) {
	return NatureFromGenerator(globalGenerator)
}

// AnimalFromGenerator returns a Animal name
func AnimalFromGenerator(generator *rand.Rand) (string, string) {
	i1 := generator.Intn(len(animalAdjectives))
	i2 := generator.Intn(len(animalNouns))
	return animalAdjectives[i1], animalNouns[i2]
}

// Animal returns a Animal name and uses the package-scoped
// random number generator.
func Animal() (string, string) {
	return AnimalFromGenerator(globalGenerator)
}
