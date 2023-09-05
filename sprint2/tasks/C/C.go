package C

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
	}

	if head != nil {
		fmt.Println(head.data)
	}

	if head != nil && head.next != nil {
		return Solution(head.next, idx-1)
	}

	return nil
}
