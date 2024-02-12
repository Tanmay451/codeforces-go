// https://codeforces.com/problemset/problem/4/A
// 4A - Watermelon

package main

import "fmt"

// isEven checks if a given integer is even.
// It returns true if the integer is even, otherwise false.
func isEven(n int) bool {
	if n&1 == 1 {
		// If the least significant bit is set, the number is odd.
		return false
	}
	// If the least significant bit is not set, the number is even.
	return true
}

// canDivideEvenly checks if a given integer can be divided into an even number of kilos.
// It returns true if the condition is met, otherwise false.
func canDivideEvenly(weight int) bool {
	return weight > 2 && isEven(weight)
}

func main() {
	var weight int

	// Read an integer from the user.
	fmt.Scanf("%d\n", &weight)

	// Check if the weight can be divided evenly.
	if canDivideEvenly(weight) {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}
