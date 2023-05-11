package main

import (
	"fmt"
	"math"
)

func Rob(houses []int) int {
	return rob(houses, 0, make(map[int]int))
}

func rob(houses []int, current int, mem map[int]int) int {
	key := current
	if result, ok := mem[key]; ok {
		return result
	}

	if current >= len(houses) {
		mem[key] = 0
		return mem[key]
	}

	mem[key] = max(houses[current]+rob(houses, current+2, mem), rob(houses, current+1, mem))
	return mem[key]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(Rob([]int{2, 10, 3, 6, 8, 1, 7}))
}
