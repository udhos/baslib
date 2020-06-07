package main

import (
	"github.com/udhos/baslib/baslib"
)

func main() {
	baslib.Begin()
	baslib.Print("hello")
	baslib.Println(" baslib")
	baslib.Println(baslib.MidNew("1234", 2, 2, "abc"))       // output: 1ab4
	baslib.Println(baslib.MidNew("1234567", 3, 3, "abcdef")) // output: 12abc67
	baslib.End()
}
