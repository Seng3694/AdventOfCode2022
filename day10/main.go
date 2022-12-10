package main

import (
	"aocutil"
	"fmt"
	"strings"
)

const (
	//values are also used for cycle timers
	NOOP = 1
	ADDX = 2
)

func main() {
	instructions := make([]int8, 0, 250)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		fields := strings.Fields(s)
		if fields[0][0] == 'n' { //can only be noop
			instructions = append(instructions, NOOP)
		} else { //because there are only 2 operation it has to be addx
			instructions = append(instructions, ADDX)
			instructions = append(instructions, int8(aocutil.Atoi(fields[1])))
		}
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

	for i := 0; i < len(instructions); i++ {
		instr := instructions[i]
		for c := int8(0); c < instr; c++ {
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

		if instr == ADDX {
			i++
			x += int(instructions[i])
		}
	}

	aocutil.AOCFinish(fmt.Sprint(part1), "\n"+string(crt))
}
