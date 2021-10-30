package main

import (
	"fmt"
	"math"
)

const n = 5

func main() {
	arr20 := [n]int{1, 0, 1, 0, 0} // 20
	arr24 := [n]int{1, 0, 1, 1, 0} // 22
	fmt.Println(convertTo10(bitSum(arr20, arr24)))
}

func bitSum(f, s [n]int) [n + 1]int {
	var result [n + 1]int

	for i := n; i > 0; i-- {
		if f[i-1] == 1 && s[i-i] == 1 {
			if result[i] == 1 {
				result[i-1] = 1
				continue
			} else {
				result[i] = 0
				result[i-1] = 1
			}
		} else if f[i-1] == 1 || s[i-1] == 1 {
			if result[i] == 1 {
				result[i-1] = 1
				result[i] = 0
				continue
			}
			result[i] = 1
		}
	}
	return result
}

func convertTo10(a [n + 1]int) int {
	var res int
	for k, v := range a {
		if v == 1 {
			res += int(math.Pow(2.0, float64(len(a)-k-1)))
		}
	}
	return res
}
