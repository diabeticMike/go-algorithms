package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertionSort(arr))
}

func insertionSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key { // a[i] < key in case of descending
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return a
}
