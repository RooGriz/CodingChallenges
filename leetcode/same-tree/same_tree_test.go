package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "EmptyTrees",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "SameTrees",
			p: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   12,
				},
				Val: 10,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   12,
				},
				Val: 10,
			},
			want: true,
		},
		{
			name: "DifferentTrees1",
			p: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   12,
				},
				Val: 10,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   22,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "DifferentTrees2",
			p: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   12,
				},
				Val: 10,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   22,
				},
				Val: 10,
			},
			want: false,
		},
		{
			name: "DifferentTrees3",
			p: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   12,
				},
				Val: 10,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: nil,
				Val:   10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			assert.Equalf(t, tt.want, got, "Want %t, got %t", tt.want, got)
		})
	}
}
