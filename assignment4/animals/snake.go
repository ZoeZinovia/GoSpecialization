package animals

import "fmt"

type Snake struct {
	FoodEaten        string
	LocomotionMethod string
	SpokenSound      string
}

func NewSnake() Snake {
	return Snake{
		FoodEaten:        "mice",
		LocomotionMethod: "slither",
		SpokenSound:      "hsss",
	}
}

func (a Snake) Eat() {
	fmt.Println(a.FoodEaten)
}

func (a Snake) Move() {
	fmt.Println(a.LocomotionMethod)
}

func (a Snake) Speak() {
	fmt.Println(a.SpokenSound)
}
