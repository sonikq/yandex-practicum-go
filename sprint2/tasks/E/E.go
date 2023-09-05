package E

type ListNode struct {
	data int
	prev *ListNode
	next *ListNode
}

func Solution(head *ListNode) *ListNode {
	next := head.next
	head.prev, head.next = head.next, head.prev

	if next != nil {
		return Solution(next)
	}
	return head
}
