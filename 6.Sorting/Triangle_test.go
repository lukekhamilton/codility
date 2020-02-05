package Sorting

import "testing"

func TestTriangle(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{10, 2, 5, 1, 8, 20}}, 1},
		{"example2", args{[]int{10, 50, 5, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Triangle(tt.args.A); got != tt.want {
				t.Errorf("Triangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
