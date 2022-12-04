package main

import (
	"aocutil"
	"strings"
)

func main() {
	part1 := 0
	part2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		pair := strings.Split(s, ",")
		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")
		elf1 := struct{ from, to int }{
			from: aocutil.Atoi(range1[0]),
			to:   aocutil.Atoi(range1[1]),
		}
		elf2 := struct{ from, to int }{
			from: aocutil.Atoi(range2[0]),
			to:   aocutil.Atoi(range2[1]),
		}

		if (elf1.from <= elf2.from && elf1.to >= elf2.to) || (elf1.from >= elf2.from && elf2.to >= elf1.to) {
			part1++
		}

		if elf1.from > elf2.from {
			//make sure that elf1 has the smallest start
			elf1, elf2 = elf2, elf1
		}
		if elf2.from <= elf1.to {
			part2++
		}
	})

	aocutil.AOCFinish(part1, part2)
}
