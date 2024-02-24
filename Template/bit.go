package main

func maxUsingBit(x, y int) int {
	var mask int
	if x < y {
		mask = -1
	}

	result := x ^ ((x ^ y) & mask)
	return result
}

func minUsingBit(x, y int) int {
	var mask int
	if x < y {
		mask = -1
	}

	result := y ^ ((x ^ y) & mask)
	return result
}

func isEven(n int) bool {
	return n&1 == 0
}

func isOdd(n int) bool {
	return n&1 == 1
}
