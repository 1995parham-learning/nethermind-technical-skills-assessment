package main

import "fmt"

// solution returns the number of vertices on the longest alternating path in the tree.
//
// Approach:
// - For each node, compute `down[i]`: the longest alternating path (in vertex count)
//   starting at node i going into its subtree.
// - For each node, combine the two best extending children to form a path going
//   through node i (child1 ↔ i ↔ child2). Track the global maximum.
// - BFS from root to get processing order, then process in reverse (children before parents).
func solution(A []int, S string) int {
	n := len(A)

	// Build children list.
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[A[i]] = append(children[A[i]], i)
	}

	// BFS from root to determine processing order.
	order := make([]int, 0, n)
	queue := []int{0}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		queue = append(queue, children[u]...)
	}

	down := make([]int, n)
	result := 1

	// Process in reverse BFS order: children before parents.
	for idx := len(order) - 1; idx >= 0; idx-- {
		i := order[idx]
		down[i] = 1
		top1, top2 := 0, 0 // two longest extending child paths

		for _, c := range children[i] {
			if S[c] != S[i] {
				// Child has different letter — we can extend through it.
				if down[c] > top1 {
					top2 = top1
					top1 = down[c]
				} else if down[c] > top2 {
					top2 = down[c]
				}
				if 1+down[c] > down[i] {
					down[i] = 1 + down[c]
				}
			}
		}

		// Path through node i using its two best alternating branches.
		pathThrough := top1 + top2 + 1
		if pathThrough > result {
			result = pathThrough
		}
	}

	return result
}

func main() {
	// Example from problem: S="abbab", A=[-1,0,0,0,2]
	// Tree: 0(a)-1(b), 0(a)-2(b), 0(a)-3(a), 2(b)-4(b)
	// Longest alternating path: 1(b) -> 0(a) -> 2(b) = 3 vertices
	fmt.Println(solution([]int{-1, 0, 0, 0, 2}, "abbab")) // 3

	// Simple chain: a-b-a-b → path of 4
	fmt.Println(solution([]int{-1, 0, 1, 2}, "abab")) // 4

	// All same letter: each node alone → 1
	fmt.Println(solution([]int{-1, 0, 0}, "aaa")) // 1

	// Star: center 'a', all children 'b' → best path b-a-b = 3
	fmt.Println(solution([]int{-1, 0, 0, 0, 0}, "abbbb")) // 3

	// Failing case: parent index can be larger than child index
	// Tree: 0(a)-2(a)-1(b)-3(b), longest alternating: 2(a)->1(b) = 2
	fmt.Println(solution([]int{-1, 2, 0, 1}, "abab")) // 2
}
