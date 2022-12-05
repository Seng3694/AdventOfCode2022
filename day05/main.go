package main

import (
	"aocutil"
	"strings"
)

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Pop() T {
	value := s.Peek()
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *Stack[T]) PushMany(values []T) {
	*s = append(*s, values...)
}

func (s *Stack[T]) PopMany(amount int) []T {
	end := len(*s)
	start := end - amount
	values := (*s)[start:end]
	*s = (*s)[:start]
	return values
}

func parse_arrangements(arrangements []string) []Stack[byte] {
	//last row contains information how many stacks there are without string splitting etc
	info := arrangements[len(arrangements)-1]
	//only works with single digits which the input has!
	crateCount := aocutil.Atoi(info[len(info)-2 : len(info)-1])

	//remove info from array
	arrangements = arrangements[:len(arrangements)-1]
	crates := make([]Stack[byte], crateCount)
	for i := 0; i < crateCount; i++ {
		crates[i] = make(Stack[byte], 0, len(arrangements))
	}

	//iterate arrangements in reverse order and push elements onto the stack
	for i := len(arrangements) - 1; i >= 0; i-- {
		if i == -1 {
			break
		}

		// each stack has a length of 3 + 1 for the space in between but not for the last one
		for j, s := 0, 0; j < len(arrangements[i]); j += 3 {
			if arrangements[i][j+1] != ' ' {
				crates[s].Push(arrangements[i][j+1])
			}
			if j-3 < len(arrangements[i]) {
				j++
			}
			s++
		}
	}

	return crates
}

type Instruction struct {
	amount, from, to int
}

func parse_instructions(instructions []string) []Instruction {
	instr := make([]Instruction, len(instructions))

	for i := range instructions {
		split := strings.Split(instructions[i], " ")
		//indices are starting at 1 so subtract 1 for 0 based indices
		instr[i] = Instruction{
			amount: aocutil.Atoi(split[1]),
			from:   aocutil.Atoi(split[3]) - 1,
			to:     aocutil.Atoi(split[5]) - 1,
		}
	}

	return instr
}

func run_crate_mover_9000(crates []Stack[byte], instructions []Instruction) {
	for _, instr := range instructions {
		for i := 0; i < instr.amount; i++ {
			crates[instr.to].Push(crates[instr.from].Pop())
		}
	}
}

func run_crate_mover_9001(crates []Stack[byte], instructions []Instruction) {
	for _, instr := range instructions {
		crates[instr.to].PushMany(crates[instr.from].PopMany(instr.amount))
	}
}

func read_top_elements(stacks []Stack[byte]) string {
	result := make([]byte, len(stacks))
	for i := range stacks {
		result[i] = stacks[i].Peek()
	}
	return string(result)
}

func copy_crates(crates []Stack[byte]) []Stack[byte] {
	output := make([]Stack[byte], len(crates))
	for i := range crates {
		output[i] = make(Stack[byte], len(crates[i]))
		copy(output[i], crates[i])
	}
	return output
}

func main() {
	arrangementStrings := make([]string, 0, 10)
	instructionsStrings := make([]string, 0, 500)
	readingArrangements := true
	aocutil.FileReadAllLines("input.txt", func(s string) {
		if len(s) == 0 {
			readingArrangements = false
		} else if !readingArrangements {
			instructionsStrings = append(instructionsStrings, s)
		} else {
			arrangementStrings = append(arrangementStrings, s)
		}
	})

	crates1 := parse_arrangements(arrangementStrings)
	crates2 := copy_crates(crates1)

	instructions := parse_instructions(instructionsStrings)
	run_crate_mover_9000(crates1, instructions)
	run_crate_mover_9001(crates2, instructions)

	aocutil.AOCFinish(read_top_elements(crates1), read_top_elements(crates2))
}
