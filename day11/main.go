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

type node struct {
	value, monkey int
	next          *node
}

type item struct {
	index, value, monkey int
}

type monkey struct {
	operation []token
	targets   []int
	test      int
}

var (
	monkeys      []monkey
	combinedTest int
)

func parse(lines []string, items *[]item) {
	m := monkey{}
	itemsString := strings.Split(strings.Replace(lines[1][18:], " ", "", -1), ",")

	for _, i := range itemsString {
		*items = append(*items, item{
			index:  len(*items),
			value:  aocutil.Atoi(i),
			monkey: aocutil.Atoi(lines[0][7:8]),
		})
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
	monkeys = append(monkeys, m)
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

func simulate(items []item, rounds int, relief bool) int {
	lookup := make([]map[int]*node, len(monkeys))
	for i := range lookup {
		lookup[i] = make(map[int]*node, 32)
	}

	visits := make([]int, len(monkeys))

	for _, item := range items {
		r := 0
		for r < rounds {
			oldNode, foundOld := lookup[item.monkey][item.value]
			if !foundOld || (foundOld && oldNode.next == nil) {
				// inspect
				oldValue := item.value
				oldMonkey := item.monkey
				item.value = do_operation(oldValue, monkeys[item.monkey].operation)
				visits[item.monkey]++

				// relief
				if relief {
					item.value /= 3
				} else {
					item.value = item.value % combinedTest
				}

				// "throw" to next monkey
				if item.value%monkeys[item.monkey].test == 0 {
					item.monkey = monkeys[item.monkey].targets[0]
				} else {
					item.monkey = monkeys[item.monkey].targets[1]
				}

				//check if next node exists. if so wire it up properly
				newNode := oldNode
				nextNode, foundNext := lookup[item.monkey][item.value]
				if newNode == nil {
					newNode = &node{value: oldValue, monkey: oldMonkey, next: nil}
				}
				if foundNext {
					newNode.next = nextNode
				}
				lookup[oldMonkey][oldValue] = newNode

				//monkeys throws from 0-N.
				//so if monkey 2 throws to monkey 3 then monkey 3 throws in the same round
				//if monkey 5 throws to monkey 2 then monkey 2 will throw in the next round
				if item.monkey < oldMonkey {
					r++
				}
			} else {
				// quick lookup due to repeating throwing patterns
				current := oldNode
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
	monkeys = make([]monkey, 0, 8)
	items := make([]item, 0, 36)

	aocutil.FileReadAllLines("input.txt", func(s string) {
		if len(s) > 0 {
			lines = append(lines, s)
		} else {
			parse(lines, &items)
			lines = lines[:0]
		}
	})
	parse(lines, &items)

	combinedTest = 1
	for _, m := range monkeys {
		combinedTest *= m.test
	}

	itemsCopy := make([]item, len(items))
	copy(itemsCopy, items)

	part1 := simulate(items, 20, true)
	part2 := simulate(itemsCopy, 10000, false)

	aocutil.AOCFinish(part1, part2)
}
