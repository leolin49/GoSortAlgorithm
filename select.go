package GoSortAlgorithm

func SelectSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		idx, val := i, arr[i]
		for j := i + 1; j < n; j++ {
			if arr[j] < val {
				val = arr[j]
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func Sort(arr []int) {
	SelectSort(arr)
}
