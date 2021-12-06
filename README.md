# go语言实现十个经典排序算法

## benchmark测试结果

```powershell
PS E:\GoSortAlgorithm> go test -bench .
goos: windows
goarch: amd64
pkg: GoSortAlgorithm
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkBubbleSort-6                 12          96679492 ns/op               0 B/op          0 allocs/op
BenchmarkSelectSort-6                 25          47674376 ns/op               0 B/op          0 allocs/op
BenchmarkInsertSort-6                 79          14008396 ns/op               0 B/op          0 allocs/op
BenchmarkShellSort-6                1701            849523 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-6                1126           1131423 ns/op         1112708 B/op       9999 allocs/op
BenchmarkQuickSort-6                1872            609207 ns/op               0 B/op          0 allocs/op
BenchmarkHeapSort-6                 1016           1200000 ns/op               0 B/op          0 allocs/op
BenchmarkCountSort-6                3242            321729 ns/op          802820 B/op          1 allocs/op
BenchmarkRadixSort-6                 918           1285081 ns/op          925213 B/op        560 allocs/op
BenchmarkBucketSort-6                100          10746767 ns/op          160240 B/op      10011 allocs/op
PASS
ok      GoSortAlgorithm 14.485s

```



## benchmark 生成 profile

参考：https://geektutu.com/post/hpg-pprof.html

1. 执行benchmark测试并生成profile

```powershell
go test -bench="Sort$" -cpuprofile="cpu.pprof" .
```

2. 在浏览器中分析cpu/内存性能数据

```powershell
go tool pprof -http=:11111 cpu.pprof
```

访问http://localhost:11111/ui/

> ps:如果提示 Graphviz 没有安装，则通过 `brew install graphviz`(MAC) 或 `apt install graphviz`(Ubuntu) 即可

