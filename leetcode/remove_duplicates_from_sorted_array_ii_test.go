package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantArray []int
		wantLen   int
	}{
		{
			name:      "Empty",
			input:     []int{},
			wantArray: []int{},
			wantLen:   0,
		},
		{
			name:      "NoDublicatesII",
			input:     []int{1, 2, 3, 4, 5},
			wantArray: []int{1, 2, 3, 4, 5},
			wantLen:   5,
		},
		{
			name:      "HasDublicatesII1",
			input:     []int{1, 1, 2, 2, 2, 3, 4, 4, 5},
			wantArray: []int{1, 1, 2, 2, 3, 4, 4, 5},
			wantLen:   8,
		},
		{
			name:      "HasDublicatesII2",
			input:     []int{1, 1, 1},
			wantArray: []int{1, 1},
			wantLen:   2,
		},
		{
			name:      "HasDublicatesII3",
			input:     []int{1, 1, 1, 2},
			wantArray: []int{1, 1, 2},
			wantLen:   3,
		},
		{
			name:      "HasDublicatesII4",
			input:     []int{1, 2, 2, 2},
			wantArray: []int{1, 2, 2},
			wantLen:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnique := removeDuplicates2(tt.input)
			require.Equalf(t, tt.wantLen, gotUnique, "Unique elements in array, want %d, got %d", tt.wantLen, gotUnique)
			require.GreaterOrEqualf(t, len(tt.input), gotUnique, "Size of input array")
			gotArray := tt.input[:gotUnique]
			assert.Equalf(t, tt.wantArray, gotArray, "Elements in array, want %v, got %v", tt.wantArray, gotArray)
		})
	}
}
