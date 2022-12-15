package main

import (
	"aocutil"
	"math"
	"regexp"
)

type Vector struct {
	x, y int
}

type Rect struct {
	left, top, right, bottom int
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

func find_bounds(sensors []Sensor) Rect {
	bounds := Rect{}
	bounds.left = math.MaxInt
	bounds.top = math.MaxInt
	bounds.right = math.MinInt
	bounds.bottom = math.MinInt
	for _, s := range sensors {
		sensorLeft := s.position.x - s.radius
		sensorRight := s.position.x + s.radius
		sensorTop := s.position.y - s.radius
		sensorBottom := s.position.y + s.radius
		if sensorLeft < bounds.left {
			bounds.left = sensorLeft
		}
		if sensorRight > bounds.right {
			bounds.right = sensorRight
		}
		if sensorTop < bounds.top {
			bounds.top = sensorTop
		}
		if sensorBottom > bounds.bottom {
			bounds.bottom = sensorBottom
		}
	}
	return bounds
}

func main() {
	sensors := make([]Sensor, 0, 128)
	aocutil.FileReadAllLines("input.txt", func(s string) {
		sensors = append(sensors, parse(s))
	})

	beacons := make(map[Vector]bool)
	for _, s := range sensors {
		beacons[s.closestBeacon] = true
	}

	y := 2000000
	bounds := find_bounds(sensors)

	part1 := 0
	for x := bounds.left; x < bounds.right; x++ {
		pos := Vector{x, y}
	done:
		for _, s := range sensors {
			dx := abs(pos.x - s.position.x)
			dy := abs(pos.y - s.position.y)
			d := dx + dy
			if d <= s.radius {
				if _, found := beacons[pos]; !found {
					part1++
					break done
				}
			}
		}
	}

	aocutil.AOCFinish(part1)
}
