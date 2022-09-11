package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, error := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Value of the number: ", aFloat)
	}
}
