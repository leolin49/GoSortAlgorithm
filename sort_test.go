package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	RandSeed  = 123456
	ArraySize = 10000
)

var arr = []int{2, 44, 38, 5, 47, 15, 36, 26, 27, 3, 48, 4, 19, 50, 48}

func GetRandSlice() []int {
	r := rand.New(rand.NewSource(RandSeed))
	array := make([]int, ArraySize)
	for i := 0; i < ArraySize; i++ {
		array[i] = r.Intn(10000)
	}
	return array
}

type SortFunc func([]int)

func HandleExample(sf SortFunc) {
	a := arr
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

func ExampleQuickSort() {
	HandleExample(QuickSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func ExampleShellSort() {
	HandleExample(ShellSort)
	// Output: [2 3 4 5 15 19 26 27 36 38 44 47 48 48 50]
}

func HandleBenchmark(b *testing.B, sf SortFunc) {
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

func BenchmarkQuickSort(b *testing.B) {
	HandleBenchmark(b, QuickSort)
}
