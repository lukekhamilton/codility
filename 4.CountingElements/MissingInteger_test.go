package CountingElements

import "testing"

func TestMissingInteger(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{[]int{1, 3, 6, 4, 1, 2}}, 5},
		{"example1", args{[]int{1, 2, 3}}, 4},
		{"example2", args{[]int{-1, -3}}, 1},
		{"example3", args{[]int{-1, -3, -2, 0, 2}}, 1},
		{"example4", args{[]int{1, 3, 6, 4, 1, 2, 5, 8, 7, 9, 9, 4}}, 10},
		{"example5", args{[]int{-1, -3, 2}}, 1},
		{"example6", args{[]int{-1, -3, 1}}, 2},
		{"example7", args{[]int{0}}, 1},
		{"example8", args{[]int{-10000000}}, 1},
		{"example9", args{[]int{-10000000, 1}}, 2},
		{"example10", args{[]int{10000000}}, 1},
		{"example11", args{[]int{999999, 999998, 1000000}}, 1},
		{"example12", args{[]int{1, 3, 999999, 999998, 1000000}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingInteger(tt.args.A); got != tt.want {
				t.Errorf("MissingInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
