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

func bubbleSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	return a
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var t int = len(a) / 2
	// fmt.Printf("%v|%v\n", a[:t], a[t:])
	l := mergeSort(a[:t])
	r := mergeSort(a[t:])
	b := make([]int, len(a))
	bi := 0
	li := 0
	ri := 0
	for {
		if len(l)-li == 0 && len(r)-ri > 0 {
			copy(b[bi:], r[ri:])
			break
		} else if len(l)-li > 0 && len(r)-ri == 0 {
			copy(b[bi:], l[li:])
			break
		}
		if l[li] < r[ri] {
			b[bi] = l[li]
			bi++
			li++
		} else {
			b[bi] = r[ri]
			bi++
			ri++
		}
	}
	// fmt.Printf("%v|%v->%v\n", l, r, b)
	copy(a, b)
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
	runASort("bubble sort", bubbleSort)
	runASort("merge sort", mergeSort)
}
