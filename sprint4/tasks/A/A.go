package A

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var q, R int

	fmt.Scanf("%d\n%d\n", &q, &R)
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	sc.Buffer(buf, 10e6)
	sc.Scan()
	str := sc.Text()

	if len(str) == 0 {
		fmt.Print(0)
		return
	}

	// fmt.Println(int('h'), int('a'), int('s'))
	// fmt.Println((((int('h')*123+int('a'))*123+int('s'))*123 + int('h')) % 100003)

	var h int
	li := len(str) - 1
	for _, s := range str[:li] {
		h = (h + int(s)) * q

		if h > R {
			h %= R
		}
	}

	h += int(str[li])
	fmt.Print(h % R)
}
