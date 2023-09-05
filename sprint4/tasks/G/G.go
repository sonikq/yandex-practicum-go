package G

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func competition(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nString))
	if n == 0 {
		writer.WriteString("0")
		writer.Flush()
		return
	}
	dataString, _ := reader.ReadString('\n')
	data := strings.Fields(dataString)

	sumOfAll := 0
	resultIndexes := map[int][]int{}

	resultIndexes[0] = append(resultIndexes[0], 0)
	for i := 0; i < n; i++ {
		if data[i] == "0" {
			sumOfAll++
			resultIndexes[sumOfAll] = append(resultIndexes[sumOfAll], i+1)
		} else {
			sumOfAll--
			resultIndexes[sumOfAll] = append(resultIndexes[sumOfAll], i+1)
		}
	}
	max := 0
	for _, v := range resultIndexes {
		if len(v) == 0 {
			continue
		}
		firstAndLastIndexesRangeOfCurrentResult := v[len(v)-1] - v[0]
		if firstAndLastIndexesRangeOfCurrentResult > max {
			max = firstAndLastIndexesRangeOfCurrentResult
		}
	}
	writer.WriteString(strconv.Itoa(max))

	writer.Flush()
}

func main() {
	competition(os.Stdin, os.Stdout)
}
