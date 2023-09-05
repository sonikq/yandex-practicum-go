package A

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, m, commands, err := getInputData()
	if err != nil {
		showError(err)
	}
	_ = solution(n, m, commands, false)
}

func solution(n, m int, matrix [][]string, isTrue bool) [][]string {
	var result [][]string
	if isTrue {
		result = make([][]string, 0, m)
	}

	for i := 0; i < m; i++ {
		var line = make([]string, 0, n)
		for i2 := 0; i2 < n; i2++ {
			line = append(line, matrix[i2][i])
		}
		if isTrue {
			result = append(result, line)
		} else {
			fmt.Println(strings.Join(line, " "))
		}
	}

	return result
}

func getInputData() (n, m int, matrix [][]string, err error) {

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

	strNum, _, _ = reader.ReadLine()
	m, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	matrix = make([][]string, 0, n)
	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		str := strings.Split(strings.TrimSpace(strNums), " ")
		matrix = append(matrix, str)
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

func showError(err error) {
	panic(err)
}
