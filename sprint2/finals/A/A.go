/*
ID: 87329894

-- ПРИНЦИП РАБОТЫ --
Я реализовал дек на основе очереди на массиве, которая давалась в теории.

В моем случае кольцевой буфер может через 0 перейти в size - 1
То есть если нужно добавить элемент в начало деки или взять из начала деки, то dir = -1
А если нужно проделать те же операции, но с концом деки, то dir = 1
Ну собственно на этом и строится весь алгоритм.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
При добавлении элемента спереди мы его добавляем с правого края, предварительно сдвигая tail.
При добавлении элемента сзади элемент добавляется с левого края.
В обоих случаях размер дека увеличивается.
При удалении элемента спереди мы берём значение, двигаем хвост или голову, значение возвращается.
В таком случае размер дека уменьшается.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
И добавление и взятие с обеих сторон O(1), но количество команд поступающих на вход не константное значение,
поэтому O(n), где n количество команд, поступивших на вход

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
O(n), где n len слайса. Определяется при инициализации структуры.

*/

package A

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deck двусторонняя очередь
type Deck struct {
	stack []int
	head,
	tail,
	maxN,
	size int
}

// cmd команда действия
type cmd struct {
	action string
	num    int
}

func main() {
	q, commands, err := getInputData()
	if err != nil {
		fmt.Errorf("error")
	}

	result := q.Solution(commands)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func (d *Deck) Solution(commands []cmd) []string {
	var (
		err    error
		result []string
	)

	for i := 0; i < len(commands); i++ {
		switch commands[i].action {
		case "push_back":
			if err = d.pushBack(commands[i].num); err != nil {
				result = append(result, err.Error())
			}
			break
		case "push_front":
			if err = d.pushFront(commands[i].num); err != nil {
				result = append(result, err.Error())
			}
			break
		case "pop_front":
			result = append(result, d.popFront())
		case "pop_back":
			result = append(result, d.popBack())
		default:
			break
		}
	}

	return result
}

// добавление элемента
func (d *Deck) push(index, direction, value int) (int, error) {
	if d.isMax() {
		return index, fmt.Errorf("error")
	}

	d.changeStack(d.getStepIndex(index, direction), value)
	d.size++

	return (index + direction) % d.maxN, nil
}

// взять элемент
func (d *Deck) pop(index, direction int) (int, string) {
	if d.isEmpty() {
		return index, "error"
	}

	x := d.stack[d.getIndex(d.getStepIndex(index, direction))]
	d.changeStack(d.getStepIndex(index, direction), 0)

	d.size--
	xRes := fmt.Sprint(x)
	return (index + direction) % d.maxN, xRes
}

// добавление элемента в начало
func (d *Deck) pushFront(value int) (err error) {
	d.head, err = d.push(d.head, -1, value)
	if err != nil {
		return err
	}

	return nil
}

// добавление элемента в конец
func (d *Deck) pushBack(value int) (err error) {
	d.tail, err = d.push(d.tail, 1, value)
	if err != nil {
		return err
	}

	return nil
}

// извлечение последнего элемента
func (d *Deck) popBack() (value string) {
	d.tail, value = d.pop(d.tail, -1)
	return
}

// извлечение первого элемента
func (d *Deck) popFront() (value string) {
	d.head, value = d.pop(d.head, 1)
	return
}

// меняет значение в стеке получить индекс хвоста или головы
func (d *Deck) changeStack(index, value int) {
	d.stack[d.getIndex(index)] = value
}

// получить индекс хвоста или головы
func (d *Deck) getStepIndex(idx, dir int) int {
	if dir > 0 {
		return idx
	}
	idx += dir
	return idx
}

// перевод "обратных" индексов
func (d *Deck) getIndex(index int) int {
	if index < 0 {
		return len(d.stack) + index
	}

	return index
}

// определяет, пуст ли дек
func (d *Deck) isEmpty() bool {
	return d.size == 0
}

// определяет, переполнен ли дек
func (d *Deck) isMax() bool {
	return d.size >= d.maxN
}

// получение input из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// парсинг входных данных
func getInputData() (d *Deck, commands []cmd, err error) {

	input, err := getInputFromFile()
	if err != nil {
		fmt.Errorf("error")
	}

	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	var n, m int
	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	strNum, _, _ = reader.ReadLine()
	m, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	d = &Deck{
		stack: make([]int, m),
		maxN:  m,
	}

	var com cmd
	commands = make([]cmd, n)

	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = cmd{
			action: comStr[0],
		}
		if len(comStr) == 2 {
			com.num, _ = strconv.Atoi(comStr[1])
		}

		commands[i] = com
	}

	defer reader.Reset(reader)
	return
}
