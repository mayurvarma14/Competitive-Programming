package main

import (
	"fmt"
)

type MinHeap struct {
	a []int
}

func NewMinHeap() *MinHeap { return &MinHeap{a: make([]int, 0)} }

func (h *MinHeap) Len() int { return len(h.a) }

func (h *MinHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.a[0], true
}

func (h *MinHeap) Push(x int) {
	h.a = append(h.a, x)
	h.up(h.Len() - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	top := h.a[0]
	h.a[0] = h.a[h.Len()-1]
	h.a = h.a[:h.Len()-1]
	h.down(0)
	return top, true

}

func (h *MinHeap) up(i int) {
	for i > 0 && h.a[parent(i)] > h.a[i] {
		h.a[parent(i)], h.a[i] = h.a[i], h.a[parent(i)]
		i = parent(i)
	}
}

func (h *MinHeap) down(i int) {
	for left(i) < h.Len() {
		child := left(i)
		if right(i) < h.Len() && h.a[right(i)] < h.a[left(i)] {
			child = right(i)
		}
		if h.a[i] <= h.a[child] {
			break
		}
		h.a[i], h.a[child] = h.a[child], h.a[i]
		i = child
	}
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

func main() {
	h := NewMinHeap()
	h.Push(5)
	h.Push(3)
	h.Push(8)
	h.Push(1)
	fmt.Println(h.a)
	fmt.Println(h.Pop()) // expect 1
	fmt.Println(h.Pop()) // expect 3
}
