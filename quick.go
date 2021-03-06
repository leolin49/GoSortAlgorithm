package GoSortAlgorithm

func QuickSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	quicksort(arr, 0, n-1)
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := partition(arr, low, high)
	quicksort(arr, low, mid-1)
	quicksort(arr, mid+1, high)
}

func partition(arr []int, low, high int) int {
	key := arr[low]
	l, r := low+1, high
	for l <= r {
		for ; l <= r && arr[l] <= key; l++ {
		}
		for ; l <= r && arr[r] >= key; r-- {
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[r], arr[low] = arr[low], arr[r]
	return r
}
