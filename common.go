package GoSortAlgorithm

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

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
