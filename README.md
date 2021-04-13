The results are the following:
```
go test -run=^$ -benchmem -v -bench=. -benchtime=10s
goos: linux
goarch: amd64
pkg: bench2
cpu: Intel Core Processor (Skylake, IBRS)
BenchmarkNew
BenchmarkNew-16       	   55899	    217737 ns/op	    1683 B/op	      36 allocs/op
BenchmarkMaster
BenchmarkMaster-16    	   51490	    238629 ns/op	    1803 B/op	      38 allocs/op
PASS
ok  	bench2	29.110s
```
``` 
go test -run=^$ -benchmem -v -bench=. -benchtime=30s
goos: linux
goarch: amd64
pkg: bench2
cpu: Intel Core Processor (Skylake, IBRS)
BenchmarkNew
BenchmarkNew-16       	  169681	    217293 ns/op	    1683 B/op	      36 allocs/op
BenchmarkMaster
BenchmarkMaster-16    	  155041	    237994 ns/op	    1803 B/op	      38 allocs/op
PASS
ok  	bench2	79.376s
```