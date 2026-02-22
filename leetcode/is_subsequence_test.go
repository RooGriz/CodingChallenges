package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "IsSubsequenceWhenBothEmpty",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "IsSubsequenceWhenOnlySEmpty",
			s:    "",
			t:    "abc",
			want: true,
		},
		{
			name: "NotSubsequenceWhenOnlyTEmpty",
			s:    "abc",
			t:    "",
			want: false,
		},
		{
			name: "IsSubsequence",
			s:    "abc",
			t:    "atbtct",
			want: true,
		},
		{
			name: "IsSubsequenceWhenBothEqual",
			s:    "abc",
			t:    "abc",
			want: true,
		},
		{
			name: "NotSubsequence1",
			s:    "abc",
			t:    "atbtt",
			want: false,
		},
		{
			name: "NotSubsequence2",
			s:    "abc",
			t:    "ab",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			assert.Equalf(t, tt.want, got, "For string t = %s and s = %s, want %t, got %t", tt.t, tt.s, tt.want, got)
		})
	}
}
