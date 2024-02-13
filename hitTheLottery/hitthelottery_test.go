package main

import (
	"testing"
)

func TestNumberOfBills(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{4, 4},     // 4 1-dollar bills
		{10, 1},    // 1 10-dollar bill
		{27, 4},    // 1 20-dollar bill, 1 5-dollar bill, 2 1-dollar bills
		{99, 6},    // 3 20-dollar bills, 1 10-dollar bill, 1 5-dollar bill, 1 1-dollar bill
		{123, 6},   // 1 100-dollar bill, 1 20-dollar bill, 1 1-dollar bill
		{1000, 10}, // 10 100-dollar bills
	}

	for _, testCase := range testCases {
		result := numberOfBills(testCase.input)
		if result != testCase.expected {
			t.Errorf("For input %d, expected %d, but got %d", testCase.input, testCase.expected, result)
		}
	}
}
