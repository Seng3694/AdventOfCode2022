package main

import (
	"aocutil"
	"fmt"
)

type List struct {
	items []any
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (list List) print() {
	fmt.Print("[")
	for i, item := range list.items {
		switch inner := item.(type) {
		case List:
			inner.print()
		case int:
			fmt.Print(inner)
		}
		if i < len(list.items)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
}

func is_number(c byte) bool {
	return c >= '0' && c <= '9'
}

func parse_number(s string, current *int) int {
	start := *current

	for is_number(s[*current]) {
		*current++
	}

	end := *current
	return aocutil.Atoi(s[start:end])
}

func parse_list(s string, current *int) List {
	list := List{items: make([]any, 0, 8)}
	for s[*current] != ']' {
		if s[*current] == '[' {
			*current++
			list.items = append(list.items, parse_list(s, current))
			*current++
		}

		if is_number(s[*current]) {
			list.items = append(list.items, parse_number(s, current))
		}

		if s[*current] == ',' {
			*current++
		}
	}

	return list
}

func parse(s string) List {
	start := 1
	return parse_list(s, &start)
}

const (
	COMPARE_EQUALS  = 0
	COMPARE_LESSER  = -1
	COMPARE_GREATER = 1
)

func compare(left, right List) int {
	if len(left.items) == 0 && len(right.items) == 0 {
		return COMPARE_EQUALS
	}

	lower := aocutil.Min(len(left.items), len(right.items))
	order := COMPARE_EQUALS
	for i := 0; i < lower; i++ {
		leftItem := left.items[i]
		rightItem := right.items[i]

		switch v1 := leftItem.(type) {
		case int:
			switch v2 := rightItem.(type) {
			case int:
				order = v1 - v2
			case List:
				order = compare(List{items: []any{v1}}, v2)
			}
		case List:
			switch v2 := rightItem.(type) {
			case int:
				order = compare(v1, List{items: []any{v2}})
			case List:
				order = compare(v1, v2)
			}
		}

		if order != COMPARE_EQUALS {
			break
		}
	}

	//if it is still equal then check whether one of them has still elements left
	//like being a subset of the other
	if order == COMPARE_EQUALS {
		if len(left.items) > len(right.items) {
			order = COMPARE_GREATER
		} else if len(left.items) < len(right.items) {
			order = COMPARE_LESSER
		}
	}

	return order
}

func main() {
	lines := make([]string, 0, 450)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		lines = append(lines, s)
	})

	packets := make([]List, 0, (len(lines)/3)*2)

	for i := 0; i < len(lines); i += 3 {
		packets = append(packets, parse(lines[i]))
		packets = append(packets, parse(lines[i+1]))
	}

	part1 := 0

	for i, j := 0, 0; i < len(packets); i += 2 {
		if compare(packets[i], packets[i+1]) < 0 {
			part1 += (j + 1)
		}
		j++
	}

	extraPacket1 := List{items: []any{List{items: []any{2}}}}
	extraPacket2 := List{items: []any{List{items: []any{6}}}}

	packets = append(packets, extraPacket1)
	packets = append(packets, extraPacket2)

	extraPacket1Index := 1
	extraPacket2Index := 1
	for i := range packets {
		if compare(packets[i], extraPacket1) < 0 {
			extraPacket1Index++
			extraPacket2Index++
		} else if compare(packets[i], extraPacket2) < 0 {
			extraPacket2Index++
		}
	}
	part2 := extraPacket1Index * extraPacket2Index

	aocutil.AOCFinish(part1, part2)
}
