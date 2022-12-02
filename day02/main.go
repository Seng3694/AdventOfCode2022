package main

import (
	"aocutil"
	"strings"
)

func calculate_score(opponent, response byte) int {
	score := 0
	switch response {
	case 'X':
		score += 1
	case 'Y':
		score += 2
	case 'Z':
		score += 3
	}

	switch opponent {
	case 'A':
		switch response {
		case 'X':
			score += 3
		case 'Y':
			score += 6
		case 'Z':
			score += 0
		}
	case 'B':
		switch response {
		case 'X':
			score += 0
		case 'Y':
			score += 3
		case 'Z':
			score += 6
		}
	case 'C':
		switch response {
		case 'X':
			score += 6
		case 'Y':
			score += 0
		case 'Z':
			score += 3
		}
	}

	return score
}

//rock <- paper <- scissors

func wrap(value, min, max int) int {
	if value < min {
		return max
	} else if value > max {
		return min
	} else {
		return value
	}
}

func main() {
	scorePart1 := 0
	scorePart2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		moves := strings.Split(s, " ")
		first := moves[0][0]
		second := moves[1][0]

		opponent := int(first - 'A' + 1)
		player := int(second - 'X' + 1)

		//both between 1 and 3 with 1 being rock, 2 being paper, 3 being scissors
		//the higher one beats the lower one (dist 1)
		//if dist is 2 then the lower one beast the higher one (wraps around)

		//rock vs paper        1 - 2 = -1 win
		//paper vs scissors    2 - 3 = -1 win
		//scissors vs rock     3 - 1 = 2 win

		//rock vs scissors     1 - 3 = -2 loss
		//paper vs rock        2 - 1 = 1 loss
		//scissors vs paper    3 - 2 = 1 loss

		dist := opponent - player
		if dist == -2 || dist == 1 { //loss
			scorePart1 += player
		} else if dist == 2 || dist == -1 { //win
			scorePart1 += player + 6
		} else { //draw
			scorePart1 += player + 3
		}

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
