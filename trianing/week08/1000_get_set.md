
## 命令
````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 1000 -t get,set -n 1000000
````

## 输出
````text
====== SET ======
  1000000 requests completed in 13.37 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.127 milliseconds (cumulative count 2)
50.000% <= 0.359 milliseconds (cumulative count 519203)
75.000% <= 0.415 milliseconds (cumulative count 751114)
87.500% <= 0.463 milliseconds (cumulative count 878272)
93.750% <= 0.511 milliseconds (cumulative count 943966)
96.875% <= 0.551 milliseconds (cumulative count 969193)
98.438% <= 0.623 milliseconds (cumulative count 985132)
99.219% <= 0.783 milliseconds (cumulative count 992300)
99.609% <= 1.103 milliseconds (cumulative count 996133)
99.805% <= 1.511 milliseconds (cumulative count 998061)
99.902% <= 2.295 milliseconds (cumulative count 999029)
99.951% <= 3.511 milliseconds (cumulative count 999514)
99.976% <= 5.215 milliseconds (cumulative count 999756)
99.988% <= 7.447 milliseconds (cumulative count 999878)
99.994% <= 15.351 milliseconds (cumulative count 999939)
99.997% <= 85.311 milliseconds (cumulative count 999973)
99.998% <= 85.503 milliseconds (cumulative count 999988)
99.999% <= 85.567 milliseconds (cumulative count 999994)
100.000% <= 85.631 milliseconds (cumulative count 999999)
100.000% <= 85.695 milliseconds (cumulative count 1000000)
100.000% <= 85.695 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.256% <= 0.207 milliseconds (cumulative count 2562)
24.054% <= 0.303 milliseconds (cumulative count 240544)
72.274% <= 0.407 milliseconds (cumulative count 722737)
93.633% <= 0.503 milliseconds (cumulative count 936332)
98.325% <= 0.607 milliseconds (cumulative count 983251)
99.018% <= 0.703 milliseconds (cumulative count 990178)
99.274% <= 0.807 milliseconds (cumulative count 992741)
99.421% <= 0.903 milliseconds (cumulative count 994205)
99.541% <= 1.007 milliseconds (cumulative count 995412)
99.613% <= 1.103 milliseconds (cumulative count 996133)
99.677% <= 1.207 milliseconds (cumulative count 996773)
99.727% <= 1.303 milliseconds (cumulative count 997269)
99.772% <= 1.407 milliseconds (cumulative count 997725)
99.804% <= 1.503 milliseconds (cumulative count 998040)
99.831% <= 1.607 milliseconds (cumulative count 998306)
99.844% <= 1.703 milliseconds (cumulative count 998436)
99.858% <= 1.807 milliseconds (cumulative count 998585)
99.870% <= 1.903 milliseconds (cumulative count 998704)
99.881% <= 2.007 milliseconds (cumulative count 998812)
99.888% <= 2.103 milliseconds (cumulative count 998882)
99.942% <= 3.103 milliseconds (cumulative count 999415)
99.958% <= 4.103 milliseconds (cumulative count 999581)
99.973% <= 5.103 milliseconds (cumulative count 999730)
99.983% <= 6.103 milliseconds (cumulative count 999833)
99.984% <= 7.103 milliseconds (cumulative count 999845)
99.988% <= 8.103 milliseconds (cumulative count 999883)
99.989% <= 9.103 milliseconds (cumulative count 999887)
99.989% <= 11.103 milliseconds (cumulative count 999892)
99.990% <= 12.103 milliseconds (cumulative count 999896)
99.990% <= 13.103 milliseconds (cumulative count 999904)
99.991% <= 14.103 milliseconds (cumulative count 999907)
99.993% <= 15.103 milliseconds (cumulative count 999931)
99.995% <= 16.103 milliseconds (cumulative count 999950)
99.996% <= 85.119 milliseconds (cumulative count 999961)
100.000% <= 86.143 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 74771.95 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.377     0.120     0.359     0.519     0.703    85.695
====== GET ======
  1000000 requests completed in 13.18 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.111 milliseconds (cumulative count 1)
50.000% <= 0.351 milliseconds (cumulative count 509071)
75.000% <= 0.407 milliseconds (cumulative count 759085)
87.500% <= 0.455 milliseconds (cumulative count 890614)
93.750% <= 0.487 milliseconds (cumulative count 939411)
96.875% <= 0.527 milliseconds (cumulative count 971135)
98.438% <= 0.567 milliseconds (cumulative count 984693)
99.219% <= 0.639 milliseconds (cumulative count 992461)
99.609% <= 0.783 milliseconds (cumulative count 996114)
99.805% <= 0.983 milliseconds (cumulative count 998056)
99.902% <= 1.215 milliseconds (cumulative count 999037)
99.951% <= 1.391 milliseconds (cumulative count 999523)
99.976% <= 1.559 milliseconds (cumulative count 999756)
99.988% <= 1.791 milliseconds (cumulative count 999882)
99.994% <= 2.111 milliseconds (cumulative count 999940)
99.997% <= 2.431 milliseconds (cumulative count 999973)
99.998% <= 2.519 milliseconds (cumulative count 999985)
99.999% <= 2.567 milliseconds (cumulative count 999993)
100.000% <= 2.767 milliseconds (cumulative count 999997)
100.000% <= 2.871 milliseconds (cumulative count 999999)
100.000% <= 2.895 milliseconds (cumulative count 1000000)
100.000% <= 2.895 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.111% <= 0.207 milliseconds (cumulative count 1108)
23.999% <= 0.303 milliseconds (cumulative count 239992)
75.909% <= 0.407 milliseconds (cumulative count 759085)
95.514% <= 0.503 milliseconds (cumulative count 955137)
99.034% <= 0.607 milliseconds (cumulative count 990338)
99.473% <= 0.703 milliseconds (cumulative count 994729)
99.643% <= 0.807 milliseconds (cumulative count 996434)
99.737% <= 0.903 milliseconds (cumulative count 997371)
99.820% <= 1.007 milliseconds (cumulative count 998203)
99.867% <= 1.103 milliseconds (cumulative count 998670)
99.902% <= 1.207 milliseconds (cumulative count 999018)
99.932% <= 1.303 milliseconds (cumulative count 999319)
99.955% <= 1.407 milliseconds (cumulative count 999552)
99.967% <= 1.503 milliseconds (cumulative count 999673)
99.980% <= 1.607 milliseconds (cumulative count 999800)
99.985% <= 1.703 milliseconds (cumulative count 999851)
99.988% <= 1.807 milliseconds (cumulative count 999884)
99.990% <= 1.903 milliseconds (cumulative count 999901)
99.992% <= 2.007 milliseconds (cumulative count 999919)
99.994% <= 2.103 milliseconds (cumulative count 999938)
100.000% <= 3.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 75866.78 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.363     0.104     0.351     0.503     0.607     2.895
````
## 分析
数据量：100000
单个 value 大小: 1000 字节
#### SET
吞吐量：74771.95/s
总耗时：13.37s
p99: 0.703ms
````text
   avg       min       p50       p95       p99       max
0.377     0.120     0.359     0.519     0.703    85.695
````

#### GET
吞吐量：75866.78/s<br>
总耗时：13.18s<br>
p99: 0.607ms
````text
 avg       min       p50       p95       p99       max
0.363     0.104     0.351     0.503     0.607     2.895
````
看起来，redis 的延迟受吞吐量的影响不大