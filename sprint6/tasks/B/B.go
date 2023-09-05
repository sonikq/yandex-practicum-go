package B

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func ConvertListOfEdgesToAdjacencyMatrix(r io.Reader, w io.Writer) {
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	nValues := strings.Fields(scaner.Text())
	n, _ := strconv.Atoi(nValues[0])
	m, _ := strconv.Atoi(nValues[1])
	matrix := make(Matrix, n)
	for i := 0; i < m; i++ {
		scaner.Scan()
		edge := strings.Fields(scaner.Text())
		edgeA, _ := strconv.Atoi(edge[0])
		edgeB, _ := strconv.Atoi(edge[1])
		edgeA--
		edgeB--
		if matrix[edgeA] == nil {
			matrix[edgeA] = make([]int, n)
		}
		matrix[edgeA][edgeB] = 1
	}
	writer := bufio.NewWriter(w)
	for _, v := range matrix {
		if v == nil {
			v = make([]int, n)
		}
		for _, vV := range v {
			writer.WriteString(strconv.Itoa(vV) + " ")
		}
		writer.WriteByte('\n')
	}
	writer.Flush()
}

func main() {
	ConvertListOfEdgesToAdjacencyMatrix(os.Stdin, os.Stdout)
}
