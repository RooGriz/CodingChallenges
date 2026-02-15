package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		wantArray  []int
		wantUnique int
	}{
		{
			name:       "Empty",
			input:      []int{},
			wantArray:  []int{},
			wantUnique: 0,
		},
		{
			name:       "NoDublicates",
			input:      []int{1, 2, 3, 4, 5},
			wantArray:  []int{1, 2, 3, 4, 5},
			wantUnique: 5,
		},
		{
			name:       "HasDublicates1",
			input:      []int{1, 1, 2, 2, 2, 3, 4, 4, 5},
			wantArray:  []int{1, 2, 3, 4, 5},
			wantUnique: 5,
		},
		{
			name:       "HasDublicates2",
			input:      []int{1, 1, 1},
			wantArray:  []int{1},
			wantUnique: 1,
		},
		{
			name:       "HasDublicates3",
			input:      []int{1, 1, 1, 2},
			wantArray:  []int{1, 2},
			wantUnique: 2,
		},
		{
			name:       "HasDublicates4",
			input:      []int{1, 2, 2, 2},
			wantArray:  []int{1, 2},
			wantUnique: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnique := removeDuplicates(tt.input)
			require.Equalf(t, tt.wantUnique, gotUnique, "Unique elements in array, want %d, got %d", tt.wantUnique, gotUnique)
			require.GreaterOrEqualf(t, len(tt.input), gotUnique, "Size of input array")
			gotArray := tt.input[:gotUnique]
			assert.Equalf(t, tt.wantArray, gotArray, "Elements in array, want %v, got %v", tt.wantArray, gotArray)
		})
	}
}
