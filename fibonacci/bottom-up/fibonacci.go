package main

import "fmt"

func Fibonacci(n int) int {
	table := make([]int, n+1)
	table[1] = 1
	for i := 2; i < len(table); i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

func main() {
	fmt.Println(Fibonacci(10))
}
