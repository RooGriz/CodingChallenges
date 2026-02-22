package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "OneNumber",
			input: []int{},
			want:  0,
		},
		{
			name:  "TwoNumbers",
			input: []int{1},
			want:  0,
		},
		{
			name:  "ThreeNumbers",
			input: []int{2, 0},
			want:  1,
		},
		{
			name:  "FourNumbers",
			input: []int{3, 0, 1},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(tt.input)
			assert.Equalf(t, tt.want, got, "Missing number want %d, got %d", tt.want, got)
		})
	}
}
