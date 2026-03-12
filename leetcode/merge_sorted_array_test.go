package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "Case1",
			nums1: []int{},
			m:     0,
			nums2: []int{},
			n:     0,
			want:  []int{},
		},
		{
			name:  "Case2",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "Case3",
			nums1: []int{0, 0, 0},
			m:     0,
			nums2: []int{1, 2, 3},
			n:     3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "Case4",
			nums1: []int{1, 2, 3},
			m:     3,
			nums2: []int{},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "Case5",
			nums1: []int{2, 3, 4, 5, 0},
			m:     4,
			nums2: []int{1},
			n:     1,
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			require.ElementsMatchf(t, tt.want, tt.nums1, "Elements in array, want %v, got %v", tt.want, tt.nums1)
		})
	}
}
