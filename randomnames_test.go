package randomnames

import "testing"

func contains(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}
	return false
}

func TestNature(t *testing.T) {
	adj, noun := Nature()
	if contains(natureAdjectives, adj) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(natureNouns, noun) == false {
		t.Error("Expected to get valid noun")
	}
}

func TestAnimal(t *testing.T) {
	adj, noun := Animal()
	if contains(animalAdjectives, adj) == false {
		t.Error("Expected to get valid adjective")
	}
	if contains(animalNouns, noun) == false {
		t.Error("Expected to get valid noun")
	}
}
