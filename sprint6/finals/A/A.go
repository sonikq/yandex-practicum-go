/*
ID 90419469

-- ПРИНЦИП РАБОТЫ --

После считывания входного потока данных, создаётся список смежности AdjacencyList, представляющий граф.

Далее считываются данные о рёбрах и добавляются в соответствующие списки смежности вершин.

Затем вызывается функция MST(которая работает по принципу алгоритма Прима) для нахождения минимального остовного дерева.
Алгоритм начинается с вершины 1 и пошагово добавляет рёбра в MST, используя приоритетную очередь (PriorityQueue) для выбора наименьшего веса ребра,
которое соединяет уже посещённую часть графа с оставшейся.

После завершения алгоритма проверяется, что все вершины были посещены.
Если это не так, то выводится сообщение об ошибке. В противном случае выводится общий вес минимального остовного дерева.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Алгоритм Прима для нахождения MST доказан корректным. Он начинает с одной вершины (в данной реализации с вершины 1) и постепенно добавляет рёбра с минимальным весом,
которые соединяют посещённые вершины с непосещёнными.
Это обеспечивает, что дерево, которое мы строим, будет соединять все вершины графа, и его общий вес будет минимальным.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Предположим, что у нас есть V вершин и E рёбер в графе.

Инициализация AdjacencyList и сканирование входных данных занимает O(V+E) времени.

Выполнение алгоритма Прима (функция MST) занимает O(E*logE) времени,
так как на каждой итерации алгоритма мы добавляем одно ребро в приоритетную очередь и выполняем операции вставки и извлечения с минимальным приоритетом.

Проверка посещения вершин в VisitVerification занимает O(V) времени.

Итак, общая временная сложность алгоритма составляет O(V + E*logE).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

Список смежности AdjacencyList занимает O(V + E) дополнительной памяти, так как мы храним структуры данных для вершин и рёбер.

Приоритетная очередь PriorityQueue занимает O(E) дополнительной памяти, так как в ней хранятся рёбра.

Итак, общая пространственная сложность алгоритма составляет O(V + E).

*/

package A

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	weight int
	point  *Node
}

type PointsTo []Edge

type Node struct {
	value    int
	pointsTo PointsTo
}

type AdjacencyList []*Node

type PriorityQueue []Edge

func main() {
	FindExpensiveWeb(os.Stdin, os.Stdout)
}

func FindExpensiveWeb(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	vertexes, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, vertexes+1)
	for i := 0; i < vertexes; i++ {
		alValue := i + 1
		AL[alValue] = &Node{
			value:    alValue,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		vertexA, _ := strconv.Atoi(edgeData[0])
		vertexB, _ := strconv.Atoi(edgeData[1])
		weight, _ := strconv.Atoi(edgeData[2])
		AL[vertexA].pointsTo = append(AL[vertexA].pointsTo, Edge{
			weight: weight,
			point:  AL[vertexB],
		})
		AL[vertexB].pointsTo = append(AL[vertexB].pointsTo, Edge{
			weight: weight,
			point:  AL[vertexA],
		})
	}

	res, visited := MST(AL)
	writer := bufio.NewWriter(w)
	if VisitVerification(visited) {
		writer.WriteString(strconv.Itoa(res))
	} else {
		writer.WriteString("Oops! I did it again")
	}
	writer.Flush()
}

func MST(al AdjacencyList) (int, []bool) {
	queue := make(PriorityQueue, 1, cap(al))
	visited := make([]bool, cap(al))
	visited[1] = true
	total := 0
	for _, edge := range al[1].pointsTo {
		queue.Push(edge)

	}
	for edge, err := queue.Pop(); err == nil; edge, err = queue.Pop() {

		if visited[edge.point.value] {
			continue
		}
		total += edge.weight
		visited[edge.point.value] = true
		for _, e := range edge.point.pointsTo {
			if !visited[e.point.value] {
				queue.Push(e)
			}
		}
	}
	return total, visited
}

func (q *PriorityQueue) Push(edge Edge) {
	*q = append(*q, edge)
	q.TreeUp(q.Len() - 1)
}

func (q *PriorityQueue) Pop() (Edge, error) {
	if q.Len() == 1 {
		return Edge{}, errors.New("queue is empty")
	}
	save := (*q)[1]
	(*q)[1] = (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	q.TreeDown(1)
	return save, nil
}

func (q PriorityQueue) TreeUp(index int) {
	for index > 1 {
		parentIndex := index >> 1
		if !q.Less(index, parentIndex) {
			break
		}
		q.Swap(index, parentIndex)
		index = parentIndex
	}
}

func (q PriorityQueue) TreeDown(index int) {
	for {
		firstChild := index * 2
		secondChild := firstChild + 1
		if firstChild >= q.Len() {
			break
		}
		best := firstChild
		if secondChild < q.Len() && q.Less(secondChild, firstChild) {
			best = secondChild
		}
		if !q.Less(best, index) {
			break
		}
		q.Swap(index, best)
		index = best
	}
}

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(a, b int) bool {
	return q[a].weight > q[b].weight
}

func (q PriorityQueue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

func VisitVerification(visited []bool) bool {
	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
