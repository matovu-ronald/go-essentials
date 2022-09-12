package main

import "fmt"

func main() {
	numberOfPupils := -1

	var result string

	if numberOfPupils < 0 {
		result = "Less than zero"
	} else if numberOfPupils > 0 {
		result = "Greather than zero"
	} else {
		result = "Pupil cannot be zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
}
