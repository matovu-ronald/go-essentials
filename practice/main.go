package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["CU"] = "Central Uganda"
	states["WNU"] = "West Nile Uganda"
	states["EU"] = "Eastern Uganda"
	states["NU"] = "Northern Uganda"
	states["WU"] = "Western Uganda"

	fmt.Println(states)
	fmt.Println(states["WNU"])

	delete(states, "NU")
	fmt.Println(states)

	for key, value := range states {
		fmt.Printf("The full form of %v is %v\n", key, value)
	}

	keys := make([]string, len(states))

	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
