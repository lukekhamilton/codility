package MaximumSliceProblem

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{[]int{23171, 21011, 21123, 21366, 21013, 21367}}, 356},
		{"example1", args{[]int{21013, 21367}}, 354},
		{"example2", args{[]int{21367, 21013}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args.A); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
