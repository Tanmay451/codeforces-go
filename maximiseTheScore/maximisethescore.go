// https://codeforces.com/contest/1930/problem/A
// 1930 A. Maximise The Score

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isEven(n int) bool {
	return n&1 == 0
}
func maximumFinalScore(arr []int) int {
	sort.Ints(arr)
	sum := 0
	for i := 0; i < len(arr); i++ {
		if isEven(i) {
			sum += arr[i]
		}
	}
	return sum
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	for i := 0; i < tc; i++ {
		_ = readInt(in)
		arr := readArrInt(in)
		fmt.Println(maximumFinalScore(arr))
	}
}
