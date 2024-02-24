// https://codeforces.com/contest/1923/problem/B
// 1923 B. Monsters Attack!
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The bufio.NewReader and os.Stdin are used to create a buffered reader for standard input.
var in = bufio.NewReader(os.Stdin)

// _scln reads a line from standard input and trims leading/trailing whitespaces.
func _scln() string {
	ln, _ := in.ReadString('\n')
	return strings.Trim(ln, " \r\n")
}

// _sc is a helper function that returns the first element of a string slice or reads a line if the slice is empty.
func _sc(s []string) string {
	if len(s) == 0 {
		return _scln()
	}
	return s[0]
}

func readStrings() []string   { return strings.Split(_scln(), " ") }
func readInt(s ...string) int { t, _ := strconv.Atoi(_sc(s)); return t }

func itrateOver(str string) {
	countA := 0
	countB := 0
	for _, e := range str {
		if string(e) == "A" {
			countA += 1
		} else if string(e) == "B" {
			countB += 1
		}
	}

	if countA > countB {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func convertStringToIntArray(str []string) []int {
	intValues := make([]int, len(str))
	for i, v := range str {
		num, _ := strconv.Atoi(v)
		intValues[i] = num
	}
	return intValues
}

func main() {
	tc := readInt()
	for tc > 0 {
		arr := convertStringToIntArray(readStrings())

		n, k := arr[0], arr[1]
		// fmt.Println()
		// fmt.Println("k:", k)
		a := convertStringToIntArray(readStrings()) // health at ith monster
		// fmt.Println("a:", a)
		x := convertStringToIntArray(readStrings()) // init position of kth monster
		// fmt.Println("x:", x)
		b := make([]int, n+1)
		for i := 0; i < len(x); i++ {
			b[abs(x[i])] += a[i]
		}
		sum := 0
		isLive := true
		for i := 0; i <= n; i++ {
			sum += b[i]
			if sum > i*k {
				isLive = false
			}
		}
		if isLive {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

		tc -= 1
	}
}
