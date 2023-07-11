## Task 1 Initial version 

1. - col : The column on which to execute the opertaion , default is 1 
2. - op : The operation to be performed , supports sum and averge

1. Create the errors.go ,csv.go and main.go files
2. Both sum and average have the same function signature , so we create the type func 

```go 
type statsFunc func([]float64) float64
```
3. Create the csv2flow function which takes a csv file and returns a channel of float64 slices


## Writing Test

1. Create a file called `csv_test.go` and add the test 
2. For the TestCSV2Float add the test for the error cases too !errors.Is(err, tc.expErr) can be used to compare the error 


```bash
touch testdata/example.csv
IP Address,Timestamp,Response Time,Byt192.168.0.199,1520698621,236,34192.168.0.88,1520698776,220,32192.168.0.199,1520699033,226,32192.168.0.100,1520699142,218,34192.168.0.199,1520699379,238,3822

```

```bash
IP Address,Timestamp,Response Time,Bytes
192.168.0.199,1520698621,236,3475
192.168.0.88,1520698776,220,3200
192.168.0.199,1520699033,226,3200
192.168.0.100,1520699142,218,3475
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
192.168.0.199,1520699379,238,3822
```


## Writing a benchmark Test

- Go provides the `testing.B` type to help write benchmarks.
- In Go, when you run benchmarks, the testing package automatically determines the number of iterations needed to obtain a stable measurement of the benchmarked code's performance. The b.N value represents this number, which is dynamically determined based on the benchmark's execution time
```bash
 go test -bench . -benchtime=10x -run ^$
```
^$ - denotes only run the benchmark tests
- benchtime flag is used to specify the number of times the benchmark should be run

## Profiling your tool 

- Go profiler shows you a breakdown of where your program spends its execution time and how much memory it uses.

```bash
go test -bench . -benchtime=10x -run ^$ -cpuprofile cpu.prof
```

- Now analyze the profile using the `go tool pprof` command

```bash
go tool pprof cpu.prof
```
- Inside the pprof tool, you can use the top command to see the top functions that consume CPU time.
- You can further investigate the function by list function name eg csv2float64
- You can use the web command to generate a report in SVG format and open it in your browser.

`- Memory profiling`

```bash
go test -bench . -benchtime=10x -run ^$ -memprofile mem.prof
```

```bash
go tool -alloc_space pprof mem.prof
```
You can use the `-benchmem` to display the memory allocation statistics for the benchmarked code.

```bash
go test -bench . -benchtime=10x -run ^$ -benchmem
```

From the above discussion we can see that the csv2float64 function is the bottleneck in our program. We can optimize this function to improve the performance of our program.

## Reducing the memory allocation 

- Instead of ReadAll function we are going to read the file line by line using the ReadLine function


## Tracing 

- Sometimes a program is spending time waiting for resources to be available, such as a network connection or a lock. In such cases, the CPU profiler won't be able to help you find the bottleneck. You can use the tracing tool to find out what your program is doing when it's not using the CPU.

```bash
go test -bench . -benchtime=10x -run ^$ -trace trace.out
```

```bash
go tool trace trace.out
```
