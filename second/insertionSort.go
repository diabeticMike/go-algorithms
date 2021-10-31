package second

func InsertionSort(a []int) []int {
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
