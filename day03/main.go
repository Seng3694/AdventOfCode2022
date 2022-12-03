package main

import (
	"aocutil"
)

func get_item_value(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	} else {
		return int(item - 'A' + 27)
	}
}

func main() {
	rucksacks := make([]string, 0, 300)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		rucksacks = append(rucksacks, s)
	})

	part1 := 0
	part2 := 0
	for i, r := range rucksacks {
		half := len(r) / 2
	exit_part1:
		for j := 0; j < half; j++ {
			for k := half; k < len(r); k++ {
				if r[j] == r[k] {
					part1 += get_item_value(r[j])
					break exit_part1
				}
			}
		}

		if i%3 == 0 {
		exit_part2:
			for r1 := 0; r1 < len(rucksacks[i]); r1++ {
				for r2 := 0; r2 < len(rucksacks[i+1]); r2++ {
					if rucksacks[i][r1] != rucksacks[i+1][r2] {
						continue
					}
					for r3 := 0; r3 < len(rucksacks[i+2]); r3++ {
						if rucksacks[i][r1] == rucksacks[i+2][r3] {
							part2 += get_item_value(rucksacks[i][r1])
							break exit_part2
						}
					}
				}
			}
		}
	}

	aocutil.AOCFinish(part1, part2)
}
