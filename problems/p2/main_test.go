package main

import "testing"

func Test_solution(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		S    string
		want int
	}{
		{
			name: "example from problem statement",
			A:    []int{-1, 0, 0, 0, 2},
			S:    "abbab",
			want: 3,
		},
		{
			name: "perfect alternating chain",
			A:    []int{-1, 0, 1, 2},
			S:    "abab",
			want: 4,
		},
		{
			name: "all same letter",
			A:    []int{-1, 0, 0},
			S:    "aaa",
			want: 1,
		},
		{
			name: "star with alternating children",
			A:    []int{-1, 0, 0, 0, 0},
			S:    "abbbb",
			want: 3,
		},
		{
			name: "parent index larger than child index",
			A:    []int{-1, 2, 0, 1},
			S:    "abab",
			want: 2,
		},
		{
			name: "single node",
			A:    []int{-1},
			S:    "a",
			want: 1,
		},
		{
			name: "two nodes same letter",
			A:    []int{-1, 0},
			S:    "aa",
			want: 1,
		},
		{
			name: "two nodes different letter",
			A:    []int{-1, 0},
			S:    "ab",
			want: 2,
		},
		{
			name: "long alternating path through root",
			A:    []int{-1, 0, 1, 2, 0, 4, 5},
			S:    "ababbab",
			want: 7,
		},
		{
			name: "path breaks at the end",
			A:    []int{-1, 0, 1, 2, 3},
			S:    "ababb",
			want: 4,
		},
		{
			name: "deep tree all alternating",
			A:    []int{-1, 0, 1, 2, 3, 4},
			S:    "ababab",
			want: 6,
		},
		{
			name: "wide tree best path through root",
			A:    []int{-1, 0, 0, 1, 1, 2, 2},
			S:    "abbabab",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution(tt.A, tt.S)
			if got != tt.want {
				t.Errorf("solution(%v, %q) = %d, want %d", tt.A, tt.S, got, tt.want)
			}
		})
	}
}
