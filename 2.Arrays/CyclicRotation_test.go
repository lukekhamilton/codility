package Arrays

import (
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example0", args{A: []int{}, K: 1}, []int{}},
		{"example1", args{A: []int{3, 8, 9, 7, 6}, K: 1}, []int{6, 3, 8, 9, 7}},
		{"example2", args{A: []int{1, 3, 8, 9, 7, 6}, K: 2}, []int{7, 6, 1, 3, 8, 9}},
		{"example3", args{A: []int{1, 3, 8, 9, 7, 6, 8, 1, 2, 5, 7, 8}, K: 3}, []int{5, 7, 8, 1, 3, 8, 9, 7, 6, 8, 1, 2}},
		{"example4", args{A: []int{0, 0, 0}, K: 1}, []int{0, 0, 0}},
		{"example5", args{A: []int{1, 3, 8, 9, 7, 6, 8, 1, 2, 5, 7, 8}, K: 15}, []int{5, 7, 8, 1, 3, 8, 9, 7, 6, 8, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CyclicRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
