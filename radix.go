package GoSortAlgorithm

import "math"

func RadixSort(arr []int) {
	maxDigit := 0
	for max, _ := GetMaxAndMin(arr); max != 0; max /= 10 {
		maxDigit++
	}
	for i := 1; i <= maxDigit; i++ {
		var cnt [10][]int
		for _, val := range arr {
			idx := val % (int(math.Pow10(i))) / int(math.Pow10(i-1))
			cnt[idx] = append(cnt[idx], val)
		}
		//
		idx := 0
		for j := 0; j < 10; j++ {
			for _, val := range cnt[j] {
				arr[idx] = val
				idx++
			}
		}
	}
}
