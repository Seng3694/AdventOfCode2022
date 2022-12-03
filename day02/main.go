package main

import (
	"aocutil"
	"strings"
)

// assumes that the difference between min-value or value-max is smaller than max-min
func wrap(value, min, max int) int {
	if value < min {
		return (max + 1) - (min - value)
	} else if value > max {
		return (min - 1) + (value - max)
	} else {
		return value
	}
}

func calculate_score(opponent, player int) int {
	score := player + 1
	distance := wrap(opponent-player, 0, 2)
	if distance == 2 { //win
		score += 6
	} else if distance == 0 { //draw
		score += 3
	}
	return score
}

func main() {
	part1 := 0
	part2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		moves := strings.Split(s, " ")
		opponent := int(moves[0][0] - 'A')
		player := int(moves[1][0] - 'X')

		// PART 1
		// both between 0 and 2 with 0 being rock, 1 being paper, 2 being scissors
		// the higher one beats the lower one (dist -1 => wraps around to 2)
		// if dist is +1 then the lower one beats the higher one (dist -2 => wraps around to +1)

		// rock vs paper        0 - 1 = -1 win
		// paper vs scissors    1 - 2 = -1 win
		// scissors vs rock     2 - 0 = 2 win

		// rock vs scissors     0 - 2 = -2 loss
		// paper vs rock        1 - 0 = 1 loss
		// scissors vs paper    2 - 1 = 1 loss

		// draw when 0

		part1 += calculate_score(opponent, player)

		// PART 2
		// 0:rock <- 1:paper <- 2:scissors
		// one number higher is always the one who beats the lower one
		// losing is always in the other way around
		// so opponent+1 is the winning move
		// opponent-1 is the losing move
		// wrap if it's below 0 or above 2
		switch player {
		case 0: //X lose
			player = wrap(opponent-1, 0, 2)
		case 1: //Y draw
			player = opponent
		case 2: //Z win
			player = wrap(opponent+1, 0, 2)
		}

		part2 += calculate_score(opponent, player)
	})

	aocutil.AOCFinish(part1, part2)
}
