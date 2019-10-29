package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := process()
	fmt.Println(strconv.FormatBool(b))
}

func process() bool {
	var k, v, inputCount int
	fmt.Scanf("%d", &k)

	fmt.Scanf("%d", &inputCount)
	list := make([]int, inputCount)

	for a := 0; a < inputCount; a++ {
		fmt.Scanf("%d", &v)
		list[a] = v
	}

	for i := 0; i < inputCount; i++ {
		for j := 0; j < inputCount; j++ {
			if i == j {
				continue
			}

			if list[i]+list[j] == k {
				return true
			}
		}
	}

	return false
}
