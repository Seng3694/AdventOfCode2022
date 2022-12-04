package main

import (
	"aocutil"
	"strings"
)

type Range struct {
	from, to int
}

func parse_pair(s string) (Range, Range) {
	pair := strings.Split(s, ",")
	return parse_range(pair[0]), parse_range(pair[1])
}

func parse_range(s string) Range {
	r := strings.Split(s, "-")
	return Range{
		from: aocutil.Atoi(r[0]),
		to:   aocutil.Atoi(r[1]),
	}
}

func (r Range) Contains(other Range) bool {
	return r.from <= other.from && r.to >= other.to
}

func (r Range) Intersects(other Range) bool {
	return (r.from <= other.from && other.from <= r.to) || (r.from > other.from && r.from <= other.to)
}

func main() {
	part1 := 0
	part2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		elf1, elf2 := parse_pair(s)

		if elf1.Contains(elf2) || elf2.Contains(elf1) {
			part1++
		}

		if elf1.Intersects(elf2) {
			part2++
		}
	})

	aocutil.AOCFinish(part1, part2)
}
