package PrefixSums

import "testing"

func TestPassingCars(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{0, 1, 0, 1, 1}}, 5},
		{"example2", args{[]int{0, 0, 0, 0, 0}}, 0},
		{"example3", args{[]int{1, 1, 1, 1, 1}}, 0},
		{"example4", args{[]int{0, 0, 0, 0, 1}}, 4},
		{"example5", args{[]int{0, 0, 0, 1, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PassingCars(tt.args.A); got != tt.want {
				t.Errorf("PassingCars() = %v, want %v", got, tt.want)
			}
		})
	}
}
