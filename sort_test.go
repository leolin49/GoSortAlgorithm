package GoSortAlgorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	RandSeed  = 123456
	ArraySize = 10000
	MaxVal    = 10000
)

var arr = []int{2, 44, 38, 5, 47, 15, 36, 26, 27, 3, 48, 4, 19, 50, 48}

func GetRandSlice() []int {
	// r := rand.New(rand.NewSource(RandSeed))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, ArraySize)
	for i := 0; i < ArraySize; i++ {
		array[i] = r.Intn(MaxVal)
	}
	return array
}

type SortFunc func([]int)

func HandleExample(sf SortFunc) {
	a := make([]int, len(arr))
	copy(a, arr)
	sf(a)
	fmt.Println(a)
}

func ExampleBubbleSort() {
	HandleExample(BubbleSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleSelectSort() {
	HandleExample(SelectSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleInsertSort() {
	HandleExample(InsertSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleShellSort() {
	HandleExample(ShellSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleMergeSort() {
	HandleExample(MergeSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleQuickSort() {
	HandleExample(QuickSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleHeapSort() {
	HandleExample(HeapSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleCountSort() {
	HandleExample(CountSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleBucketSort() {
	HandleExample(BucketSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}
func ExampleRadixSort() {
	HandleExample(RadixSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func HandleBenchmark(b *testing.B, sf SortFunc) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := GetRandSlice()
		sf(a)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	HandleBenchmark(b, BubbleSort)
}

func BenchmarkSelectSort(b *testing.B) {
	HandleBenchmark(b, SelectSort)
}

func BenchmarkInsertSort(b *testing.B) {
	HandleBenchmark(b, InsertSort)
}

func BenchmarkShellSort(b *testing.B) {
	HandleBenchmark(b, ShellSort)
}

func BenchmarkMergeSort(b *testing.B) {
	HandleBenchmark(b, MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
	HandleBenchmark(b, QuickSort)
}

func BenchmarkHeapSort(b *testing.B) {
	HandleBenchmark(b, HeapSort)
}

func BenchmarkCountSort(b *testing.B) {
	HandleBenchmark(b, CountSort)
}

func BenchmarkRadixSort(b *testing.B) {
	HandleBenchmark(b, RadixSort)
}

func BenchmarkBucketSort(b *testing.B) {
	HandleBenchmark(b, BucketSort)
}
