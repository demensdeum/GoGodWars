package main

import (
	//"container/list"
	"fmt"
)

type Point struct{ X, Y int }

type HeightMap struct {
	Width  int
	Height int
	Data   [][]int
}

func NewHeightMap(w, h int) *HeightMap {
	data := make([][]int, h)
	for i := range data {
		data[i] = make([]int, w)
	}
	return &HeightMap{Width: w, Height: h, Data: data}
}

func (hm *HeightMap) SmoothMap() {
	w, h := hm.Width, hm.Height

	fmt.Println(w,h)
}

func (hm *HeightMap) Print() {
	for _, row := range hm.Data {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	hm := NewHeightMap(6, 6)

	hm.Data = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 1, 3, 1, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	fmt.Println("Before smoothing:")
	hm.Print()

	hm.SmoothMap()

	fmt.Println("\nAfter smoothing:")
	hm.Print()
}
