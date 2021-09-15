package GoSortAlgorithm

func InsertSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
}
