package main

import (
	"bufio"
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

// The following functions provide a convenient way to read different data types from standard input.
func readLine() string               { return _scln() }
func readString() string             { return _scln() }
func readStrings() []string          { return strings.Split(_scln(), " ") }
func readBool(s ...string) bool      { t, _ := strconv.ParseBool(_sc(s)); return t }
func readByte(s ...string) byte      { t, _ := strconv.ParseUint(_sc(s), 10, 8); return byte(t) }
func readDouble(s ...string) float64 { t, _ := strconv.ParseFloat(_sc(s), 64); return t }
func readFloat(s ...string) float32  { t, _ := strconv.ParseFloat(_sc(s), 32); return float32(t) }
func readInt(s ...string) int        { t, _ := strconv.Atoi(_sc(s)); return t }
func readLong(s ...string) int64     { t, _ := strconv.ParseInt(_sc(s), 10, 64); return t }
func readULong(s ...string) uint64   { t, _ := strconv.ParseUint(_sc(s), 10, 64); return t }

func readBools() []bool {
	strs := readStrings()
	arr := make([]bool, len(strs))
	for i, s := range strs {
		arr[i] = readBool(s)
	}
	return arr
}

func readBytes() []byte {
	strs := readStrings()
	arr := make([]byte, len(strs))
	for i, s := range strs {
		arr[i] = readByte(s)
	}
	return arr
}

func readDoubles() []float64 {
	strs := readStrings()
	arr := make([]float64, len(strs))
	for i, s := range strs {
		arr[i] = readDouble(s)
	}
	return arr
}

func readFloats() []float32 {
	strs := readStrings()
	arr := make([]float32, len(strs))
	for i, s := range strs {
		arr[i] = readFloat(s)
	}
	return arr
}

func readInts() []int {
	strs := readStrings()
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i] = readInt(s)
	}
	return arr
}

func readLongs() []int64 {
	strs := readStrings()
	arr := make([]int64, len(strs))
	for i, s := range strs {
		arr[i] = readLong(s)
	}
	return arr
}

func readULongs() []uint64 {
	strs := readStrings()
	arr := make([]uint64, len(strs))
	for i, s := range strs {
		arr[i] = readULong(s)
	}
	return arr
}
