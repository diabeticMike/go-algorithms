package second

import (
	"math/rand"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	l := rand.Intn(100)
	arr := make([]int, 0, l)
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(100))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = MergeSort(arr)
	}
}

var a []int
