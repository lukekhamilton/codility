package CountingElements

import "testing"

func TestPermCheck(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{[]int{4, 1, 3, 2}}, 1},
		{"example1", args{[]int{4, 1, 3}}, 0},
		{"example2", args{[]int{}}, 0},
		{"example3", args{[]int{-1, 0, 1, 2}}, 0},
		{"example4", args{[]int{5, 0, 1, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermCheck(tt.args.A); got != tt.want {
				t.Errorf("PermCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
