package second

func MergeSort(A []int) []int {
	if len(A) == 1 {
		return A
	}
	middle := len(A) / 2
	L := append(A[:0:0], A[:middle]...)
	R := append(A[:0:0], A[middle:]...)
	L = MergeSort(L)
	R = MergeSort(R)
	i, j := 0, 0
	for k := 0; i < len(L) && j < len(R); k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
	if j == len(R) {
		A = append(A[:i+j], L[i:]...)
	} else {
		A = append(A[:i+j], R[j:]...)
	}
	return A
}
