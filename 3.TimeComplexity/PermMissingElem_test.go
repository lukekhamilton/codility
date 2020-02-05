package TimeComplexity

import "testing"

func TestPermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example0", args{A: []int{2, 3, 1, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermMissingElem(tt.args.A); got != tt.want {
				t.Errorf("PermMissingElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
