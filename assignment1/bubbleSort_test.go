package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		testName       string
		input          *[]int
		expectedResult *[]int
	}{
		{
			testName:       "case 1: all positive numbers",
			input:          &[]int{9, 4, 7, 1, 6, 3, 1, 1, 1, 2},
			expectedResult: &[]int{1, 1, 1, 1, 2, 3, 4, 6, 7, 9},
		},
		{
			testName:       "case 1: one negative number",
			input:          &[]int{4, 7, 1, 6, -9, 3, 1, 1, 1, 2},
			expectedResult: &[]int{-9, 1, 1, 1, 1, 2, 3, 4, 6, 7},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			BubbleSort(test.input)

			assert.Equal(t, test.input, test.expectedResult, "actual and expected should be equal")
		})
	}
}
