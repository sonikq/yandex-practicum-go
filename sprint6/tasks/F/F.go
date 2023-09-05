package F

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Queued interface {
	Push(node *Node)
	Pop() (*Node, error)
	Grow()
}

type Queue struct {
	CurrentLength int
	QueueData     AdjacencyList
	PushTo        int
	TakeFrom      int
}

func (q *Queue) Push(node *Node) {
	q.QueueData[q.PushTo] = node
	q.PushTo = (q.PushTo + 1) % q.CurrentLength
	predictedIndex := (q.PushTo + 1) % q.CurrentLength
	if predictedIndex == q.TakeFrom {
		q.Grow()
	}
}

func (q *Queue) Pop() (*Node, error) {
	if q.TakeFrom == q.PushTo {
		return nil, errors.New("queue is empty")
	}
	save := q.QueueData[q.TakeFrom]
	q.TakeFrom = (q.TakeFrom + 1) % q.CurrentLength
	return save, nil
}

func (q *Queue) Grow() {
	newLength := q.CurrentLength * 2
	newList := make(AdjacencyList, newLength)
	newListIndex := 0
	for node, err := q.Pop(); err == nil; node, err = q.Pop() {
		newList[newListIndex] = node
		newListIndex++
	}
	q.TakeFrom = 0
	q.PushTo = newListIndex
	q.QueueData = newList
	q.CurrentLength = newLength
}

type Node struct {
	Value    int
	PointsTo AdjacencyList
}

type AdjacencyList []*Node

func DistanceBetweenVertices(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	vertexes, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, vertexes+1)
	for i := 0; i < vertexes; i++ {
		correctIndex := i + 1
		AL[correctIndex] = &Node{
			Value:    correctIndex,
			PointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(edgeData[0])
		peakB, _ := strconv.Atoi(edgeData[1])
		AL[peakA].PointsTo = append(AL[peakA].PointsTo, AL[peakB])
		AL[peakB].PointsTo = append(AL[peakB].PointsTo, AL[peakA])
	}
	scanner.Scan()
	searchData := strings.Fields(scanner.Text())
	from, _ := strconv.Atoi(searchData[0])
	to, _ := strconv.Atoi(searchData[1])
	writer := bufio.NewWriter(w)
	if from == to {
		writer.WriteString("0")
		writer.Flush()
		return
	}
	colors := make([]string, vertexes+1)
	distance := make([]int, vertexes+1)
	queue := Queue{
		CurrentLength: 16,
		QueueData:     make(AdjacencyList, 16),
		PushTo:        0,
		TakeFrom:      0,
	}
	queue.Push(AL[from])
	ok := bfs(&queue, colors, distance, to)

	if !ok {
		writer.WriteString("-1")
	} else {
		writer.WriteString(strconv.Itoa(distance[to]))
	}
	writer.Flush()

}

func bfs(queue Queued, colors []string, distance []int, searchFor int) bool {
	node, err := queue.Pop()
	if err != nil {
		return false
	}
	for _, n := range node.PointsTo {
		if colors[n.Value] == "" {
			distance[n.Value] = distance[node.Value] + 1
			colors[n.Value] = "gray"
			queue.Push(n)
			if n.Value == searchFor {
				return true
			}
		}
	}
	colors[node.Value] = "black"
	return bfs(queue, colors, distance, searchFor)
}

func main() {
	DistanceBetweenVertices(os.Stdin, os.Stdout)
}
