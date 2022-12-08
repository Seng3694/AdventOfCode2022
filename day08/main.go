package main

import (
	"aocutil"
	"math"
)

const (
	LEFT = iota
	TOP
	RIGHT
	BOTTOM
)

type vector struct {
	x, y int
}

func check_direction(x, y, w, h, direction int, trees [][]byte) (int, bool) {
	tree := trees[y][x]
	move := vector{}
	to := vector{x, y}

	switch direction {
	case LEFT:
		move.x = -1
		to.x = 0
	case TOP:
		move.y = -1
		to.y = 0
	case RIGHT:
		move.x = 1
		to.x = w - 1
	case BOTTOM:
		move.y = 1
		to.y = h - 1
	}

	current := vector{x, y}
	visible := true
	score := 0
	for {
		current.x += move.x
		current.y += move.y
		score++
		if trees[current.y][current.x] >= tree {
			visible = false
			break
		}
		if current == to {
			break
		}
	}

	return score, visible
}

func get_scenic_overview(x, y, w, h int, trees [][]byte) (int, bool) {
	score := 1
	visible := false
	for i := 0; i < 4; i++ {
		s, v := check_direction(x, y, w, h, i, trees)
		score *= s
		if v {
			visible = true
		}
	}
	return score, visible
}

func main() {
	trees := make([][]byte, 0, 32)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		trees = append(trees, []byte(s))
	})

	width := len(trees[0])
	height := len(trees)

	//edges minus corners which are count twice
	visible := width*2 + height*2 - 4
	score := math.MinInt
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			s, v := get_scenic_overview(x, y, width, height, trees)
			if v {
				visible++
			}
			if score < s {
				score = s
			}
		}
	}

	aocutil.AOCFinish(visible, score)
}
