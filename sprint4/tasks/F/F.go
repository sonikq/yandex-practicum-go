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
	s := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 4096)
	s.Buffer(buf, 6000*100+6000)
	s.Scan()
	s.Scan()
	ws := strings.Split(s.Text(), " ")
	anagrams := make(map[string][]int)
	for i, w := range ws {
		r := []rune(w)
		sort.Sort(SortRunes(r))
		sr := string(r)
		if _, ok := anagrams[sr]; !ok {
			anagrams[sr] = make([]int, 0, 10)
		}

		anagrams[sr] = append(anagrams[sr], i)
	}

	sorted := make([][]int, 0, len(anagrams))
	for _, indx := range anagrams {
		sorted = append(sorted, indx)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	var res strings.Builder
	for _, s := range sorted {
		for _, ss := range s {
			res.WriteString(strconv.Itoa(ss) + " ")
		}
		res.WriteString("\n")
	}

	fmt.Print(res.String())
}

type SortRunes []rune

func (sr SortRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr SortRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr SortRunes) Len() int {
	return len(sr)
}
