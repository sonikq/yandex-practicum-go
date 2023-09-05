package C

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str1, str2, err := getInputData()
	if err != nil {
		showError(err)
	}

	fmt.Println(solution(str1, str2))
}

func solution(s, t string) string {
	lenS, lenT := len(s), len(t)
	if lenS == 0 || s == " " || s == t {
		return "True"
	}

	if lenS > lenT {
		return "False"
	}

	var i, j int
	for j < lenT {
		if s[i] == t[j] {
			i++
			if i >= lenS {
				return "True"
			}
		}
		j++
	}

	return "False"
}

func getInputData() (s, t string, err error) {
	var input *os.File

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)
	s, _ = reader.ReadString('\n')
	t, _ = reader.ReadString('\n')

	return
}

func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func showError(err interface{}) {
	panic(err)
}
