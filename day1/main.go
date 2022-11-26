package main

import (
	"aocutil"
	"fmt"
)

func main() {
	for s := range aocutil.FileReadAllLines("day1/main.go") {
		fmt.Printf("%v\n", s)
	}

	aocutil.AOCFinish(10, 20)
}
