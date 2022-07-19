package animals

import "fmt"

type Cow struct {
	FoodEaten        string
	LocomotionMethod string
	SpokenSound      string
}

func NewCow() Cow {
	return Cow{
		FoodEaten:        "grass",
		LocomotionMethod: "walk",
		SpokenSound:      "moo",
	}
}

func (a Cow) Eat() {
	fmt.Println(a.FoodEaten)
}

func (a Cow) Move() {
	fmt.Println(a.LocomotionMethod)
}

func (a Cow) Speak() {
	fmt.Println(a.SpokenSound)
}
