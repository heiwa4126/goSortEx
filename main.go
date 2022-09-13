package main

// sort algorithm examples
// from https://cs50.jp/x/2021/week3/notes/

import (
	"fmt"
)

func genList() []int {
	return []int{6, 3, 8, 5, 2, 7, 4, 1}
}

func selectionSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	for i := range a {
		minIdx := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			a[minIdx], a[i] = a[i], a[minIdx]
		}
	}
	return a
}

func runASort(name string, f func([]int) []int) {
	fmt.Printf("%s\n", name)
	a := genList()
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", f(a))
}

func main() {
	runASort("selection sort", selectionSort)
}
