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

type item_node struct {
	value, monkey int
	next          *item_node
}

var (
	globalItemIndex = 0
)

type item struct {
	index, value, monkey int
}

type monkey struct {
	items     []item
	operation []token
	test      int
	targets   []int
}

func parse(lines []string) monkey {
	m := monkey{}
	items := strings.Split(strings.Replace(lines[1][18:], " ", "", -1), ",")
	// 36 is the total number of items (counted manually) => less dynamic allocations
	m.items = make([]item, 0, 36)
	for _, i := range items {
		m.items = append(m.items, item{
			index:  globalItemIndex,
			value:  aocutil.Atoi(i),
			monkey: aocutil.Atoi(lines[0][7:8]),
		})
		globalItemIndex++
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
	itemMap := make([]map[int]*item_node, len(monkeys))
	for i := range itemMap {
		itemMap[i] = make(map[int]*item_node, 32)
	}

	items := make([]item, 0, globalItemIndex)
	for _, m := range monkeys {
		items = append(items, m.items...)
	}

	visits := make([]int, len(monkeys))

	test := 1
	for i := range monkeys {
		test *= monkeys[i].test
	}

	for _, item := range items {
		r := 0
		for r < rounds {
			value, found := itemMap[item.monkey][item.value]
			if !found {
				old := item.value
				oldMonkey := item.monkey
				item.value = do_operation(old, monkeys[item.monkey].operation)
				visits[item.monkey]++

				if relief {
					item.value /= 3
				} else {
					item.value = item.value % test
				}

				m2 := 0
				if item.value%monkeys[item.monkey].test == 0 {
					m2 = monkeys[item.monkey].targets[0]
				} else {
					m2 = monkeys[item.monkey].targets[1]
				}
				item.monkey = m2

				value, found = itemMap[item.monkey][item.value]
				if !found {
					thisNode := &item_node{value: old, monkey: oldMonkey, next: nil}
					itemMap[oldMonkey][old] = thisNode
				} else {
					thisNode := &item_node{value: old, monkey: oldMonkey, next: value}
					itemMap[oldMonkey][old] = thisNode
				}
				if item.monkey < oldMonkey {
					r++
				}
			} else if found && value.next == nil {
				old := item.value
				oldMonkey := item.monkey
				item.value = do_operation(old, monkeys[item.monkey].operation)
				visits[item.monkey]++

				if relief {
					item.value /= 3
				} else {
					item.value = item.value % test
				}

				m2 := 0
				if item.value%monkeys[item.monkey].test == 0 {
					m2 = monkeys[item.monkey].targets[0]
				} else {
					m2 = monkeys[item.monkey].targets[1]
				}
				item.monkey = m2

				value2, found2 := itemMap[item.monkey][item.value]
				if found2 {
					value.next = value2
				}
				if item.monkey < oldMonkey {
					r++
				}
			} else {
				current := value
				for current.next != nil && r < rounds {
					oldMonkey := current.monkey
					current = current.next
					item.value = current.value
					item.monkey = current.monkey
					if item.monkey < oldMonkey {
						r++
					}
					visits[oldMonkey]++
				}
			}
		}
	}

	// sort in descending order
	sort.SliceStable(visits, func(i, j int) bool {
		return visits[i] > visits[j]
	})

	return visits[0] * visits[1]
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
		monkeysCopy[i].items = make([]item, len(monkeys[i].items))
		copy(monkeysCopy[i].items, monkeys[i].items)
	}

	part1 := simulate(monkeys, 20, true)
	part2 := simulate(monkeysCopy, 10000, false)

	aocutil.AOCFinish(part1, part2)
}
