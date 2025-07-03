package main

import . "fmt"

type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(Name) Dog
)

// The reason we chose int for Breed and Gender is we will construct those using Go's equivalent of Enum definition

const (
	BullDog Breed = iota
	Havanese
	Cavalier
	Poodle
	GoldenRetriever
)

// define possible genders
const (
	Male Gender = iota
	Female
)

// This shows that type aliases compile down to the underlying type, in this case, the int type for which iota is defined

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

func createDogsWithoutPartialApplication() []Dog {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}
	rocky := Dog{
		Name:   "Rocky",
		Breed:  GoldenRetriever,
		Gender: Male,
	}
	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}

	return []Dog{bucky, rocky, tipsy}
}

// DogSpawner function allows us to create new functions where the dog’s breed and gender are already partially applied,
// but the name is still expected as input.
func DogSpawner(breed Breed, gender Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  breed,
			Gender: gender,
			Name:   n,
		}
	}
}

// Using the DogSpawner function, we can create two new functions, maleHavaneseSpawner and femalePoodleSpawner. These
// functions will allow us to create male Havanese dogs and female poodles, by only providing a name for our dogs.
var (
	maleHaveneseSpawner        = DogSpawner(Havanese, Male)
	femalePoodleSpawner        = DogSpawner(Poodle, Female)
	maleGoldenRetrieverSpawner = DogSpawner(GoldenRetriever, Male)
)

// After this definition, the maleHavaneseSpawner and femalePoodleSpawner functions are available anywhere in that package

func main() {
	bucky := maleHaveneseSpawner("Bucky")
	rocky := maleGoldenRetrieverSpawner("Rocky")
	tipsy := femalePoodleSpawner("Tipsy")
	dogs := []Dog{bucky, rocky, tipsy}
	Printf("Dogs = %v\n", dogs)
	dogsViaTraditionalFunc := createDogsWithoutPartialApplication()
	Printf("Dogs (traditional) = %v\n", dogsViaTraditionalFunc)
}
