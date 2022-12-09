package main

import (
	"aocutil"
	"strings"
)

type vector struct {
	x, y int
}

const (
	LEFT = iota
	UP
	RIGHT
	DOWN
)

type command struct {
	direction, amount int
}

func direction_to_int(direction string) int {
	switch direction {
	case "L":
		return LEFT
	case "U":
		return UP
	case "R":
		return RIGHT
	case "D":
		return DOWN
	default:
		return -1
	}
}

func parse_command(line string) command {
	split := strings.Split(line, " ")
	return command{
		direction: direction_to_int(split[0]),
		amount:    aocutil.Atoi(split[1]),
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

func move(position *vector, direction int) {
	switch direction {
	case LEFT:
		position.x--
	case UP:
		position.y++
	case RIGHT:
		position.x++
	case DOWN:
		position.y--
	}
}

func main() {
	commands := make([]command, 0, 2000)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		commands = append(commands, parse_command(s))
	})

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

	for _, c := range commands {
		for s := 0; s < c.amount; s++ {
			move(&rope[0], c.direction)

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
