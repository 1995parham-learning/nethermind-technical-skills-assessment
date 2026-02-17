package main

import "testing"

func Test_solution(t *testing.T) {
	tests := []struct {
		name string
		H    []int
		X    int
		Y    int
		want int
	}{
		{
			name: "all cars fit across two lines",
			H:    []int{2, 3, 4, 5},
			X:    7,
			Y:    8,
			want: 4,
		},
		{
			name: "only two out of three fit",
			H:    []int{4, 4, 4},
			X:    5,
			Y:    5,
			want: 2,
		},
		{
			name: "all small cars fit",
			H:    []int{1, 1, 1, 1},
			X:    2,
			Y:    2,
			want: 4,
		},
		{
			name: "single car fits line 1",
			H:    []int{5},
			X:    5,
			Y:    3,
			want: 1,
		},
		{
			name: "single car fits line 2 only",
			H:    []int{5},
			X:    3,
			Y:    5,
			want: 1,
		},
		{
			name: "single car too large for both",
			H:    []int{10},
			X:    3,
			Y:    5,
			want: 0,
		},
		{
			name: "no cars",
			H:    []int{},
			X:    10,
			Y:    10,
			want: 0,
		},
		{
			name: "all cars identical fit perfectly",
			H:    []int{3, 3, 3, 3},
			X:    6,
			Y:    6,
			want: 4,
		},
		{
			name: "total exceeds combined capacity",
			H:    []int{5, 5, 5},
			X:    6,
			Y:    6,
			want: 2,
		},
		{
			name: "one line has zero capacity",
			H:    []int{1, 2, 3},
			X:    0,
			Y:    6,
			want: 3,
		},
		{
			name: "greedy split needed",
			H:    []int{1, 2, 3, 7},
			X:    3,
			Y:    7,
			want: 3,
		},
		{
			name: "large line capacities",
			H:    []int{10, 20, 30, 40, 50},
			X:    60,
			Y:    90,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution(tt.H, tt.X, tt.Y)
			if got != tt.want {
				t.Errorf("solution(%v, %d, %d) = %d, want %d", tt.H, tt.X, tt.Y, got, tt.want)
			}
		})
	}
}
