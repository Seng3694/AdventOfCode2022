package main

import (
	"aocutil"
)

type vector struct {
	x, y int
}

type command struct {
	move   vector
	amount int
}

func direction_to_move(direction byte) vector {
	switch direction {
	case 'L':
		return vector{-1, 0}
	case 'U':
		return vector{0, -1}
	case 'R':
		return vector{1, 0}
	case 'D':
		return vector{0, 1}
	default:
		return vector{0, 0}
	}
}

func parse_command(line string) command {
	return command{
		move:   direction_to_move(line[0]),
		amount: aocutil.Atoi(line[2:]),
	}
}

func sign(value int) int {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

func main() {
	commands := make([]command, 0, 2000)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		commands = append(commands, parse_command(s))
	})

	//head at rope[0]
	rope := make([]vector, 10)

	part1 := make(map[vector]bool)
	part2 := make(map[vector]bool)

	part1[rope[0]] = true
	part2[rope[0]] = true

	//precalculated square roots. numbers from 0-8
	//longest distance:
	// . . .      h . .      h . .      h . .
	// h . .  UP  . . .  S1  1 . .  S2  1 . .
	// . 1 .      . 1 .      . . .      . 2 .
	// . . 2      . . 2      . . 2      . . .
	//there is a distance of two after S1 in each direction before "2" catches up
	//sqrt(2*2 + 2*2) = sqrt(8)
	sqrts := []int{0, 1, 1, 2, 2, 2, 2, 2, 2}

	for _, cmd := range commands {
		for step := 0; step < cmd.amount; step++ {
			rope[0].x += cmd.move.x
			rope[0].y += cmd.move.y

			for i := 1; i < len(rope); i++ {
				dx := rope[i-1].x - rope[i].x
				dy := rope[i-1].y - rope[i].y
				dist := sqrts[dx*dx+dy*dy]
				if dist > 1 {
					rope[i].x += sign(dx)
					rope[i].y += sign(dy)
					if i == 1 {
						part1[rope[1]] = true
					} else if i == 9 {
						part2[rope[9]] = true
					}
				}
			}
		}
	}

	aocutil.AOCFinish(len(part1), len(part2))
}
