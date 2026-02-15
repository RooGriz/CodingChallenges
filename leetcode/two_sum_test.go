package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Empty input",
			nums:   []int{},
			target: 0,
			want:   []int{0, 0},
		},
		{
			name:   "HasPairOfNumbers",
			nums:   []int{1, -1, 4, 5},
			target: 6,
			want:   []int{0, 3},
		},
		{
			name:   "NoPairOfNumbers",
			nums:   []int{1, -1, 4, 5},
			target: 7,
			want:   []int{0, 0},
		},
		{
			name:   "HasPairWhenTwoZeros",
			nums:   []int{0, 0},
			target: 0,
			want:   []int{0, 1},
		},
		{
			name:   "HasPairWhenAllNegative",
			nums:   []int{-1, -2, -3},
			target: -3,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			assert.ElementsMatchf(t, tt.want, got, "For nums %v and target %d, want %d, got %d", tt.nums, tt.target, tt.want, got)
		})
	}
}
