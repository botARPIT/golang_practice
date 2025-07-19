package main

import "fmt"

func main() {
	fmt.Println(findMaxNumber([]int{50, 60, 80, 10, 99, 200}))
}

func findMaxNumber(numbers []int) int {
	var max int = 0
	if len(numbers) == 0 {
		return max
	}
	for i := 0; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	return max

}
