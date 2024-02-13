// https://codeforces.com/problemset/problem/996/A
// 996 A. Hit the Lottery

package main

import "fmt"

// numberOfBills function takes an amount in dollars and returns the minimum number of bills required.
func numberOfBills(n int) int {
	bills := 0
	bills += countBills(&n, 100)
	bills += countBills(&n, 20)
	bills += countBills(&n, 10)
	bills += countBills(&n, 5)
	bills += countBills(&n, 1)

	return bills
}

// countBills function calculates the number of bills of a specific denomination needed for the given amount.
// It updates the amount and returns the count of bills.
func countBills(n *int, denomination int) int {
	count := 0
	if *n/denomination > 0 {
		count += *n / denomination
		*n = *n % denomination
	}
	return count
}

// The main function reads the amount from the user and prints the minimum number of bills required.
func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(numberOfBills(n))
}
