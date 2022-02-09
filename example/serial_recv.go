package main

import (
	"fmt"
	"time"

	"github.com/udhos/baslib/baslib"
)

func main() {
	baslib.Begin()
	baslib.Open("com3:9600,N,8,1,RS,CS0,DS0,CD0", 1, 1)
	for {
		time.Sleep(1 * time.Second)
		input := baslib.FileInputCount(1, 1)
		fmt.Print(input)
	}
	baslib.End()
}
