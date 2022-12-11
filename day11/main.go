package main

import (
	"aocutil"
	"sort"
	"strings"
)

const (
	TOKEN_OLD = iota
	TOKEN_NUMBER
	TOKEN_ADD
	TOKEN_MUL
)

type token struct {
	ttype, value int
}

type monkey struct {
	items     []int
	operation []token
	test      int
	inspected int
	targets   []int
}

func parse(lines []string) monkey {
	m := monkey{}
	items := strings.Split(strings.Replace(lines[1][18:], " ", "", -1), ",")
	// 36 is the total number of items (counted manually) => less dynamic allocations
	m.items = make([]int, 0, 36)
	for _, item := range items {
		m.items = append(m.items, aocutil.Atoi(item))
	}

	operation := strings.Fields(lines[2][19:])
	m.operation = make([]token, 0, 3)

	for _, op := range operation {
		t := token{}
		switch op {
		case "old":
			t.ttype = TOKEN_OLD
		case "*":
			t.ttype = TOKEN_MUL
		case "+":
			t.ttype = TOKEN_ADD
		default:
			t.ttype = TOKEN_NUMBER
			t.value = aocutil.Atoi(op)
		}
		m.operation = append(m.operation, t)
	}

	m.test = aocutil.Atoi(lines[3][21:])
	m.targets = []int{
		aocutil.Atoi(lines[4][29:]),
		aocutil.Atoi(lines[5][30:]),
	}

	return m
}

func get_token_value(old int, t token) int {
	switch t.ttype {
	case TOKEN_OLD:
		return old
	case TOKEN_NUMBER:
		return t.value
	default:
		panic("invalid operand token")
	}
}

func do_operation(old int, tokens []token) int {
	op1 := get_token_value(old, tokens[0])
	op2 := get_token_value(old, tokens[2])
	switch tokens[1].ttype {
	case TOKEN_ADD:
		return op1 + op2
	case TOKEN_MUL:
		return op1 * op2
	default:
		panic("unknown operator token")
	}
}

func simulate(monkeys []monkey, rounds int, relief bool) int {
	//create test value which is dividable by all unique test values
	//would be least common multiple (LCM) but all test values are primes
	test := 1
	for i := range monkeys {
		test *= monkeys[i].test
	}

	for r := 0; r < rounds; r++ {
		for mi := range monkeys {
			m := &monkeys[mi]
			for i := range m.items {
				//inspect
				m.items[i] = do_operation(m.items[i], m.operation)
				m.inspected++

				//relief
				if relief {
					m.items[i] /= 3
				} else {
					m.items[i] = m.items[i] % test
				}

				//throw
				var m2 *monkey
				if m.items[i]%m.test == 0 {
					m2 = &monkeys[m.targets[0]]
				} else {
					m2 = &monkeys[m.targets[1]]
				}
				m2.items = append(m2.items, m.items[i])
			}

			//all items thrown. remove them
			m.items = m.items[:0]
		}
	}

	// sort in descending order
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	// multiply the first two (largest)
	return monkeys[0].inspected * monkeys[1].inspected
}

func main() {
	lines := make([]string, 0, 6)
	monkeys := make([]monkey, 0, 8)

	aocutil.FileReadAllLines("input.txt", func(s string) {
		if len(s) > 0 {
			lines = append(lines, s)
		} else {
			monkeys = append(monkeys, parse(lines))
			lines = lines[:0]
		}
	})
	monkeys = append(monkeys, parse(lines))

	//copy "deep" copy monkeys for part2 (only need to deep copy items)
	monkeysCopy := make([]monkey, len(monkeys))
	copy(monkeysCopy, monkeys)
	for i := range monkeysCopy {
		monkeysCopy[i].items = make([]int, len(monkeys[i].items))
		copy(monkeysCopy[i].items, monkeys[i].items)
	}

	part1 := simulate(monkeys, 20, true)
	part2 := simulate(monkeysCopy, 10000, false)

	aocutil.AOCFinish(part1, part2)
}
