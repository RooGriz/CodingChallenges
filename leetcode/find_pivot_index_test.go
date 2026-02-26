package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case1",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Case2",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "Case3",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "Case4",
			nums: []int{2, 1, -1},
			want: 0,
		},
		{
			name: "Case5",
			nums: []int{1, 1, 1},
			want: 1,
		},
		{
			name: "Case6",
			nums: []int{0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			assert.Equalf(t, tt.want, got, "For nums %v pivot index want %d, got %d", tt.nums, tt.want, got)
		})
	}
}
