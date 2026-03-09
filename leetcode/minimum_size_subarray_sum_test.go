package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "Case1",
			target: 7,
			nums:   []int{7},
			want:   1,
		},
		{
			name:   "Case2",
			target: 6,
			nums:   []int{2, 3, 1, 2, 4},
			want:   2,
		},
		{
			name:   "Case3",
			target: 4,
			nums:   []int{1, 1, 1, 1, 4, 1},
			want:   1,
		},
		{
			name:   "Case4",
			target: 5,
			nums:   []int{2, 3},
			want:   2,
		},
		{
			name:   "Case5",
			target: 5,
			nums:   []int{2, 3, 5},
			want:   1,
		},
		{
			name:   "Case6",
			target: 6,
			nums:   []int{1, 2, 3, 3},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			require.Equalf(t, tt.want, got, "For target %d and nums %v, want %d, got %d", tt.target, tt.nums, tt.want, got)
		})
	}
}
