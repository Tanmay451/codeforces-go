// https://codeforces.com/contest/1931/problem/B
// 1931 B. Make Equal
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

// canMakeEqual checks if it's possible to make the weights of the array elements equal
// by redistributing excess weights to the next element.
func canMakeEqual(arr []int) bool {
	tW := 0
	for _, w := range arr {
		tW += w
	}

	aW := tW / len(arr)

	for i, w := range arr {
		if w < aW {
			return false
		}
		if w > aW && i < len(arr)-1 {
			arr[i+1] += w - aW
		}
	}
	return true
}

// main function to read the number of test cases, read and process each test case,
// and print the result (YES or NO) based on whether it's possible to make weights equal.
func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in) // Read the number of test cases
	for i := 0; i < tc; i++ {
		_ = readInt(in)       // Discard the size of the array (not used in the current implementation)
		arr := readArrInt(in) // Read the array of weights
		if canMakeEqual(arr) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
