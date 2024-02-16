// https://codeforces.com/contest/1929/problem/B
// 1929 B. Sasha and the Drawing

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readLineNumbs reads a line of space-separated numbers and returns them as a slice of strings.
func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

// readArrInt reads a line of space-separated numbers and returns them as a slice of integers.
func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

// readInt reads a single integer from the input.
func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func min(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func colorDiagonals(n, k int) int {
	c := n + (n - 2)
	if k <= 2*c {
		return (k + 1) / 2
	}
	r := k - 2*c
	return r + c
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	for i := 0; i < tc; i++ {
		arr := readArrInt(in)
		fmt.Println(colorDiagonals(arr[0], arr[1]))
	}
}
