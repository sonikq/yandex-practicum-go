package E

import "errors"

//type Node struct {
//	value int
//	left  *Node
//	right *Node
//}

func Solution(root *Node) bool {
	_, err := checkSortTree(root)
	if err != nil {
		return false
	}
	return true
}

func checkSortTree(node *Node) ([]int, error) {
	if node == nil {
		return []int{}, nil
	}

	if node.left == nil && node.right == nil {
		return []int{node.value}, nil
	}

	lVals, err := checkSortTree(node.left)
	if err != nil {
		return nil, errors.New("Wrong left tree ")
	}

	rVals, err := checkSortTree(node.right)
	if err != nil {
		return nil, errors.New("Wrong right tree ")
	}

	for _, v := range lVals {
		if v >= node.value {
			return nil, errors.New("wrong left tree")
		}
	}
	for _, v := range rVals {
		if v <= node.value {
			return nil, errors.New("wrong right tree")
		}
	}
	return append(lVals, rVals...), nil
}
