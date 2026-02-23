package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Case1",
			input: []int{1},
			want:  1,
		},
		{
			name:  "Case2",
			input: []int{-1},
			want:  -1,
		},
		{
			name:  "Case3",
			input: []int{-1, 0, -1, 0, 5},
			want:  5,
		},
		{
			name:  "Case4",
			input: []int{-1, 0, -1},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.input)
			assert.Equalf(t, tt.want, got, "Missing number want %d, got %d", tt.want, got)
		})
	}
}
