package GoSortAlgorithm

func CountSort(arr []int) {
	max, min := GetMaxAndMin(arr)
	cnt := make([]int, max-min+1)
	for _, val := range arr {
		cnt[val-min]++
	}
	idx := 0
	for k, v := range cnt {
		for i := 0; i < v; i++ {
			arr[idx] = k + min
			idx++
		}
	}
}
