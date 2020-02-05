package PrefixSums

import "testing"

func TestMinAvgTwoSlice(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{[]int{4, 2, 2, 5, 1, 5, 8}}, 1},
		{"example1", args{[]int{2, 2}}, 0},
		{"example2", args{[]int{5, 1}}, 0},
		{"example3", args{[]int{5, 1, 2, 2}}, 1},
		{"example4", args{[]int{9, 9, 1, 1, 9}}, 2},
		{"example5", args{[]int{-9, -9, -1, -1, -9}}, 0},
		{"example6", args{[]int{-1, -1, -9, -9, -1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinAvgTwoSlice(tt.args.A); got != tt.want {
				t.Errorf("MinAvgTwoSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
