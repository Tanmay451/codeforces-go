package main

import (
	"testing"
)

func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEven(100)
	}
}

func BenchmarkCanDivideEvenly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canDivideEvenly(100)
	}
}
