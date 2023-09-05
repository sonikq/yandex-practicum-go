package E

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 100000)
	s.Scan()
	m := make(map[rune]int)
	max := 0
	str := s.Text()
	j := -1
	for i, c := range str {
		if nj, ok := m[c]; ok && nj > j {
			j = nj
		}
		if i-j > max {
			max = i - j
		}
		m[c] = i
	}

	fmt.Print(max)
}
