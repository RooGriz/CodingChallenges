package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "WhenBothEmpty",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "WhenFirstIsEmpty",
			list1: nil,
			list2: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
		},
		{
			name: "WhenSecondIsEmpty",
			list1: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
			list2: nil,
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  11,
					Next: nil,
				},
			},
		},
		{
			name: "WhenListsAreSame",
			list1: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  12,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
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
				Val: 10,
				Next: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val: 11,
						Next: &ListNode{
							Val: 11,
							Next: &ListNode{
								Val: 12,
								Next: &ListNode{
									Val:  12,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "WhenFirstListFollowedBySecond",
			list1: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 14,
					Next: &ListNode{
						Val:  15,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
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
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 14,
								Next: &ListNode{
									Val:  15,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "WhenSecondListFollowedByFirst",
			list1: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val:  12,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 14,
					Next: &ListNode{
						Val:  15,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 14,
								Next: &ListNode{
									Val:  15,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "WhenOrderedRandomly",
			list1: &ListNode{
				Val: 11,
				Next: &ListNode{
					Val: 12,
					Next: &ListNode{
						Val:  15,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 13,
					Next: &ListNode{
						Val:  14,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val: 11,
					Next: &ListNode{
						Val: 12,
						Next: &ListNode{
							Val: 13,
							Next: &ListNode{
								Val: 14,
								Next: &ListNode{
									Val:  15,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead := mergeTwoLists(tt.list1, tt.list2)
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
