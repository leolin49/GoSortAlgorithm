package sort

func QuickSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	quicksort(arr, 0, n-1)
}

func quicksort(arr []int, low, high int) {
	//	fmt.Println(low, high)
	if low >= high {
		return
	}
	mid := partition(arr, low, high)
	quicksort(arr, low, mid-1)
	quicksort(arr, mid+1, high)
}

func partition(arr []int, low, high int) int {
	//	fmt.Println(arr[low : high+1])
	//fmt.Println(low, high)
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
	//fmt.Println(arr[low], arr[r])
	arr[r], arr[low] = arr[low], arr[r]
	//fmt.Println(arr)
	return r
}

/*
func main() {
	arr := []int{2, 44, 38, 5, 47, 15, 36, 26, 27, 3, 48, 4, 19, 50, 48}
	QuickSort(arr)
	fmt.Println(arr)
}
*/
