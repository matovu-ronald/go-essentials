package main

import (
	"fmt"
	"math"
)

func main() {
	int1, int2, int3 := 12, 45, 97
	intSum := int1 + int2 + int3
	fmt.Println("Integer sum: ", intSum)

	float1, float2, float3 := 23.5, 65.1, 76.3
	floatSum := float1 + float2 + float3
	fmt.Println("Float sum: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now", floatSum)
}
