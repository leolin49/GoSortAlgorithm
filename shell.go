package sortalgorithm

func ShellSort(arr []int) {
	n := len(arr)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			val := arr[i]
			idx := i - gap
			for idx >= 0 && val < arr[idx] {
				arr[idx+gap] = arr[idx]
				idx -= gap
			}
			arr[idx+gap] = val
		}
		gap /= 2
	}
}
