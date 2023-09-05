/*
ID: 87333729

-- ПРИНЦИП РАБОТЫ --
Я использую стек на массиве,ну собственно как и сказано в задании)

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм реализован в соответствии с заданием, как я раньше и указал:
Сам Алгоритм:
Если на вход подан операнд, он помещается на вершину стека.
Если на вход подан знак операции, то эта операция выполняется над требуемым количеством значений из стека, взятых в порядке добавления. Результат выполненной операции помещается на вершину стека.
Если входной набор символов обработан не полностью, возвращаемся к началу алгоритма.
После полной обработки входного набора символов результат вычисления выражения находится в вершине стека. Если в стеке осталось несколько чисел, то выводить последний добавленный элемент.

И кстати насчет кейса деления:
Округление всегда происходит вниз.
если a / b = c, то b * c – это наибольшее число, которое не превосходит a и одновременно делится без остатка на b.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
0(n), где n - количество токенов на входе

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Сложность O(n),где n - количество чисел
*/

package B

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []interface{}
	size int
}

func NewStack() *Stack {

	return &Stack{
		data: make([]interface{}, 0),
		size: 0,
	}
}

func (s *Stack) Push(val interface{}) {
	if len(s.data) == s.size {
		s.data = append(s.data, val)
	} else {
		s.data[s.size] = val
	}

	s.size++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.size == 0 {
		return "", errors.New("empty")
	}

	s.size--
	return s.data[s.size], nil
}

func main() {
	initReader()
	values := readStrings()
	s := NewStack()

	for i := 0; i < len(values); i++ {
		if val, err := strconv.Atoi(values[i]); err == nil {
			s.Push(val)
		} else {
			second, err := s.Pop()
			if err != nil {
				return
			}

			first, err := s.Pop()
			if err != nil {
				return
			}
			s.Push(Solution(first.(int), second.(int), values[i]))
		}
	}

	pop, err := s.Pop()
	if err != nil {
		return
	}
	writeData(strconv.Itoa(pop.(int)))
}

func Solution(firstNumber int, secondNumber int, operand string) int {
	switch strings.TrimSpace(operand) {
	case "+":
		return firstNumber + secondNumber
	case "-":
		return firstNumber - secondNumber
	case "*":
		return firstNumber * secondNumber
	case "/":
		if firstNumber < 0 && secondNumber > 0 && firstNumber%secondNumber != 0 {
			return firstNumber/secondNumber - 1
		}

		return firstNumber / secondNumber
	}

	return -1
}

var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader() {
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
}

func readStrings() []string {
	var data []string

	for scanner.Scan() {
		rawString := scanner.Text()
		data = append(data, rawString)
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}
