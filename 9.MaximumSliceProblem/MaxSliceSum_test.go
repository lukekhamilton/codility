package MaximumSliceProblem

import "testing"

func TestMaxSliceSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ex0", args{[]int{3, 2, -6, 4, 0}}, 5},
		{"ex1", args{[]int{1, -3, 2, -5, 7, 6, -1, -4, 11, -23}}, 19},
		{"ex2", args{[]int{-2, -3, -6, -12, -1, -52}}, -1},
		{"ex3", args{[]int{2, 3, 6, 12, 1, 52}}, 76},
		{"ex4", args{[]int{1, -3, 2, 1, -1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSliceSum(tt.args.A); got != tt.want {
				t.Errorf("MaxSliceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
