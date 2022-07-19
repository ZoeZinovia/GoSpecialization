package animals

import "fmt"

type Bird struct {
	FoodEaten        string
	LocomotionMethod string
	SpokenSound      string
}

func NewBird() Bird {
	return Bird{
		FoodEaten:        "worms",
		LocomotionMethod: "fly",
		SpokenSound:      "peep",
	}
}

func (a Bird) Eat() {
	fmt.Println(a.FoodEaten)
}

func (a Bird) Move() {
	fmt.Println(a.LocomotionMethod)
}

func (a Bird) Speak() {
	fmt.Println(a.SpokenSound)
}
