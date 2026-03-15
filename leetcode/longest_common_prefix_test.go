package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Case1",
			strs: []string{},
			want: "",
		},
		{
			name: "Case2",
			strs: []string{"abc"},
			want: "abc",
		},
		{
			name: "Case3",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Case4",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			assert.Equalf(t, tt.want, got, "Longest common prefix among %v want %s, got %s", tt.strs, tt.want, got)
		})
	}
}
