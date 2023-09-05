package D

type ListNode struct {
	data string
	next *ListNode
}

var idx = 0

func Solution(head *ListNode, elem string) int {
	if elem == head.data {
		return idx
	}

	if head.next != nil {
		idx++
	} else {
		return -1
	}

	return Solution(head.next, elem)
}
