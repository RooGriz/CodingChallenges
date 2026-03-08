package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "WhenTreeHasNoNodes",
			root: nil,
			want: 0,
		},
		{
			name: "WhenTreeHasOneNode",
			root: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   10,
			},
			want: 1,
		},
		{
			name: "WhenTreeHasTheSameAmountOfNodesOnLeftAndRight",
			root: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   11,
				},
				Val: 10,
			},
			want: 2,
		},
		{
			name: "WhenTreeHasMoreNodesOnLeft",
			root: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{
							Left:  nil,
							Right: nil,
							Val:   12,
						},
						Right: nil,
						Val:   15,
					},
					Right: nil,
					Val:   11,
				},
				Right: &TreeNode{
					Left: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   12,
					},
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   9,
					},
					Val: 11,
				},
				Val: 10,
			},
			want: 4,
		},
		{
			name: "WhenTreeHasMoreNodesOnRight",
			root: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   9,
				},
				Right: &TreeNode{
					Left: nil,
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   9,
					},
					Val: 9,
				},
				Val: 10,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			assert.Equalf(t, tt.want, got, "Maximum depth of Binary Tree want %d, got %d", tt.want, got)
		})
	}
}
