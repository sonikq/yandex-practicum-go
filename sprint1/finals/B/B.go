// id:86208824
package B

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sleightOfHands(matrixSize int, k int, numbers []string) int {
	var result int
	var mas = make([]int, 9)
	k *= 2

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			x := numbers[i][j]
			if x != '.' {
				if contains(mas, mas[numbers[i][j]-49]); false {
					mas[numbers[i][j]-49] = 1
				} else {
					mas[numbers[i][j]-49]++
				}
			}
		}
	}

	for i := range mas {
		if k >= mas[i] && mas[i] != 0 {
			result++
		}
	}

	return result
}

func main() {
	scanner := makeScanner()
	const matrixSize = 4
	k := readInt(scanner)
	matrix := readMatrix(scanner, matrixSize)
	hands := sleightOfHands(matrixSize, k, matrix)
	fmt.Print(hands)
}

func contains(arr []int, e int) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, matrixSize int) []string {
	matrix := make([]string, 4)
	for i := 0; i < matrixSize; i++ {
		matrix[i] = readLine(scanner)
	}
	return matrix
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
