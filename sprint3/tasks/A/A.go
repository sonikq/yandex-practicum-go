package A

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	bracketSequence(n, 0, 0, "")
}

func bracketSequence(n, close, open int, seq string) {
	if len(seq) == 2*n {
		fmt.Println(seq)
		return
	}

	if open < n {
		bracketSequence(n, close, open+1, seq+"(") // use sb
	}

	if close < open {
		bracketSequence(n, close+1, open, seq+")")
	}
}
