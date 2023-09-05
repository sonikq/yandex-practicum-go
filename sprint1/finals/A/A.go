// id:86208561
package A

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getNearZero(n int, houses []int) []int {

	var zeros = make([]int, n)
	var count = n * 2
	var check int

	for i := 0; i < n; i++ {
		if houses[i] == 0 {
			if i+1 < n && houses[i+1] == 0 {
				continue
			}
			zeros[i] = count
			count = 1
			continue
		}

		if count < n {
			houses[i] = count
		} else {
			houses[i] = n
		}

		count++
	}

	for i := n - 1; i >= 0; i-- {
		if houses[i] == 0 {
			if i+1 < n && houses[i+1] == 0 {
				continue
			}
			count = 1
			check = zeros[i]
			continue
		}

		if check > 0 {
			if count < (check - count) {
				houses[i] = count
			}
			count++
		}
	}

	return houses
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(scanner)
	nearZero := getNearZero(n, arr)
	printArray(nearZero)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 11 * 1024 * 1024
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

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
