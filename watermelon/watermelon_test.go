package main

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("isEven(%d)", testCase.input), func(t *testing.T) {
			result := isEven(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}

func TestCanDivideEvenly(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, true},
		{0, false},
		{-2, false},
		{-3, false},
		{10, true},
		{11, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("canDivideEvenly(%d)", testCase.input), func(t *testing.T) {
			result := canDivideEvenly(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}
