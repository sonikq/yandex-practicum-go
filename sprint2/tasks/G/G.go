package G

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	stack,
	max []int
}

type method struct {
	Name string
	Step int
}

func main() {
	n, methods := getInputData()

	result := solution(n, methods)
	fmt.Println(strings.Join(result, "\n"))
}

func solution(n int, methods []method) []string {
	var result []string
	var q = &queue{
		stack: []int{},
		max:   []int{},
	}

	for i := 0; i < n; i++ {
		switch methods[i].Name {
		case "get_max":
			if q.isEmpty() {
				result = append(result, "None")
			} else {
				result = append(result, q.getMax())
			}
			break
		case "push":
			q.push(methods[i].Step)
			break
		case "pop":
			if q.isEmpty() {
				result = append(result, "error")
			} else {
				q.pop()
			}
			break
		default:
			break
		}
	}

	return result
}

func (q *queue) push(num int) {

	if len(q.max) > 0 && q.max[len(q.max)-1] <= num {
		q.max = append(q.max, num)
	} else if len(q.max) == 0 {
		q.max = append(q.max, num)
	}

	q.stack = append(q.stack, num)
}

func (q *queue) pop() {
	var lastIndex int
	lastIndex, q.stack = q.stack[len(q.stack)-1], q.stack[:len(q.stack)-1]

	if q.max[len(q.max)-1] == lastIndex {
		q.max = q.max[:len(q.max)-1]
	}
}

func (q *queue) getMax() string {
	return fmt.Sprint(q.max[len(q.max)-1])
}

func (q *queue) isEmpty() bool {
	return len(q.stack) == 0
}

func getInputData() (n int, methods []method) {

	inp, err := getInputFromFile()
	if err != nil {
		return
	}
	defer func(inp *os.File) {
		_ = inp.Close()
	}(inp)

	reader := bufio.NewReader(inp)

	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	methods = make([]method, 0, n+1)

	var num int
	var com method
	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = method{
			Name: comStr[0],
		}
		if len(comStr) == 2 {
			num, _ = strconv.Atoi(comStr[1])
			com.Step = num
		}

		methods = append(methods, com)
	}

	defer reader.Reset(reader)
	return
}

func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}
