package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Println(">")

		buffer := bufio.NewReader(os.Stdin)
		nextLine, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Printf("reading input, %v\n", err)
			continue
		}

		info := strings.Split(strings.TrimSuffix(nextLine, "\n"), " ")
		if len(info) != 2 {
			fmt.Println("incorrect arguments")
			continue
		}

		animal, err := getAnimal(info[0])
		if err != nil {
			fmt.Printf("getting animal, %v\n", err)
			continue
		}

		if err = animal.doAction(info[1]); err != nil {
			fmt.Printf("doing animal action, %v\n", err)
		}
	}

}

var (
	cow = Animal{
		FoodEaten:        "grass",
		LocomotionMethod: "walk",
		SpokenSound:      "moo",
	}

	bird = Animal{
		FoodEaten:        "worms",
		LocomotionMethod: "fly",
		SpokenSound:      "peep",
	}

	snake = Animal{
		FoodEaten:        "mice",
		LocomotionMethod: "slither",
		SpokenSound:      "hsss",
	}
)

type Animal struct {
	FoodEaten        string
	LocomotionMethod string
	SpokenSound      string
}

func (a Animal) Eat() {
	fmt.Println(a.FoodEaten)
}

func (a Animal) Move() {
	fmt.Println(a.LocomotionMethod)
}

func (a Animal) Speak() {
	fmt.Println(a.SpokenSound)
}

func getAnimal(animalString string) (Animal, error) {
	switch animalString {
	case "cow":
		return cow, nil
	case "bird":
		return bird, nil
	case "snake":
		return snake, nil
	default:
		return Animal{}, errors.New("unknown animal")
	}
}

func (a Animal) doAction(action string) error {
	switch action {
	case "eat":
		a.Eat()
		return nil
	case "move":
		a.Move()
		return nil
	case "speak":
		a.Speak()
		return nil
	default:
		return errors.New("unknown action")
	}
}
