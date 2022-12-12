package main

import (
	"aocutil"
	"math"
)

type Vector struct {
	x, y int
}

var (
	LEFT  = Vector{-1, 0}
	UP    = Vector{0, -1}
	RIGHT = Vector{1, 0}
	DOWN  = Vector{0, 1}
)

func (v Vector) Add(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y}
}

type Queue[T any] []T

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() {
	*q = (*q)[1:len(*q)]
}

func (q *Queue[T]) DequeueMany(count int) {
	*q = (*q)[count:]
}

func (q *Queue[T]) Clear() {
	*q = (*q)[:0]
}

type BfsData struct {
	pos  Vector
	dist int
}

func is_in_bounds(pos Vector, w, h int) bool {
	return pos.x >= 0 && pos.x < w && pos.y >= 0 && pos.y < h
}

func is_valid_move(from, to Vector, w, h int, hm [][]int8) (Vector, bool) {
	//you can go down as far as you want (negative numbers) but only go up by 1
	return to,
		is_in_bounds(to, w, h) && (hm[to.y][to.x]-hm[from.y][from.x]) <= 1
}

func main() {
	hm := make([][]int8, 0, 41)
	start := Vector{}
	dest := Vector{}
	h := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		arr := make([]int8, len(s))
		for x := range s {
			if s[x] == 'S' {
				start.x = x
				start.y = h
				arr[x] = 'a'
			} else if s[x] == 'E' {
				dest.x = x
				dest.y = h
				arr[x] = 'z'
			} else {
				arr[x] = int8(s[x])
			}
		}
		hm = append(hm, arr)
		h++
	})
	w := len(hm[0])

	starts := make([]Vector, 0, 32)
	//collect all starting positions 'a' at the edges of the map
	for y := 0; y < h; y++ {
		if y != 0 && y != h-1 {
			if hm[y][0] == 'a' {
				starts = append(starts, Vector{0, y})
			}
			if hm[y][w-1] == 'a' {
				starts = append(starts, Vector{w - 1, y})
			}
		} else {
			for x := 0; x < w; x++ {
				if hm[y][x] == 'a' {
					starts = append(starts, Vector{x, y})
				}
			}
		}
	}

	queue := make(Queue[BfsData], 0, w*h)
	visited := make(map[Vector]int, w*h)
	part1 := 0
	part2 := math.MaxInt

	for _, s := range starts {
		queue.Enqueue(BfsData{pos: s})
		visited[start] = 0

	done:
		for len(queue) > 0 {
			currentLength := len(queue)
			for i := 0; i < currentLength; i++ {
				current := queue[i]
				if current.pos == dest {
					if part2 > current.dist {
						part2 = current.dist
					}
					if s == start {
						part1 = current.dist
					}
					queue.Clear()
					for v := range visited {
						delete(visited, v)
					}
					break done
				}

				positions := []Vector{
					current.pos.Add(LEFT),
					current.pos.Add(UP),
					current.pos.Add(RIGHT),
					current.pos.Add(DOWN),
				}

				for _, next := range positions {
					if pos, yes := is_valid_move(current.pos, next, w, h, hm); yes {
						if _, yes := visited[pos]; !yes {
							visited[pos] = current.dist + 1
							queue.Enqueue(BfsData{pos: pos, dist: current.dist + 1})
						}
					}
				}
			}
			queue.DequeueMany(currentLength)
		}
	}

	aocutil.AOCFinish(part1, part2)
}
