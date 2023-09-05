package E

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	if n == 0 {
		fmt.Print("0")
		return
	}

	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 10000*1000+1000)
	s.Scan()

	inputStr := strings.Split(strings.TrimSpace(s.Text()), " ")
	nums := make([]int, n)

	for i, s := range inputStr {
		nums[i], _ = strconv.Atoi(s)
	}

	sort.Ints(nums)

	sum := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum <= k {
			cnt++
		}
	}

	fmt.Print(cnt)
}
