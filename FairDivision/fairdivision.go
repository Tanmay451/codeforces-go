// https://codeforces.com/contest/1472/problem/B
// B. Fair Division

package main

import (
	"fmt"
	"sort"
)

func canDivideEqually(arr []int) bool {
	sort.Ints(arr)
	s1, s2 := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		if s1 <= s2 {
			s1 += arr[i]
		} else {
			s2 += arr[i]
		}
	}
	return s1 == s2
}

func readIntArray(size int) []int {
	arr := make([]int, size)
	for j := 0; j < size; j++ {
		fmt.Scanf("%d", &arr[j])
	}
	return arr
}

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var size int
		fmt.Scanf("%d\n", &size)
		arr := readIntArray(size)
		if canDivideEqually(arr) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
