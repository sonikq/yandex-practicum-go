/*
id 87754831

-- ПРИНЦИП РАБОТЫ --

Принцип работы алгоритма заключается в том, что при добавлении элемента в хеш-таблицу сначала вычисляется хеш ключа,

Затем по этому хешу вычисляется индекс в массиве buckets, который служит указателем на начало связного списка.

Если элемент с таким же ключом уже есть в таблице, то его значение обновляется.

Если элемента с таким ключом нет, то он добавляется в связный список соответствующего индекса.

Алгоритм данного кода реализует хеш-таблицу с методом разрешения коллизий методом цепочек.

Для хэширования ключа, я вопользовался функцией Дженкинса: https://ru.wikipedia.org/wiki/%D0%A5%D0%B5%D1%88-%D1%84%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D1%8F_%D0%94%D0%B6%D0%B5%D0%BD%D0%BA%D0%B8%D0%BD%D1%81%D0%B0#%D0%A5%D0%B5%D1%88-%D1%84%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D0%B8



-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Корректность алгоритма обеспечивается тем, что при поиске, добавлении или удалении элемента происходит вычисление индекса массива buckets с помощью вычисления хеша ключа.

Это позволяет быстро находить элементы в таблице и обеспечивает хорошее распределение ключей в массиве, что уменьшает количество коллизий.

Кроме того, алгоритм обрабатывает возможные коллизии, используя метод цепочек для хранения элементов, которые имеют одинаковый индекс массива.

Если два ключа имеют одинаковый индекс хеш-таблицы, то они будут храниться в одном связном списке.



-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

Временная сложность зависит от количества коллизий и размера хеш-таблицы, мне кажется,

в среднем она составляет O(k*n) => O(n) для операций добавления, поиска и удаления элементов в хеш-таблице, где k - количество операций get/put/delete,

A n - количество элементов в хеш-таблице



-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

O(n + l), где n - количество элементов в хеш-таблице, а l - количество связных списков, используемых для разрешения коллизий.



*/

package B

import (
	"bufio"

	"fmt"

	"os"

	"strconv"

	"strings"
)

const (
	ActionGet = "get"

	ActionPut = "put"

	ActionDelete = "delete"

	ActionNone = "None"
)

// Node односвязный список

type Node struct {
	key int

	value int

	next *Node
}

type List struct {
	curr,

	prev *Node
}

// nodeList для обхода списка

var nodeList *List

// HashTable хеш таблица

type HashTable struct {
	buckets []*Node

	n int
}

func main() {

	initReader()

	data := readStrings()

	writer := bufio.NewWriter(os.Stdout)

	h := NewHashTable(1000001)

	var arrStr []string

	for i := 0; i < len(data); i++ {

		arrStr = strings.Split(data[i], " ")

		switch arrStr[0] {

		case ActionGet:

			if value, ok := h.GetKey(toInt(arrStr[1])); ok {

				writer.WriteString(fmt.Sprintf("%d\n", value))

			} else {

				writer.WriteString(ActionNone + "\n")

			}

			break

		case ActionPut:

			h.PutKey(toInt(arrStr[1]), toInt(arrStr[2]))

			break

		case ActionDelete:

			if value, ok := h.DeleteKey(toInt(arrStr[1])); ok {

				writer.WriteString(fmt.Sprintf("%d\n", value))

			} else {

				writer.WriteString(ActionNone + "\n")

			}

			break

		default:

			break

		}

	}

	writer.Flush()

}

func NewHashTable(size int) *HashTable {

	return &HashTable{

		buckets: make([]*Node, size),

		n: size,
	}

}

func (h *HashTable) PutKey(key int, value int) {

	index := h.getIndex(key)

	node := Node{key, value, nil}

	if h.buckets[index] == nil {

		h.buckets[index] = &node

		return

	}

	nodeList = &List{

		curr: h.buckets[index],

		prev: &Node{},
	}

	if find := h.findElement(key); find {

		nodeList.curr.value = value

		return

	}

	nodeList.prev.next = &node

}

func (h *HashTable) GetKey(key int) (int, bool) {

	index := h.getIndex(key)

	nodeList = &List{

		curr: h.buckets[index],
	}

	if find := h.findElement(key); find {

		return nodeList.curr.value, true

	}

	return -1, false

}

func (h *HashTable) DeleteKey(key int) (val int, find bool) {

	index := h.getIndex(key)

	nodeList = &List{

		curr: h.buckets[index],
	}

	if nodeList.curr == nil {

		return

	} else if nodeList.curr.key == key {

		h.buckets[index] = nodeList.curr.next

		return nodeList.curr.value, true

	} else {

		if find = h.findElement(key); find {

			nodeList.prev.next = nodeList.curr.next

			return nodeList.curr.value, true

		}

	}

	return

}

func (h *HashTable) hash(key int) (hash uint32) {

	hash = 0

	return uint32(key % h.n)

}

func (h *HashTable) getIndex(key int) int {

	return int(h.hash(key)) % h.n

}

func (h *HashTable) findElement(key int) bool {

	for nodeList.curr != nil {

		if nodeList.curr.key == key {

			return true

		}

		nodeList.prev = nodeList.curr

		nodeList.curr = nodeList.curr.next

	}

	return false

}

var reader *bufio.Reader

var scanner *bufio.Scanner

func initReader() {

	reader = bufio.NewReader(os.Stdin)

	scanner = bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

}

func readStrings() []string {

	var data []string

	for {

		if !scanner.Scan() {

			break

		}

		line := scanner.Text()

		if line == "" {

			break

		}

		data = append(data, line)

	}

	if err := scanner.Err(); err != nil {

		showError(err)

	}

	return data

}

func toInt(s string) int {

	intValue, err := strconv.Atoi(s)

	if err != nil {

		return -1

	}

	return intValue

}

func showError(err interface{}) {

	panic(err)

}
