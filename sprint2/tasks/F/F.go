package F

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type method struct {
	Name string
	Num  int
}

func solution(n int, cases []method) []string {

	var stack []int
	var res []string

	getMax := func() string {

		if len(stack) == 0 {
			return "None"
		}

		var max int
		for i, v := range stack {
			if i == 0 || v > max {
				max = v
			}
		}
		maxElemOfStack := fmt.Sprint(max)
		return maxElemOfStack
	}
	for i := 0; i < n; i++ {
		switch cases[i].Name {
		case "get_max":
			res = append(res, getMax())
			break
		case "push":
			stack = append(stack, cases[i].Num)
			break
		case "pop":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				res = append(res, "error")
			}
			break
		default:
			break
		}
	}
	return res
}

func main() {
	n, cases := getInput()

	res := solution(n, cases)
	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}
}

func getInput() (n int, cases []method) {

	inp, err := getInputFile()
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

	cases = make([]method, 0, n+1)

	var num int
	var met method
	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		metStr := strings.Split(strings.TrimSpace(strNums), " ")

		met = method{
			Name: metStr[0],
		}
		if len(metStr) == 2 {
			num, _ = strconv.Atoi(metStr[1])
			met.Num = num
		}

		cases = append(cases, met)
	}

	defer reader.Reset(reader)
	return
}

func getInputFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}
