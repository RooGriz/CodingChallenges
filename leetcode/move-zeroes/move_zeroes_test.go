package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "ZeroesMoved1",
			input: []int{0, 0, 0, 10, 11, 12},
			want:  []int{10, 11, 12, 0, 0, 0},
		},
		{
			name:  "ZeroesMoved2",
			input: []int{10, 0, 11, 0, 12, 0},
			want:  []int{10, 11, 12, 0, 0, 0},
		},
		{
			name:  "ZeroesNotMoved",
			input: []int{10, 11, 12, 0, 0, 0},
			want:  []int{10, 11, 12, 0, 0, 0},
		},
		{
			name:  "NoZeroes",
			input: []int{10, 11, 12},
			want:  []int{10, 11, 12},
		},
		{
			name:  "AllZeroes",
			input: []int{0, 0, 0},
			want:  []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.input)
			assert.Equalf(t, tt.want, tt.input, "want %v, got %v", tt.want, tt.input)
		})
	}
}
