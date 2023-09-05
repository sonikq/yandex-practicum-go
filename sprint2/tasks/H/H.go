package H

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := getInputData()
	if err != nil {
		showError(err)
	}
	result := isCorrectBracketSeq(input)
	runes := []rune(strconv.FormatBool(result))
	runes[0] = unicode.ToUpper(runes[0])
	fmt.Println(string(runes))
}

func isCorrectBracketSeq(s string) bool {
	var stack []rune
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (v == ')' && pop != '(') || (v == '}' && pop != '{') || (v == ']' && pop != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}
func getInputData() (input string, err error) {

	inp, err := getInputFromFile()
	if err != nil {
		return
	}
	defer func(inp *os.File) {
		_ = inp.Close()
	}(inp)

	reader := bufio.NewReader(inp)

	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

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
