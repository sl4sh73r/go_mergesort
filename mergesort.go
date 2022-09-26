package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(items [][]int) [][]int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func size(a []int) int {
	return len(a)
}

func Randint(min int, max int) int {

	return min + rand.Intn(max-min)
}

func merge(a [][]int, b [][]int) [][]int {
	final := make([][]int, 1)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if size(a[i]) < size(b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func createMssive(size int) [][]int {
	m := make([][]int, size)
	reversed := make([][]int, size)
	for i := 0; i < size; i++ {
		t := Randint(1, size)
		for j := 0; j < t; j++ {
			m[i] = append(m[i], 1)
		}
	}
	for i := range m {
		n := m[len(m)-1-i]
		//fmt.Println(n) -- sanity check
		reversed = append(reversed, n)
	}
	return reversed
}

func main() {
	var n int
	fmt.Scanln(&n)
	now := time.Now()
	unsorted := createMssive(n)
	fmt.Println("unsorted")
	for i := 0; i < len(unsorted); i++ {
		if len(unsorted[i]) < 1 {
			unsorted = append(unsorted[:i], unsorted[i+1:]...)
			i--
		}
	}
	fmt.Println(unsorted)
	sorted := mergeSort(unsorted)
	for i := 0; i < len(sorted); i++ {
		if len(sorted[i]) < 1 {
			sorted = append(sorted[:i], sorted[i+1:]...)
			i--
		}
	}
	fmt.Println("sorted")
	fmt.Println(sorted)
	durat := time.Since(now)
	fmt.Println(durat)
}
