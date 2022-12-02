package main

import (
	"aocutil"
	"strings"
)

func wrap(value, min, max int) int {
	if value < min {
		return (max + 1) - (min - value)
	} else if value > max {
		return (min - 1) + (value - max)
	} else {
		return value
	}
}

func main() {
	scorePart1 := 0
	scorePart2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		moves := strings.Split(s, " ")
		opponent := int(moves[0][0] - 'A' + 1)
		player := int(moves[1][0] - 'X' + 1)

		// PART 1
		// both between 1 and 3 with 1 being rock, 2 being paper, 3 being scissors
		// the higher one beats the lower one (dist 1)
		// if dist is 2 then the lower one beats the higher one (wraps around)

		// rock vs paper        1 - 2 = -1 win
		// paper vs scissors    2 - 3 = -1 win
		// scissors vs rock     3 - 1 = 2 win

		// rock vs scissors     1 - 3 = -2 loss
		// paper vs rock        2 - 1 = 1 loss
		// scissors vs paper    3 - 2 = 1 loss

		// draw when 0

		scorePart1 += player //add player value
		distance := wrap(opponent-player, 1, 3)
		if distance == 2 {
			scorePart1 += 6
		} else if distance != 1 {
			scorePart1 += 3
		}

		// PART 2
		// 1:rock <- 2:paper <- 3:scissors
		// one number higher is always the one who beats the lower one
		// losing is always in the other way around
		// so opponent+1 is the winning move
		// opponent-1 is the losing move
		// wrap if it's below 1 or above 3

		switch player {
		case 1: //X lose
			player = wrap(opponent-1, 1, 3)
			scorePart2 += player
		case 2: //Y draw
			player = opponent
			scorePart2 += player + 3
		case 3: //Z win
			player = wrap(opponent+1, 1, 3)
			scorePart2 += player + 6
		}
	})

	//14163
	//12091
	aocutil.AOCFinish(scorePart1, scorePart2)
}
