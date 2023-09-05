package B

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func Solution(head *ListNode) {
	fmt.Println(head.data)
	if head.next != nil {
		Solution(head.next)
	}
}
