package main

import "fmt"

func main() {
	anInteger := 26

	var p = &anInteger
	fmt.Println("Value of p: ", *p)

	value1 := 24.13
	pointer1 := &value1

	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = *pointer1 * 2
	fmt.Println("Pointer 1: ", *pointer1)
	fmt.Println("Value 1: ", value1)
}
