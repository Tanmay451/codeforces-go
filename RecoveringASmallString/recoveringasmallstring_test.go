package main

import (
	"testing"
)

func TestCalculateCharacters(t *testing.T) {
	tests := []struct {
		input int
		f     int
		s     int
		t     int
	}{
		{60, 8, 26, 26},
		{30, 1, 3, 26},
	}

	for _, test := range tests {
		f, s, th := calculateCharacters(test.input)

		if f != test.f || s != test.s || th != test.t {
			t.Errorf("For input %d, expected (%d, %d, %d), but got (%d, %d, %d)", test.input, test.f, test.s, test.t, f, s, th)
		}
	}
}
