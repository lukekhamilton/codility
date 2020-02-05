package TimeComplexity

import "testing"

func TestTapeEquilibrium(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ex0", args{[]int{3, 1, 2, 4, 3}}, 1},
		{"ex1", args{[]int{3, 1, 2, 4, 3, 1, 2, 3, 10}}, 1},
		{"ex2", args{[]int{2, 3, 10}}, 5},
		{"ex3", args{[]int{0, 1}}, 1},
		{"ex4", args{[]int{1, 1}}, 0},
		{"ex5", args{[]int{-3, 1, 2, -4, 3}}, 1},
		{"ex6", args{[]int{-1000, 1000, -500, 990}}, 490},
		{"ex7", args{[]int{1, 2}}, 1},
		{"ex8", args{[]int{100, -25}}, 125},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TapeEquilibrium(tt.args.A); got != tt.want {
				t.Errorf("TapeEquilibrium() = %v, want %v", got, tt.want)
			}
		})
	}
}
