package main

import (
	"aocutil"
	"sort"
)

func main() {
	calories := make([]int, 0, 32)
	current := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		if len(s) > 0 {
			current += aocutil.Atoi(s)
		} else {
			calories = append(calories, current)
			current = 0
		}
	})

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	part1 := calories[0]
	part2 := calories[0] + calories[1] + calories[2]

	aocutil.AOCFinish(part1, part2)
}
