package A

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	pointsTo AdjacenceList
	value    int
}

type AdjacenceList []*Node

func BuildAdjacencyList(r io.Reader, w io.Writer) {
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	nValues := strings.Fields(scaner.Text())
	n, _ := strconv.Atoi(nValues[0])
	m, _ := strconv.Atoi(nValues[1])
	al := make(AdjacenceList, n+1)
	for i := 0; i < m; i++ {
		scaner.Scan()
		edge := strings.Fields(scaner.Text())
		edgeA, _ := strconv.Atoi(edge[0])
		edgeB, _ := strconv.Atoi(edge[1])
		if al[edgeA] == nil {
			al[edgeA] = &Node{value: edgeA}
		}
		if al[edgeB] == nil {
			al[edgeB] = &Node{value: edgeB}
		}
		al[edgeA].pointsTo = append(al[edgeA].pointsTo, al[edgeB])
	}
	writer := bufio.NewWriter(w)
	for i, v := range al {
		if i == 0 {
			continue
		}
		if v == nil {
			writer.WriteString("0")
			writer.WriteByte('\n')
			continue
		}
		if len(v.pointsTo) == 0 {
			writer.WriteString("0")
			writer.WriteByte('\n')
			continue
		}
		writer.WriteString(strconv.Itoa(len(v.pointsTo)) + " ")
		for _, g := range v.pointsTo {
			writer.WriteString(strconv.Itoa(g.value) + " ")
		}
		writer.WriteByte('\n')

	}
	writer.Flush()
}

func main() {
	BuildAdjacencyList(os.Stdin, os.Stdout)
}
