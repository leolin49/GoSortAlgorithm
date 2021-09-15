package sort

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

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

func GetMaxAndMin(arr []int) (max, min int) {
	max, min = INT_MIN, INT_MAX
	for _, val := range arr {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max, min
}
