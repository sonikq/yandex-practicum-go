package F

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	strs := strings.Split(strings.TrimSpace(s.Text()), " ")

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(strs[i])
	}

	sort.Ints(nums)

	var res int
LOOP:
	for i := n - 1; i > 1; i-- {
		for j := i - 1; j > 0; j-- {
			for k := i - 2; k >= 0; k-- {
				if nums[i] < nums[j]+nums[k] {
					res = nums[i] + nums[j] + nums[k]
					break LOOP
				}
			}
		}
	}

	fmt.Print(res)
}
