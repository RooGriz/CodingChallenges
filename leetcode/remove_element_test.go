package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		val       int
		wantArray []int
		wantLen   int
	}{
		{
			name:      "Empty",
			nums:      []int{},
			val:       1,
			wantArray: []int{},
			wantLen:   0,
		},
		{
			name:      "NoValuesInArray",
			nums:      []int{1, 2, 3, 4, 5},
			val:       6,
			wantArray: []int{1, 2, 3, 4, 5},
			wantLen:   5,
		},
		{
			name:      "AllValuesInArray",
			nums:      []int{1, 1, 1},
			val:       1,
			wantArray: []int{},
			wantLen:   0,
		},
		{
			name:      "SeveralValuesInArray",
			nums:      []int{1, 2, 1, 3, 1, 4},
			val:       1,
			wantArray: []int{2, 3, 4},
			wantLen:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeElement(tt.nums, tt.val)
			require.Equalf(t, tt.wantLen, gotLen, "Amount of elements in the array, want %d, got %d", tt.wantLen, gotLen)
			require.GreaterOrEqualf(t, len(tt.nums), gotLen, "Size of input array")
			gotArray := tt.nums[:gotLen]
			assert.Equalf(t, tt.wantArray, gotArray, "Elements in array, want %v, got %v", tt.wantArray, gotArray)
		})
	}
}
