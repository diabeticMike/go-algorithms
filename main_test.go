package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_BitSum(t *testing.T) {
	n := 64
	for i := 0; i < n; i++ {
		a := convertTo2(i)
		b := convertTo2(n - i)
		if len(a) < len(b) {
			tmp := a
			a = make([]int, len(b)-len(a))
			a = append(a, tmp...)
		} else if len(b) < len(a) {
			tmp := b
			b = make([]int, len(a)-len(b))
			b = append(b, tmp...)
		}
		if i+n-i != convertTo10(bitSum(a, b)) {
			log.Fatal(i+n-i, "  ", convertTo10(bitSum(a, b)))
		}
	}
	fmt.Println("SUCCESS")
}
func Benchmark_BitSum(b *testing.B) {
	n := 64
	for j := 0; j < b.N; j++ {
		for i := 0; i < n; i++ {
			a := convertTo2(i)
			b := convertTo2(n - i)
			if len(a) < len(b) {
				tmp := a
				a = make([]int, len(b)-len(a))
				a = append(a, tmp...)
			} else if len(b) < len(a) {
				tmp := b
				b = make([]int, len(a)-len(b))
				b = append(b, tmp...)
			}
			if i+n-i != convertTo10(bitSum(a, b)) {
				log.Fatal(i+n-i, "  ", convertTo10(bitSum(a, b)))
			}
		}
	}
}
func Benchmark_BitSum2(b *testing.B) {
	n := 64
	for j := 0; j < b.N; j++ {
		for i := 0; i < n; i++ {
			a := convertTo2(i)
			b := convertTo2(n - i)
			if len(a) < len(b) {
				tmp := a
				a = make([]int, len(b)-len(a))
				a = append(a, tmp...)
			} else if len(b) < len(a) {
				tmp := b
				b = make([]int, len(a)-len(b))
				b = append(b, tmp...)
			}
			if i+n-i != convertTo10(bitSum2(a, b)) {
				log.Fatal(i+n-i, "  ", convertTo10(bitSum2(a, b)))
			}
		}
	}
}
