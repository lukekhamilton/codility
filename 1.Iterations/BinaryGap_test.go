package Iterations

import "testing"

func TestBinaryGap(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ex0", args{9}, 2},
		{"ex1", args{529}, 4},
		{"ex2", args{20}, 1},
		{"ex3", args{15}, 0},
		{"ex4", args{32}, 0},
		{"ex5", args{1041}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryGap(tt.args.N); got != tt.want {
				t.Errorf("BinaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
