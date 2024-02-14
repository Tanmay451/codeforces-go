// https://codeforces.com/contest/1931/problem/A
// 1931 A. Recovering a Small String

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateCharacters(x int) (int, int, int) {
	f := max(1, x-52)
	s := max(1, x-f-26)
	t := max(1, x-f-s)
	return f, s, t
}

func printResult(f, s, t int) {
	asciiValueA := int('a')
	fmt.Printf("%c%c%c\n", asciiValueA+f-1, asciiValueA+s-1, asciiValueA+t-1)
}

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var x int
		fmt.Scanf("%d\n", &x)

		f, s, t := calculateCharacters(x)
		printResult(f, s, t)
	}
}
