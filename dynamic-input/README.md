# Dynamic User Input

Modify the code below to work as described:

> The program should receive an integer `n` as **user input** and store `n` integers in an array. It must, then, print the sum of all integers stored in the array

```go
func main() {
​
	var inputCount
	fmt.Scanf("%d", &inputCount)
​
	var v int
	var sum [inputCount]int
	for i := 0; i < inputCount; i++ {
		fmt.Scanf("%d", &v)
		sum[i] = v
	}
​
	total := 0
	for _, value := range sum {
		total += value
	}
​
	fmt.Println("Sum of array is ", total)
}
```