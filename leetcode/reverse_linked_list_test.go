package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "WhenEmpty",
			head: nil,
			want: nil,
		},
		{
			name: "WhenOneNodeList",
			head: &ListNode{
				Val:  10,
				Next: nil,
			},
			want: &ListNode{
				Val:  10,
				Next: nil,
			},
		},
		{
			name: "WhenTwoNodesList",
			head: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 11,
				Next: &ListNode{
					Val:  10,
					Next: nil,
				},
			},
		},
		{
			name: "WhenThreeNodesList",
			head: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  12,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 12,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  10,
						Next: nil,
					},
				},
			},
		},
		{
			name: "WhenFourNodesList",
			head: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val:  13,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 12,
					Next: &ListNode{
						Val: 11,
						Next: &ListNode{
							Val:  10,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead := reverseList(tt.head)
			gotCur := gotHead
			wantCur := tt.want
			for i := 0; ; i++ {
				require.Equalf(t, wantCur != nil, gotCur != nil, "Wrong node %d", i)
				if wantCur == nil || gotCur == nil {
					break
				}
				require.Equalf(t, wantCur.Val, gotCur.Val, "Wrong node %d", i)
				gotCur = gotCur.Next
				wantCur = wantCur.Next
			}
		})
	}
}
