/*

ID 89806686

Хотел бы уточнить пвру важных моментов, почему вы предлагаете использовать snake case, вместо camel case(вы предложили назвать поле структуры tasks_solved)?
И второй момент - это рекурсия в функциях SiftUp и SiftDown, почему я ее не могу использовать?!
Если даже в теории написан единственный пример на просеивание - и это как раз с использованием рекурсии(привожу пример псевдокода из теории ниже).

Псевдокод из раздела "Приоритетная очередь. Вставка и удаление":
функция sift_up(heap, index):
    если index == 1, то
        завершить работу
    parent_index = index / 2  (целочисленное деление)
    если heap[parent_index] < heap[index], то
        обменять местами heap[parent_index] и heap[index]
        sift_up(heap, parent_index)

-- ПРИНЦИП РАБОТЫ --

Считывание входных данных: Сначала считывается число n - количество элементов, а затем n строк, представляющих элементы, состоящие из имени, количества завершенных задач (tasksSolved), и количества неудачных задач (penalty).

Инициализация структуры "Пирамида" (Queue): Создается структура данных "Пирамида" (min-куча) с начальным пустым массивом queue и указателем lastIndex на последний элемент в очереди.

Добавление элементов в "Пирамиду": Каждый считанный элемент добавляется в "пирамиду" с помощью метода Append.
При добавлении элемента, он помещается в конец массива queue, затем выполняется просеивание вверх SiftUp, которая перемещает элемент вверх по пирамиде до тех пор, пока не будет восстановлено условие мин-кучи.

Извлечение элементов из "Пирамиды" и сортировка: После того как все элементы добавлены в "пирамиду", начинается процесс извлечения элементов методом GetElem.
При извлечении элемента, он меняется местами с корневым элементом пирамиды, затем выполняется просеивание вниз SiftDown, которое опускает элемент вниз по пирамиде до тех пор,
пока не будет восстановлено условие мин-кучи. Это позволяет извлекать элементы в порядке убывания приоритета (сначала по завершенным задачам, затем по неудачным, и по имени).

Вывод отсортированных элементов: Отсортированные элементы выводятся в заданном порядке в формате: имя элемента, за которым следует символ новой строки.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Построение мин-кучи (мин-пирамиды):
	Корректность алгоритма начинается с того, что он строит мин-кучу.
	Это означает, что для каждого элемента i его приоритет (сравниваемые значения) меньше или равен приоритету его потомков,
	что обеспечивает правильное размещение элементов в структуре данных.

Извлечение с наивысшим приоритетом:
	При извлечении элемента из мин-кучи, извлекается элемент с наивысшим приоритетом,
	что соответствует определенному правилу сортировки (сначала по завершенным задачам, затем по неудачным, и по имени).
	После извлечения элемента, происходит переупорядочивание пирамиды, чтобы сохранить свойства мин-кучи.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Добавление элемента в "пирамиду" (Append): выполняется через просеивание вниз(ShiftUp) => O(log n)
Извлечение элемента из "пирамиды" (GetElem): выполняется через просеивание вниз(ShiftDown) => O(log n)
Построение "пирамиды" (queue): поскольку каждый из n элементов добавляется в пирамиду, общая сложность построения пирамиды составляет O(n * log n), где n - кол-во входных данных.
Вывод отсортированных элементов: O(n)

Итоговая временная сложность алгоритма: O(n * log n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Хранение пирамиды:
	Пирамида (мин-куча) хранится в массиве размером "n + 1" (queue), где "n" - количество элементов.
	Дополнительно, несколько переменных (например, lastIndex) также используются для управления состоянием пирамиды.
	Пространственная сложность хранения пирамиды составляет O(n).

Прочие переменные и структуры данных:
	Кроме массива для пирамиды, другие переменные и структуры данных (например, переменные для временного хранения
	считанных данных и временных переменных внутри методов) занимают константное количество памяти.
	Пространственная сложность остальных переменных - O(1).

Функции SiftUp и SiftDown, вызываемые рекурсивно, могут добавить до O(log n) дополнительной памяти в стек вызовов.
Таким образом, общая пространственная сложность алгоритма будет составлять O(n + log n), где n - количество элементов во входных данных,
а O(n) идет на хранение данных в структуре Queue, а O(log n) - на рекурсивные вызовы функций.

Итоговая пространственная сложность алгоритма: O(n+log n)

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
	if index == 1 {
		return
	}
	parentIndex := index >> 1
	if q.Less(index, parentIndex) {
		q.Swap(index, parentIndex)
		q.TreeUp(parentIndex)
	}
}

func (q PriorityQueue) TreeDown(index int) {
	firstChild := index * 2
	secondChild := firstChild + 1
	if firstChild >= q.Len() {
		return
	}
	best := firstChild
	if secondChild < q.Len() && q.Less(secondChild, firstChild) {
		best = secondChild
	}
	if q.Less(best, index) {
		q.Swap(best, index)
		q.TreeDown(best)
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
