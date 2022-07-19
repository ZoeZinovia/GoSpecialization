package main

import (
	"Personal/GoSpecialization/assignment4/animals"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var userAnimals = make(map[string]animals.Animal)

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
		if len(info) != 3 {
			fmt.Println("incorrect arguments")
			continue
		}

		command := info[0]
		name := info[1]
		animalOrAction := info[2]

		switch command {
		case "newanimal":
			if err := newAnimalRequest(name, animalOrAction); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Created it!")
		case "query":
			if err := queryAnimal(name, animalOrAction); err != nil {
				fmt.Println(err)
				continue
			}
		default:
			fmt.Println("unknown command")
			return
		}
	}
}

func newAnimalRequest(name, animal string) error {
	switch animal {
	case "cow":
		userAnimals[name] = animals.NewCow()
	case "bird":
		userAnimals[name] = animals.NewBird()
	case "snake":
		userAnimals[name] = animals.NewSnake()
	default:
		return errors.New("unknown animal")
	}
	return nil
}

func queryAnimal(name, action string) error {
	animal, ok := userAnimals[name]
	if !ok {
		return errors.New("animal name not found")
	}
	switch action {
	case "eat":
		animal.Eat()
	case "speak":
		animal.Speak()
	case "move":
		animal.Move()
	default:
		return errors.New("unknown action")
	}
	return nil
}
