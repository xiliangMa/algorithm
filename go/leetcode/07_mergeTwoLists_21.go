package leetcode


type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre, tail  *ListNode

	if l1 == nil{
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val{
		pre, tail, l1 = l1, l1, l1.Next
	} else {
		pre, tail, l2 = l2, l2, l2.Next
	}

	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 == nil{
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return pre
}