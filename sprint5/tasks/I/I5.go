package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func DiffBST(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	writer := bufio.NewWriter(w)
	result := countUniqBST(1, n)
	writer.WriteString(strconv.Itoa(len(result)))
	writer.Flush()
}

func countUniqBST(left, right int) []*Node {
	nodes := make([]*Node, 0)
	if left > right {
		return append(nodes, nil)
	}

	for i := left; i <= right; i++ {
		leftNodes := countUniqBST(left, i-1)
		rightNodes := countUniqBST(i+1, right)
		for _, v := range leftNodes {
			for _, r := range rightNodes {
				nodes = append(nodes, &Node{left: v, right: r})
			}
		}
	}
	return nodes
}

//func countNodes(nodes []*Node) int {
//	count := 0
//	for _, node := range nodes {
//		if node != nil {
//			count++
//			count += countNodes([]*Node{node.left, node.right})
//		}
//	}
//	return count
//}

func main() {
	DiffBST(os.Stdin, os.Stdout)
}
