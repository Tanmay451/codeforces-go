// https://codeforces.com/contest/1935/problem/A
// A. Entertainment in MAC

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

func findPosition(n int) int {
	var res int
	for i := n; i > 0; i-- {
		if i&(i-1) == 0 {
			res = i
			break
		}
	}
	return res
}

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func lls(n int, s string) string {
	if n == 0 {
		return s
	}
	rs := reverse(s)
	if rs == s {
		return s
	}

	if s < rs {
		if n%2 == 0 {
			return s
		} else {
			return s + rs
		}
	} else {
		if n%2 == 1 {
			return rs
		} else {
			return rs + s
		}
	}
}

func solve932Div2() {
	tc := readInt()
	for tc > 0 {
		n := readInt()
		s := readStrings()[0]
		ls := lls(n, s)
		fmt.Println(ls)
		tc--
	}
}

func main() {
	solve932Div2()
}
