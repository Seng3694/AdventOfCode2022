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

func main() {
	scorePart1 := 0
	scorePart2 := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		moves := strings.Split(s, " ")
		scorePart1 += calculate_score(moves[0][0], moves[1][0])

		switch moves[0][0] {
		case 'A':
			switch moves[1][0] {
			case 'X':
				scorePart2 += 3
			case 'Y':
				scorePart2 += 1 + 3
			case 'Z':
				scorePart2 += 2 + 6
			}
		case 'B':
			switch moves[1][0] {
			case 'X':
				scorePart2 += 1
			case 'Y':
				scorePart2 += 2 + 3
			case 'Z':
				scorePart2 += 3 + 6
			}
		case 'C':
			switch moves[1][0] {
			case 'X':
				scorePart2 += 2
			case 'Y':
				scorePart2 += 3 + 3
			case 'Z':
				scorePart2 += 1 + 6
			}
		}
	})

	aocutil.AOCFinish(scorePart1, scorePart2)
}
