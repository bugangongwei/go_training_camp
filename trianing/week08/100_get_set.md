
## 命令
````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 100 -t get,set -n 1000000
````

## 输出
````text
====== SET ======
  1000000 requests completed in 13.68 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 1)
50.000% <= 0.367 milliseconds (cumulative count 503521)
75.000% <= 0.431 milliseconds (cumulative count 774268)
87.500% <= 0.471 milliseconds (cumulative count 875046)
93.750% <= 0.519 milliseconds (cumulative count 939912)
96.875% <= 0.575 milliseconds (cumulative count 971284)
98.438% <= 0.647 milliseconds (cumulative count 984837)
99.219% <= 0.783 milliseconds (cumulative count 992216)
99.609% <= 1.007 milliseconds (cumulative count 996164)
99.805% <= 1.231 milliseconds (cumulative count 998063)
99.902% <= 1.479 milliseconds (cumulative count 999030)
99.951% <= 1.831 milliseconds (cumulative count 999512)
99.976% <= 2.327 milliseconds (cumulative count 999757)
99.988% <= 3.239 milliseconds (cumulative count 999878)
99.994% <= 5.999 milliseconds (cumulative count 999939)
99.997% <= 6.679 milliseconds (cumulative count 999970)
99.998% <= 7.567 milliseconds (cumulative count 999985)
99.999% <= 9.807 milliseconds (cumulative count 999993)
100.000% <= 11.495 milliseconds (cumulative count 999997)
100.000% <= 11.551 milliseconds (cumulative count 999999)
100.000% <= 11.559 milliseconds (cumulative count 1000000)
100.000% <= 11.559 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.154% <= 0.207 milliseconds (cumulative count 1537)
19.581% <= 0.303 milliseconds (cumulative count 195813)
68.730% <= 0.407 milliseconds (cumulative count 687302)
92.354% <= 0.503 milliseconds (cumulative count 923539)
97.933% <= 0.607 milliseconds (cumulative count 979326)
98.904% <= 0.703 milliseconds (cumulative count 989038)
99.283% <= 0.807 milliseconds (cumulative count 992828)
99.462% <= 0.903 milliseconds (cumulative count 994618)
99.616% <= 1.007 milliseconds (cumulative count 996164)
99.717% <= 1.103 milliseconds (cumulative count 997168)
99.792% <= 1.207 milliseconds (cumulative count 997922)
99.842% <= 1.303 milliseconds (cumulative count 998416)
99.883% <= 1.407 milliseconds (cumulative count 998831)
99.908% <= 1.503 milliseconds (cumulative count 999081)
99.926% <= 1.607 milliseconds (cumulative count 999259)
99.938% <= 1.703 milliseconds (cumulative count 999384)
99.949% <= 1.807 milliseconds (cumulative count 999492)
99.959% <= 1.903 milliseconds (cumulative count 999585)
99.967% <= 2.007 milliseconds (cumulative count 999666)
99.971% <= 2.103 milliseconds (cumulative count 999714)
99.987% <= 3.103 milliseconds (cumulative count 999874)
99.989% <= 4.103 milliseconds (cumulative count 999895)
99.990% <= 5.103 milliseconds (cumulative count 999898)
99.994% <= 6.103 milliseconds (cumulative count 999943)
99.998% <= 7.103 milliseconds (cumulative count 999983)
99.999% <= 8.103 milliseconds (cumulative count 999989)
99.999% <= 9.103 milliseconds (cumulative count 999991)
99.999% <= 10.103 milliseconds (cumulative count 999993)
100.000% <= 12.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 73083.39 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.381     0.152     0.367     0.535     0.727    11.559
====== GET ======
  1000000 requests completed in 14.27 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.127 milliseconds (cumulative count 2)
