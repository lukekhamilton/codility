package Sorting

import "testing"

func TestDistinct(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{[]int{2, 1, 1, 2, 3, 1}}, 3},
		{"example1", args{[]int{1}}, 1},
		{"example2", args{[]int{}}, 0},
		{"example3", args{[]int{-1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.A); got != tt.want {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
