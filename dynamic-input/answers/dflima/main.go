package main

import "fmt"

func main() {
	var inputCount int
	fmt.Scanf("%d", &inputCount)

	var v int
	sum := make([]int, inputCount)
	for i := 0; i < inputCount; i++ {
		fmt.Scanf("%d", &v)
		sum[i] = v
	}

	total := 0
	for _, value := range sum {
		total += value
	}

	fmt.Printf("Sum of array is %d", total)
}
