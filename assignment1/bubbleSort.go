package main

import (
	"fmt"
)

func main() {
	// Get user name
	fmt.Println("Please enter your 10 integers: ")

	// ReadString will block until the delimiter is entered
	var inputSlice []int
	for i := 0; i < 10; i++ {
		var nextInt int
		if _, err := fmt.Scanf("%d", &nextInt); err != nil {
			fmt.Printf("reading input, %v\n", err)
			return
		}

		inputSlice = append(inputSlice, nextInt)
	}

	// Sort slice
	BubbleSort(inputSlice)

	// Print slice
	fmt.Println(intSliceToString(inputSlice))
}

func intSliceToString(inputSlice []int) (result string) {
	for _, i := range inputSlice {
		result += fmt.Sprintf("%d", i)
	}
	return
}

func BubbleSort(intSlice []int) {
	// Get slice length
	len := len(intSlice)

	// Sort slice
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if (intSlice)[j] > (intSlice)[j+1] {
				Swap(intSlice, j)
			}
		}
	}
}

func Swap(intSlice []int, index int) {
	// Make copy of dereferenced intSlice
	intSlice[index], intSlice[index+1] = intSlice[index+1], intSlice[index]
}
