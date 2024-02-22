// https://codeforces.com/problemset/problem/1926/C
// 1926 C. Vlad and a Sum of Sum of Digits

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

func digitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number = number / 10
	}
	return sum
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	s := 2e5 + 1
	arr := make([]int, int(s))
	for i := 1; i < int(s); i++ {
		arr[i] = arr[i-1] + digitSum(i)
	}
	for i := 0; i < tc; i++ {
		n := readInt(in)
		fmt.Println(arr[n])
	}
}
