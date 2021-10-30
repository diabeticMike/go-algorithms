package main

import (
	"math"
)

func main() {
}

func bitSum(f, s []int) []int {
	result := make([]int, len(f)+1)
	for i := len(f); i > 0; i-- {
		if f[i-1] == 1 && s[i-1] == 1 {
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

func convertTo2(in int) []int {
	out := make([]int, 0, 0)
	for {
		out = append(out, in%2)
		in /= 2
		if in == 0 {
			return reverse(out)
		}
	}
}

func reverse(in []int) []int {
	out := make([]int, 0, len(in))
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return out
}

func convertTo10(a []int) int {
	var res int
	for k, v := range a {
		if v == 1 {
			res += int(math.Pow(2.0, float64(len(a)-k-1)))
		}
	}
	return res
}
