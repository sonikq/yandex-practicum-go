package K

import "fmt"

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func Solution(root *Node, left int, right int) {
	if root == nil {
		return
	}

	if left <= root.value {
		Solution(root.left, left, right)
	}

	if left <= root.value && right >= root.value {
		fmt.Println(root.value)
	}

	if right >= root.value {
		Solution(root.right, left, right)
	}
}
