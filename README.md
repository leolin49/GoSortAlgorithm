# go语言实现十个经典排序算法

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

![image-20211206172918652](C:\Users\linyufeng\AppData\Roaming\Typora\typora-user-images\image-20211206172918652.png)

