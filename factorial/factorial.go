package main

import "fmt"

func factorial(val int) uint64 {
	table := make([]uint64, val+1)

	table[0] = 1
	for i := 1; i <= val; i++ {
		table[i] = table[i-1] * uint64(i)
	}

	return table[val]
}

func main() {
	fmt.Print("Enter Integer: ")
	var val int
	fmt.Scanf("%d", &val)
	result := factorial(val)
	fmt.Println("Result: ", result)
}
