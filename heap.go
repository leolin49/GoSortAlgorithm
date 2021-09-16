package GoSortAlgorithm

func HeapSort(arr []int) {
	n := len(arr)
	build(arr)
	for n > 0 {
		arr[0], arr[n-1] = arr[n-1], arr[0]
		n--
		adjust(arr, 0, n)
	}
}

func build(arr []int) {
	for idx := len(arr)>>1 - 1; idx >= 0; idx-- {
		adjust(arr, idx, len(arr))
	}
}

func adjust(arr []int, idx int, n int) {
	maxIdx := idx
	if l := idx<<1 + 1; l < n && arr[l] > arr[maxIdx] {
		maxIdx = l
	}
	if r := idx<<1 + 2; r < n && arr[r] > arr[maxIdx] {
		maxIdx = r
	}
	if maxIdx != idx {
		arr[maxIdx], arr[idx] = arr[idx], arr[maxIdx]
		adjust(arr, maxIdx, n)
	}
}
