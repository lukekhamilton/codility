package TimeComplexity

import "testing"

func TestFrogJmp(t *testing.T) {
	type args struct {
		X int
		Y int
		D int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{X: 10, Y: 85, D: 30}, 3},
		{"example1", args{X: 1, Y: 855, D: 3}, 285},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrogJmp(tt.args.X, tt.args.Y, tt.args.D); got != tt.want {
				t.Errorf("FrogJmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
