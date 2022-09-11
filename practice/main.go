package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Orange")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 23
	numbers[1] = 2
	numbers[2] = 42
	numbers[3] = 90
	numbers[4] = 19

	numbers = append(numbers, 33)
	numbers = append(numbers, 11)

	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
