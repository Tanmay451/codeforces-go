package main

import "fmt"

func write(arg ...interface{})                 { fmt.Print(arg...) }
func writeLine(arg ...interface{})             { fmt.Println(arg...) }
func writeFormat(f string, arg ...interface{}) { fmt.Printf(f, arg...) }
