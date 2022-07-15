package main

import "fmt"

func main() {

	// Read values from input
	var acceleration, velocity_0, displacement_0, time float64

	fmt.Println("Please enter the value for acceleration: ")
	if _, err := fmt.Scanf("%f", &acceleration); err != nil {
		fmt.Printf("reading input, %v\n", err)
		return
	}

	fmt.Println("Please enter the value for initial velocity: ")
	if _, err := fmt.Scanf("%f", &velocity_0); err != nil {
		fmt.Printf("reading input, %v\n", err)
		return
	}

	fmt.Println("Please enter the value for initial displacement: ")
	if _, err := fmt.Scanf("%f", &displacement_0); err != nil {
		fmt.Printf("reading input, %v\n", err)
		return
	}

	dispFn := GenDisplaceFn(acceleration, velocity_0, displacement_0)

	fmt.Println("Please enter the value for time: ")
	if _, err := fmt.Scanf("%f", &time); err != nil {
		fmt.Printf("reading input, %v\n", err)
		return
	}

	fmt.Println(dispFn(time))
}

func GenDisplaceFn(acceleration, velocity_0, displacement_0 float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * time * time) + (velocity_0 * time) + displacement_0
	}
}
