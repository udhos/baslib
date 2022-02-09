package main

import (
	"github.com/udhos/baslib/baslib"
)

func main() {
	baslib.Begin()
	baslib.Open("com3:9600,N,8,1,RS,CS0,DS0,CD0", 1, 1)
	baslib.FilePrint(1, "hello serial\r\n")
	baslib.End()
}
