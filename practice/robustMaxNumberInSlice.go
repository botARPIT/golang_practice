package main

import (
	"errors"
	"fmt"
)

func maxNumberInSlice(numbers []int) (int, error) {
	var max int = 0
	if len(numbers) == 0 {
		return 0, errors.New("Empty slice of numbers")
	}
	max = numbers[0]
	for _, val := range numbers {
		if max < val {
			max = val
		}
	}
	return max, nil
}

func main() {
	fmt.Println(maxNumberInSlice([]int{10, 30, 40}))
}
