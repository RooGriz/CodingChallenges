package leetcode

// Time complexity - O(n), space complexity - O(1)
func swapPairs(head *ListNode) *ListNode {
	var prev *ListNode
	first := head
	for first != nil {
		if first.Next == nil {
			if prev != nil {
				prev.Next = first
			}
			break
		}
		second := first.Next

		secondNext := second.Next
		second.Next = first
		first.Next = secondNext
		if prev != nil {
			prev.Next = second
		} else {
			head = second
		}

		prev = first
		first = secondNext
	}
	return head
}
