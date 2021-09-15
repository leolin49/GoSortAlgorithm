package sort

func MergeSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	MergeSort(arr[:n>>1])
	MergeSort(arr[n>>1:])
	merge(arr)
}

func merge(arr []int) {
	n := len(arr)
	help := make([]int, n)
	copy(help, arr)
	mid := n >> 1
	l, r, idx := 0, mid, 0
	for l < mid && r < n {
		if help[l] < help[r] {
			arr[idx] = help[l]
			l++
		} else {
			arr[idx] = help[r]
			r++
		}
		idx++
	}
	for ; l < mid; l, idx = l+1, idx+1 {
		arr[idx] = help[l]
	}
}
