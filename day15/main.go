package main

import (
	"aocutil"
	"regexp"
	"sort"
)

type Vector struct {
	x, y int
}

type Range struct {
	left, right int
}

type Sensor struct {
	position, closestBeacon Vector
	radius                  int
}

var (
	numericRegex = regexp.MustCompile(`[\d-]+`)
)

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func distance(start, end Vector) int {
	return abs(start.x-end.x) + abs(start.y-end.y)
}

func parse(s string) Sensor {
	numbers := numericRegex.FindAllString(s, -1)

	sensor := Sensor{
		position: Vector{
			aocutil.Atoi(numbers[0]),
			aocutil.Atoi(numbers[1]),
		},
		closestBeacon: Vector{
			aocutil.Atoi(numbers[2]),
			aocutil.Atoi(numbers[3]),
		},
	}

	sensor.radius = distance(sensor.position, sensor.closestBeacon)

	return sensor
}

func get_ranges(sensors []Sensor, y int, ranges *[]Range) {
	for _, s := range sensors {
		yDistance := abs(y - s.position.y)
		if yDistance < s.radius {
			dist := abs(s.radius - yDistance)
			r := Range{s.position.x - dist, s.position.x + dist}
			(*ranges) = append((*ranges), r)
		}
	}
}

func merge_ranges(ranges []Range, output *[]Range) {
	(*output) = append((*output), ranges[0])
	for _, r2 := range ranges {
		r1 := (*output)[len((*output))-1]
		if r1.right >= r2.left && r2.right > r1.right {
			r1.right = r2.right
			(*output)[len((*output))-1] = r1
		} else if r1.right < r2.left {
			(*output) = append((*output), r2)
		}
	}
}

func part1(sensors []Sensor) int {
	y := 2000000
	ranges := make([]Range, 0, len(sensors))
	get_ranges(sensors, 2000000, &ranges)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].left < ranges[j].left
	})

	mergedRanges := make([]Range, 0, len(ranges))
	merge_ranges(ranges, &mergedRanges)

	solution := 0
	for _, r := range mergedRanges {
		solution += (r.right - r.left) + 1
	}
	beacons := make(map[Vector]bool)
	for _, s := range sensors {
		beacons[s.closestBeacon] = true
	}
	for k := range beacons {
		if k.y == y {
			solution--
		}
	}
	return solution
}

func part2(sensors []Sensor) int {
	solution := 0
	ranges := make([]Range, 0, len(sensors))
	mergedRanges := make([]Range, 0, 2)

	for y := 0; y <= 4000000; y++ {
		ranges := ranges[:0]
		get_ranges(sensors, y, &ranges)

		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].left < ranges[j].left
		})

		mergedRanges = mergedRanges[:0]
		merge_ranges(ranges, &mergedRanges)

		if len(mergedRanges) == 2 {
			solution = (mergedRanges[0].right+1)*4000000 + y
			break
		}
	}
	return solution
}

func main() {
	sensors := make([]Sensor, 0, 128)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		sensors = append(sensors, parse(s))
	})
	aocutil.AOCFinish(part1(sensors), part2(sensors))
}
