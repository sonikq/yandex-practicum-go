package E

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	pointsTo AdjacencyList
}

type AdjacencyList []*Node

func ConnectivityComponents(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, peaks+1)
	for i := 0; i < peaks; i++ {
		alIndex := i + 1
		AL[alIndex] = &Node{
			value:    alIndex,
			pointsTo: nil,
		}
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(edgeData[0])
		peakB, _ := strconv.Atoi(edgeData[1])
		AL[peakA].pointsTo = append(AL[peakA].pointsTo, AL[peakB])
		AL[peakB].pointsTo = append(AL[peakB].pointsTo, AL[peakA])
	}
	colors := make([]int, peaks+1)
	currentColor := 0
	for _, node := range AL {
		if node == nil {
			continue
		}
		if colors[node.value] == 0 {
			currentColor++
			DFC(node, colors, currentColor)
		}
	}
	result := make(map[int]*strings.Builder)
	for i, color := range colors {
		if color == 0 {
			continue
		}
		builder, ok := result[color]
		if !ok {
			builder = &strings.Builder{}
			result[color] = builder
		}
		builder.WriteString(strconv.Itoa(i) + " ")
	}
	writer := bufio.NewWriter(w)
	_, err := writer.WriteString(strconv.Itoa(currentColor))
	if err != nil {
		return
	}
	err = writer.WriteByte('\n')
	if err != nil {
		return
	}
	for i := 0; i < currentColor; i++ {
		builder := result[i+1]
		_, err = writer.WriteString(builder.String())
		if err != nil {
			return
		}
		err = writer.WriteByte('\n')
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

func DFC(node *Node, colors []int, currentColor int) {
	colors[node.value] = currentColor
	for _, n := range node.pointsTo {
		if colors[n.value] == 0 {
			DFC(n, colors, currentColor)
		}
	}
}

func main() {
	ConnectivityComponents(os.Stdin, os.Stdout)
}
