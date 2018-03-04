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
	if contains(natureAdjectives, name.Adjective) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(natureNouns, name.Noun) == false {
		t.Error("Expected to get valid noun")
	}
}

func TestAnimal(t *testing.T) {
	name := Animal()
	if contains(animalAdjectives, name.Adjective) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(animalNouns, name.Noun) == false {
		t.Error("Expected to get valid noun")
	}
}

func TestAnimalFromSeed(t *testing.T) {
	name := AnimalFromStringSeed("foo")
	if name.Adjective != "wonderful" || name.Noun != "magpie" {
		t.Error("Expected seed 1 to give wonderful magpie as name from seed foo")
	}
}

func TestNatureFromSeed(t *testing.T) {
	name := NatureFromStringSeed("foo")
	if name.Adjective != "rough" || name.Noun != "spirit" {
		t.Error("Expected seed 1 to give rough spirit as name from seed foo")
	}
}

func TestSeed(t *testing.T) {
	source := rand.NewSource(1)
	generator := rand.New(source)
	name := AnimalFromGenerator(generator)
	if name.Adjective != "blushing" || name.Noun != "crow" {
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

func TestMatches(t *testing.T) {
	name1 := Name{"blushing", "meerkat"}
	name2 := Name{"blushing", "meerkat"}
	if nameMatches(name1, name2) != true {
		t.Error("Expected names to match")
	}
}
