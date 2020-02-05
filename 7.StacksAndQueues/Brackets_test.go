package StacksAndQueues

import "testing"

func TestBrackets(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"{[()()]}"}, 1},
		{"example2", args{"([)()]"}, 0},
		{"example3", args{"([)()"}, 0},
		{"example4", args{")"}, 0},
		{"example5", args{"{[(){([])}()]}"}, 1},
		{"example6", args{"{}"}, 1},
		{"example7", args{"()"}, 1},
		{"example8", args{"{}"}, 1},
		{"example9", args{"[([][])()]"}, 1},
		{"example10", args{"([)()]"}, 0},
		{"example11", args{"(]"}, 0},
		{"example12", args{""}, 1},
		{"example13", args{"{{{{{"}, 0},
		{"example14", args{"))"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Brackets(tt.args.S); got != tt.want {
				t.Errorf("Brackets() = %v, want %v", got, tt.want)
			}
		})
	}
}
