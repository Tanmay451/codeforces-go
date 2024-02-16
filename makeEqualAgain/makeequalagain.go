// https://codeforces.com/contest/1931/problem/C
// 1931 C. Make Equal Again

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

// lenSameOfElementFromLeft returns the length of the contiguous subarray of the same element from the left side.
func lenSameOfElementFromLeft(ele []int) int {
	if len(ele) == 0 {
		return 0
	}
	e, l := ele[0], 0
	for i := 0; i < len(ele); i++ {
		if ele[i] == e {
			l++
		} else {
			break
		}
	}
	return l
}

// lenSameOfElementFromRight returns the length of the contiguous subarray of the same element from the right side.
func lenSameOfElementFromRight(ele []int) int {
	if len(ele) == 0 {
		return 0
	}
	e, l := ele[len(ele)-1], 0
	for i := len(ele) - 1; i >= 0; i-- {
		if ele[i] == e {
			l++
		} else {
			break
		}
	}
	return l
}

// numberOfBurles calculates the number of elements to be changed to make the array equal.
func numberOfBurles(ele []int) int {
	ll := lenSameOfElementFromLeft(ele)
	lr := lenSameOfElementFromRight(ele)

	// Check if the first and last elements are the same.
	if ele[0] == ele[len(ele)-1] {
		if ll+lr >= len(ele) {
			// If the entire array is composed of the same element, return 0 (no changes needed).
			return 0
		} else {
			// Otherwise, return the number of elements that need to be changed to make the array equal.
			return len(ele) - (ll + lr)
		}
	}

	// If the first and last elements are different, return the number of elements that need to be changed.
	return len(ele) - max(ll, lr)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	for i := 0; i < tc; i++ {
		_ = readInt(in)
		arr := readArrInt(in)
		fmt.Println(numberOfBurles(arr))
	}
}
