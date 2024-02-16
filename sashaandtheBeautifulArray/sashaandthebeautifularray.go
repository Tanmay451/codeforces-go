// https://codeforces.com/contest/1929/problem/A
// 1929 A. Sasha and the Beautiful Array

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

// maximumBeauty calculates the maximum beauty of an array by sorting it and computing the difference
// between the maximum and minimum values in the sorted array.
func maximumBeauty(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1] - arr[0]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	for i := 0; i < tc; i++ {
		_ = readInt(in)
		arr := readArrInt(in)
		fmt.Println(maximumBeauty(arr))
	}
}