50.000% <= 0.375 milliseconds (cumulative count 530595)
75.000% <= 0.431 milliseconds (cumulative count 770149)
87.500% <= 0.479 milliseconds (cumulative count 882794)
93.750% <= 0.535 milliseconds (cumulative count 942063)
96.875% <= 0.615 milliseconds (cumulative count 969581)
98.438% <= 0.791 milliseconds (cumulative count 984423)
99.219% <= 1.047 milliseconds (cumulative count 992320)
99.609% <= 1.327 milliseconds (cumulative count 996097)
99.805% <= 1.647 milliseconds (cumulative count 998076)
99.902% <= 1.991 milliseconds (cumulative count 999025)
99.951% <= 2.359 milliseconds (cumulative count 999516)
99.976% <= 2.815 milliseconds (cumulative count 999756)
99.988% <= 4.199 milliseconds (cumulative count 999878)
99.994% <= 5.383 milliseconds (cumulative count 999939)
99.997% <= 7.727 milliseconds (cumulative count 999970)
99.998% <= 8.055 milliseconds (cumulative count 999985)
99.999% <= 8.191 milliseconds (cumulative count 999993)
100.000% <= 8.271 milliseconds (cumulative count 999997)
100.000% <= 8.303 milliseconds (cumulative count 999999)
100.000% <= 8.335 milliseconds (cumulative count 1000000)
100.000% <= 8.335 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.097% <= 0.207 milliseconds (cumulative count 970)
17.167% <= 0.303 milliseconds (cumulative count 171668)
68.142% <= 0.407 milliseconds (cumulative count 681422)
91.505% <= 0.503 milliseconds (cumulative count 915052)
96.803% <= 0.607 milliseconds (cumulative count 968029)
97.955% <= 0.703 milliseconds (cumulative count 979553)
98.515% <= 0.807 milliseconds (cumulative count 985146)
98.867% <= 0.903 milliseconds (cumulative count 988672)
99.145% <= 1.007 milliseconds (cumulative count 991446)
99.333% <= 1.103 milliseconds (cumulative count 993328)
99.483% <= 1.207 milliseconds (cumulative count 994834)
99.586% <= 1.303 milliseconds (cumulative count 995859)
99.672% <= 1.407 milliseconds (cumulative count 996722)
99.733% <= 1.503 milliseconds (cumulative count 997329)
99.789% <= 1.607 milliseconds (cumulative count 997892)
99.830% <= 1.703 milliseconds (cumulative count 998296)
99.861% <= 1.807 milliseconds (cumulative count 998609)
99.884% <= 1.903 milliseconds (cumulative count 998839)
99.905% <= 2.007 milliseconds (cumulative count 999053)
99.920% <= 2.103 milliseconds (cumulative count 999201)
99.982% <= 3.103 milliseconds (cumulative count 999817)
99.987% <= 4.103 milliseconds (cumulative count 999874)
99.993% <= 5.103 milliseconds (cumulative count 999931)
99.995% <= 6.103 milliseconds (cumulative count 999950)
99.999% <= 8.103 milliseconds (cumulative count 999987)
100.000% <= 9.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 70086.91 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.391     0.120     0.375     0.551     0.951     8.335
````
## 分析
数据量：100000
单个 value 大小: 100 字节
#### SET
吞吐量：73083.39.27/s
总耗时：13.68s
p99: 0.727ms
````text
  avg       min       p50       p95       p99       max
0.381     0.152     0.367     0.535     0.727    11.559
````
对于字节大小为 100B 的 value ，相比于字节大小为 10B 的吞吐量肯定是有增加的<br>
但是相对来说耗时和 p99 反而表现更好，说明 10 倍增的吞吐量不太影响延迟

#### GET
吞吐量：70086.91/s<br>
总耗时：14.27s<br>
p99: 0.951
````text
  avg       min       p50       p95       p99       max
0.391     0.120     0.375     0.551     0.951     8.335
````

#### 总结
机器性能是会有一些波动的，表现在 redis 性能上是会有 100B 比 10B 表现好的情况，暂时来看，10B 和 100B 的性能差距不大；
