package Arrays

import "testing"

func TestOddOccurrencesInArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{[]int{7, 9, 3, 9, 3, 9, 9}}, 7},
		{"example0", args{[]int{1, 2, 1, 3, 5, 2, 3}}, 5},
		{"example1", args{[]int{1, 2, 1, 3, 5, 2, 3, 1, 1, 2, 2}}, 5},
		{"example2", args{[]int{1000000, 2000000, 1000000, 30000000, 5000000, 2000000, 30000000}}, 5000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddOccurrencesInArray(tt.args.A); got != tt.want {
				t.Errorf("OddOccurrencesInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
