package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Case1",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "Case2",
			input: []int{1, 2, 3},
			want:  []int{1, 3, 6},
		},
		{
			name:  "Case3",
			input: []int{1, 1, 1, 1, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runningSum(tt.input)
			require.Equalf(t, tt.want, got, "For input %b want %v, got %v", tt.input, tt.want, got)
		})
	}
}
