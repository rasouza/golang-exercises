package main

import "fmt"

func main() {
	var inputCount int
	var result int = 0
	fmt.Print("Enter size: ")
	fmt.Scanf("%d", &inputCount)

	var v int
	sum := make([]int, inputCount)

	for j := 0; j < inputCount; j++ {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &v)
		sum[j] = v
	}

	for i := 0; i < len(sum); i++ {
		result += sum[i]
	}

	fmt.Println("Sum:", result)
}
