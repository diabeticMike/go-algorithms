package main

import (
	"fmt"

	"github.com/algorithms/second"
)

func main() {
	a := []int{1, 0, 1, 0, 0}
	b := []int{1, 1, 0, 0, 0}
	fmt.Println(second.BinSum2(a, b))
}
