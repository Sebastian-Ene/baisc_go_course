package main

import "fmt"

func main() {
	ints := generateIntSlices(19)
	printOddOrEven(ints)
}

func printOddOrEven(i []int) {
	for _, num := range i {
		if num%2 == 0 {
			fmt.Printf("Num %v is even.\n", num)
		} else {
			fmt.Printf("Num %v is odd.\n", num)
		}
	}
}

func generateIntSlices(n int) []int {
	start := 1
	ints := []int{}
	for start < n {
		ints = append(ints, start)
		start += 1
	}
	return ints
}
