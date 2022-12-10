package main

import (
	"aocutil"
	"fmt"
	"strings"
)

type instruction struct {
	cycles, operand int
}

const (
	NOOP = 1
	ADDX = 2
)

func parse(s string) instruction {
	fields := strings.Fields(s)
	op := NOOP
	operand := 0
	if fields[0] == "addx" {
		op = ADDX
		operand = aocutil.Atoi(fields[1])
	}
	return instruction{
		cycles:  op,
		operand: operand,
	}
}

func main() {
	instructions := make([]instruction, 0, 100)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		instructions = append(instructions, parse(s))
	})

	x := 1
	cycles := 1

	signals := []int{20, 60, 100, 140, 180, 220, 260}
	nextSignalIndex := 0

	crt := []byte(`........................................
........................................
........................................
........................................
........................................
........................................
`)

	//scanline position
	scx := 0
	scy := 0

	part1 := 0

	for _, instr := range instructions {
		for i := 0; i < instr.cycles; i++ {
			//check signal
			if cycles == signals[nextSignalIndex] {
				part1 += cycles * x
				nextSignalIndex++
			}

			//draw sprite
			if scx >= x-1 && scx <= x+1 {
				crt[scy*41+scx] = '#'
			}

			//advance scanline
			scx++
			if scx == 40 {
				scx = 0
				scy++
			}
			cycles++
		}

		x += instr.operand
	}

	aocutil.AOCFinish(fmt.Sprint(part1), "\n"+string(crt))
}
