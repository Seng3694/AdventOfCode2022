package main

import (
	"aocutil"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Line []Point

const (
	TILE_EMPTY uint8 = iota
	TILE_SOLID
	TILE_SAND
	TILE_SOURCE
)

const (
	STATE_FALLING = iota
	STATE_LEFT
	STATE_RIGHT
	STATE_SETTLED
	STATE_DONE
)

func parse_point(s string) Point {
	coords := strings.Split(s, ",")
	return Point{
		x: aocutil.Atoi(coords[0]),
		y: aocutil.Atoi(coords[1]),
	}
}

func parse(s string) Line {
	points := strings.Split(strings.Replace(s, " ", "", -1), "->")
	line := make(Line, len(points))
	for p := range points {
		line[p] = parse_point(points[p])
	}
	return line
}

func find_bottom(lines []Line) int {
	bottom := 0
	for _, line := range lines {
		for _, point := range line {
			if point.y > bottom {
				bottom = point.y
			}
		}
	}
	return bottom
}

func translate(point *Point, translation Point) {
	point.x -= translation.x
	point.y -= translation.y
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

func get_type(caveMap []uint8, w, h int, coord Point) uint8 {
	return caveMap[coord.y*w+coord.x]
}

func draw_point(caveMap *[]uint8, w, h int, coord Point, tiletype uint8) {
	(*caveMap)[coord.y*w+coord.x] = tiletype
}

func draw_line(caveMap *[]uint8, w, h int, line Line) {
	draw_point(caveMap, w, h, line[0], TILE_SOLID)
	for i := 0; i < len(line)-1; i++ {
		pt1 := line[i]
		pt2 := line[i+1]

		draw_point(caveMap, w, h, pt2, TILE_SOLID)
		x := sign(pt2.x - pt1.x)
		y := sign(pt2.y - pt1.y)

		for j := pt1.x; j != pt2.x; j += x {
			draw_point(caveMap, w, h, Point{j, pt1.y}, TILE_SOLID)
		}
		for j := pt1.y; j != pt2.y; j += y {
			draw_point(caveMap, w, h, Point{pt1.x, j}, TILE_SOLID)
		}
	}
}

func draw_lines(caveMap *[]uint8, w, h int, lines []Line) {
	for _, line := range lines {
		draw_line(caveMap, w, h, line)
	}
}

func print_tile(tiletype uint8) {
	switch tiletype {
	case TILE_EMPTY:
		fmt.Print(".")
	case TILE_SOLID:
		fmt.Print("#")
	case TILE_SAND:
		fmt.Print("o")
	case TILE_SOURCE:
		fmt.Print("+")
	}
}

func print_map(caveMap []uint8, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			print_tile(caveMap[y*w+x])
		}
		fmt.Println()
	}
}

func is_in_bounds(x, y, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

func simulate(caveMap []uint8, w, h int, sandSource Point) int {
	sand := sandSource
	state := STATE_FALLING
	count := 0
	done := false

	for !done {
		switch state {
		case STATE_FALLING:
			next := Point{sand.x, sand.y + 1}
			if !is_in_bounds(next.x, next.y, w, h) {
				state = STATE_DONE
				break
			}
			switch get_type(caveMap, w, h, next) {
			case TILE_EMPTY:
				sand = next
			case TILE_SOLID:
				fallthrough
			case TILE_SAND:
				state = STATE_LEFT
			}
		case STATE_LEFT:
			next := Point{sand.x - 1, sand.y + 1}
			if !is_in_bounds(next.x, next.y, w, h) {
				state = STATE_DONE
				break
			}
			switch get_type(caveMap, w, h, next) {
			case TILE_EMPTY:
				sand = next
				state = STATE_FALLING
			case TILE_SOLID:
				fallthrough
			case TILE_SAND:
				state = STATE_RIGHT
			}
		case STATE_RIGHT:
			next := Point{sand.x + 1, sand.y + 1}
			if !is_in_bounds(next.x, next.y, w, h) {
				state = STATE_DONE
				break
			}
			switch get_type(caveMap, w, h, next) {
			case TILE_EMPTY:
				sand = next
				state = STATE_FALLING
			case TILE_SOLID:
				fallthrough
			case TILE_SAND:
				state = STATE_SETTLED
			}
		case STATE_SETTLED:
			draw_point(&caveMap, w, h, sand, TILE_SAND)
			count++
			if sand == sandSource {
				state = STATE_DONE
			} else {
				state = STATE_FALLING
			}
			sand = sandSource
		case STATE_DONE:
			done = true
		}
	}

	return count
}

func part1(lines []Line) int {
	sandSource := Point{500, 0}

	bottom := find_bottom(lines)
	h := bottom + 1
	w := h * 2
	caveMap := make([]uint8, w*h)
	origin := Point{sandSource.x - (w / 2), 0}
	for i := range lines {
		for j := range lines[i] {
			translate(&lines[i][j], origin)
		}
	}

	draw_lines(&caveMap, w, h, lines)

	translate(&sandSource, origin)
	draw_point(&caveMap, w, h, sandSource, TILE_SOURCE)

	return simulate(caveMap, w, h, sandSource)
}

func part2(lines []Line) int {
	sandSource := Point{500, 0}

	bottom := find_bottom(lines)
	h := bottom + 1 + 2 //+1 to include the last one, +2 for part 2
	w := h * 2
	caveMap := make([]uint8, w*h)
	origin := Point{sandSource.x - (w / 2), 0}
	for i := range lines {
		for j := range lines[i] {
			translate(&lines[i][j], origin)
		}
	}

	draw_lines(&caveMap, w, h, lines)
	//draw extra line at the bottom
	draw_line(&caveMap, w, h, Line{Point{0, h - 1}, Point{w - 1, h - 1}})

	translate(&sandSource, origin)
	draw_point(&caveMap, w, h, sandSource, TILE_SOURCE)

	return simulate(caveMap, w, h, sandSource)
}

func main() {
	lines := make([]Line, 0, 128)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		lines = append(lines, parse(s))
	})

	linesCopy := make([]Line, len(lines))
	for i := range lines {
		linesCopy[i] = make(Line, len(lines[i]))
		copy(linesCopy[i], lines[i])
	}

	p1 := part1(lines)
	p2 := part2(linesCopy)

	aocutil.AOCFinish(p1, p2)
}
