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

type scenic_overview struct {
	score   int
	visible bool
}

func get_scenic_overview(x, y, w, h int, trees [][]byte) scenic_overview {
	tree := trees[y][x]

	visibleFromTop := true
	topScore := 0
	for top := y - 1; top >= 0; top-- {
		if top < 0 {
			break
		}
		topScore++
		if trees[top][x] >= tree {
			visibleFromTop = false
			break
		}
	}

	visibleFromLeft := true
	leftScore := 0
	for left := x - 1; left >= 0; left-- {
		if left < 0 {
			break
		}
		leftScore++
		if trees[y][left] >= tree {
			visibleFromLeft = false
			break
		}
	}

	visibleFromRight := true
	rightScore := 0
	for right := x + 1; right < w; right++ {
		rightScore++
		if trees[y][right] >= tree {
			visibleFromRight = false
			break
		}
	}

	visibleFromBottom := true
	bottomScore := 0
	for bottom := y + 1; bottom < h; bottom++ {
		bottomScore++
		if trees[bottom][x] >= tree {
			visibleFromBottom = false
			break
		}
	}
	overview := scenic_overview{
		score:   topScore * leftScore * rightScore * bottomScore,
		visible: visibleFromTop || visibleFromLeft || visibleFromRight || visibleFromBottom,
	}
	return overview
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
	maxScore := math.MinInt
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			overview := get_scenic_overview(x, y, width, height, trees)
			if overview.visible {
				visible++
			}
			if maxScore < overview.score {
				maxScore = overview.score
			}
		}
	}

	aocutil.AOCFinish(visible, maxScore)
}
