package main

import (
	"aocutil"
)

type MaxHeap struct {
	length int
	data   []int
}

func CreateMaxHeap(capacity int) MaxHeap {
	heap := MaxHeap{}
	heap.data = make([]int, 0, capacity)
	heap.length = 0
	return heap
}

func (h *MaxHeap) Peek() int {
	return h.data[0]
}

func (h *MaxHeap) Push(value int) {
	h.data = append(h.data, value)
	h.length++
	h.HeapifyUp()
}

func (h *MaxHeap) Pop() int {
	value := h.Peek()
	h.length--
	h.data[0] = h.data[h.length]
	h.HeapifyDown()
	return value
}

func (h *MaxHeap) HeapifyUp() {
	i := h.length - 1
	for {
		parent := (i - 1) / 2
		if parent < 0 || h.data[parent] >= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MaxHeap) HeapifyDown() {
	i := 0
	for {
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if leftChild > h.length {
			break
		}

		biggerChild := leftChild
		if rightChild < h.length && h.data[rightChild] > h.data[leftChild] {
			biggerChild = rightChild
		}

		if h.data[i] >= h.data[biggerChild] {
			break
		}

		h.data[biggerChild], h.data[i] = h.data[i], h.data[biggerChild]
		i = biggerChild
	}
}

func main() {
	calories := CreateMaxHeap(256)
	current := 0
	aocutil.FileReadAllLines("input.txt", func(s string) {
		if len(s) > 0 {
			current += aocutil.Atoi(s)
		} else {
			calories.Push(current)
			current = 0
		}
	})

	part1 := calories.Pop()
	part2 := part1 + calories.Pop() + calories.Pop()

	aocutil.AOCFinish(part1, part2)
}
