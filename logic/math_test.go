package logic

import (
	"testing"
)

func TestIsSquare(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{1, true},
		{2, false},
		{4, true},
		{9, true},
	}
	for _, tt := range tests {
		if got := IsSquare(tt.input); got != tt.want {
			t.Errorf("IsSquare(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsFibonacci(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
		{8, true},

		{317811, true},
		{317812, false},
	}
	for _, tt := range tests {
		if got := IsFibonacci(tt.input); got != tt.want {
			t.Errorf("IsFibonacci(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
