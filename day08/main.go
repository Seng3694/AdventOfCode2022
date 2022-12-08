package main

import (
	"aocutil"
	"math"
)

type vector struct {
	x, y int
}

var (
	LEFT   = vector{-1, 0}
	TOP    = vector{0, -1}
	RIGHT  = vector{1, 0}
	BOTTOM = vector{0, 1}
)

type scenic_overview struct {
	score   int
	visible bool
}

func check_direction(x, y int, move, to vector, trees [][]byte) scenic_overview {
	tree := trees[y][x]

	current := vector{x, y}
	visible := true
	score := 0
	for current != to && visible {
		current.x += move.x
		current.y += move.y
		score++
		visible = trees[current.y][current.x] < tree
	}

	return scenic_overview{
		score,
		visible,
	}
}

func get_scenic_overview(x, y, w, h int, trees [][]byte) scenic_overview {
	score := 1
	visible := false
	overviews := []scenic_overview{
		check_direction(x, y, LEFT, vector{0, y}, trees),
		check_direction(x, y, TOP, vector{x, 0}, trees),
		check_direction(x, y, RIGHT, vector{w - 1, y}, trees),
		check_direction(x, y, BOTTOM, vector{x, h - 1}, trees),
	}
	for _, o := range overviews {
		score *= o.score
		visible = visible || o.visible
	}

	return scenic_overview{
		score,
		visible,
	}
}

func main() {
	trees := make([][]byte, 0, 32)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		trees = append(trees, []byte(s))
	})

	width := len(trees[0])
	height := len(trees)

	//edges minus corners which are counted twice
	visible := width*2 + height*2 - 4
	score := math.MinInt
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			overview := get_scenic_overview(x, y, width, height, trees)
			if overview.visible {
				visible++
			}
			if score < overview.score {
				score = overview.score
			}
		}
	}

	aocutil.AOCFinish(visible, score)
}
