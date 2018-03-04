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

// Name holds an adjective and noun
type Name struct {
	Adjective string
	Noun      string
}

// nameMatches tests if two names are the same
func nameMatches(name1 Name, name2 Name) bool {
	if name1.Adjective == name2.Adjective && name1.Noun == name2.Noun {
		return true
	}
	return false
}

// nameSliceContains tests if a slice of names contains
// a given name
func nameSliceContains(names []Name, name Name) bool {
	for _, name2 := range names {
		if nameMatches(name, name2) {
			return true
		}
	}
	return false
}

// NatureFromGenerator returns a nature name
func NatureFromGenerator(generator *rand.Rand) Name {
	i1 := generator.Intn(len(natureAdjectives))
	i2 := generator.Intn(len(natureNouns))
	return Name{natureAdjectives[i1], natureNouns[i2]}
}

// Nature returns a nature name and uses the package-scoped
// random number generator.
func Nature() Name {
	return NatureFromGenerator(globalGenerator)
}

// AnimalFromGenerator returns a Animal name
func AnimalFromGenerator(generator *rand.Rand) Name {
	i1 := generator.Intn(len(animalAdjectives))
	i2 := generator.Intn(len(animalNouns))
	return Name{animalAdjectives[i1], animalNouns[i2]}
}

// Animal returns a Animal name and uses the package-scoped
// random number generator.
func Animal() Name {
	return AnimalFromGenerator(globalGenerator)
}

// UniqueAnimalFromGenerator generates a name that is not in the
// provided slice of names.
func UniqueAnimalFromGenerator(generator *rand.Rand, names []Name) Name {
	name := AnimalFromGenerator(generator)
	for true {
		if nameSliceContains(names, name) == false {
			return name
		}
		name = AnimalFromGenerator(generator)
	}
	return name
}

// UniqueAnimal generates a name that is not in the
// provided slice of names.
func UniqueAnimal(names []Name) Name {
	return UniqueAnimalFromGenerator(globalGenerator, names)
}
