package main

import (
	"aocutil"

	"golang.org/x/exp/maps"
)

func find_start(buffer []byte, length int) int {
	set := make(map[byte]bool)
	for i := 0; (i + length - 1) < len(buffer); i++ {
		maps.Clear(set)
		for j := 0; j < length; j++ {
			set[buffer[i+j]] = true
		}
		if len(set) == length {
			return i + length
		}
	}
	return -1
}

func main() {
	buffer := aocutil.FileReadAll[[]byte]("input.txt")
	aocutil.AOCFinish(find_start(buffer, 4), find_start(buffer, 14))
}
