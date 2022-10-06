package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customStruct struct {
	stack [][]int
	size  int
}

func LRandint(min int, max int) int {
	return min + rand.Intn(max-min)
}
func sise(a []int) int {
	return len(a)
}

func PushFront(size, value int) [][]int {
	c := customStruct{size: size}
	m := make([][]int, c.size)
	for i := 0; i < len(m); i++ {
		t := LRandint(1, len(m))
		for j := 0; j < t; j++ {
			m[i] = append(m[i], value)
		} //create massive
		if len(m) > 1 {
			for h := 0; h < len(m); h++ {
				if len(m[i]) < 1 {
					m = append(m[:h], m[h+1:]...)
					h--
				} //reversed massive
			}
		}
	}
	return m
}

func mergen(a [][]int, b [][]int) [][]int { //double merge sort
	final := make([][]int, 1)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if sise(a[i]) < sise(b[j]) {
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

func mergeSorti(items [][]int) [][]int {
	if len(items) < 2 {
		return items
	}
	first := mergeSorti(items[:len(items)/2])
	second := mergeSorti(items[len(items)/2:])
	return mergen(first, second)
}

func main() {

	var n int
	fmt.Scanln(&n)
	now := time.Now()
	c := customStruct{stack: PushFront(n, 1), size: n}
	fmt.Println(c.stack)
	sorted := mergeSorti(c.stack)
	for i := 0; i < len(sorted); i++ {
		if len(sorted[i]) < 1 {
			sorted = append(sorted[:i], sorted[i+1:]...)
			i--
		}
	}
	durat := time.Since(now)
	fmt.Println(sorted)
	fmt.Println(durat)
	//TODO an OP как считать кол-во операций в разных функциях?
	//TODO посчитать big O

}
