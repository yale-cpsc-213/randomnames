package randomnames

import (
	"math/rand"
	"testing"
)

func contains(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}
	return false
}

func TestNature(t *testing.T) {
	name := Nature()
	if contains(natureAdjectives, name.adjective) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(natureNouns, name.noun) == false {
		t.Error("Expected to get valid noun")
	}
}

func TestAnimal(t *testing.T) {
	name := Animal()
	if contains(animalAdjectives, name.adjective) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(animalNouns, name.noun) == false {
		t.Error("Expected to get valid noun")
	}
}

func TestSeed(t *testing.T) {
	source := rand.NewSource(1)
	generator := rand.New(source)
	name := AnimalFromGenerator(generator)
	if name.adjective != "blushing" || name.noun != "crow" {
		t.Error("Expected seed 1 to give blushing cow as first name")
	}
}

func TestUnique(t *testing.T) {
	source := rand.NewSource(1)
	generator1 := rand.New(source)
	generator2 := rand.New(source)
	names := []Name{AnimalFromGenerator(generator1)}
	name := UniqueAnimalFromGenerator(generator2, names)
	if nameMatches(name, names[0]) {
		t.Error("Expected to get a unique name")
	}
}
