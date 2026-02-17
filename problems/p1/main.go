package main

import (
	"fmt"
	"sort"
)

// solution returns the maximum number of different cars that can be produced.
//
// Approach:
// 1. Sort cars by production time (ascending) — to maximize count, prefer cheapest cars.
// 2. Incrementally add cars and maintain a subset-sum DP over the first assembly line's capacity.
// 3. For each count k, check if there's a way to split the k cheapest cars across
//    the two lines (sum on line1 ≤ X, remainder ≤ Y).
func solution(H []int, X int, Y int) int {
	n := len(H)
	sort.Ints(H)

	result := 0
	total := 0

	// dp[s] = true means we can achieve exactly sum s on line 1
	// using some subset of the cars considered so far.
	dp := make([]bool, X+1)
	dp[0] = true

	for k := 0; k < n; k++ {
		total += H[k]
		if total > X+Y {
			break
		}

		// 0-1 knapsack update: traverse in reverse to avoid reusing the same car.
		for s := X; s >= H[k]; s-- {
			if dp[s-H[k]] {
				dp[s] = true
			}
		}

		// Check if any valid partition exists.
		// We need some s where dp[s] && s <= X && (total - s) <= Y
		// i.e. s >= total - Y.
		lo := total - Y
		if lo < 0 {
			lo = 0
		}
		for s := X; s >= lo; s-- {
			if dp[s] {
				result = k + 1
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(solution([]int{2, 3, 4, 5}, 7, 8)) // 4
	fmt.Println(solution([]int{4, 4, 4}, 5, 5))     // 2
	fmt.Println(solution([]int{1, 1, 1, 1}, 2, 2))  // 4
}
