````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 20 -t get,set -n 1000000
````

````text
====== SET ======
  1000000 requests completed in 14.69 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.375 milliseconds (cumulative count 526729)
75.000% <= 0.439 milliseconds (cumulative count 758078)
87.500% <= 0.503 milliseconds (cumulative count 878454)
93.750% <= 0.591 milliseconds (cumulative count 937979)
96.875% <= 0.743 milliseconds (cumulative count 969329)
98.438% <= 0.967 milliseconds (cumulative count 984499)
99.219% <= 1.335 milliseconds (cumulative count 992235)
99.609% <= 1.863 milliseconds (cumulative count 996097)
99.805% <= 2.703 milliseconds (cumulative count 998050)
99.902% <= 3.815 milliseconds (cumulative count 999028)
99.951% <= 5.535 milliseconds (cumulative count 999513)
99.976% <= 7.567 milliseconds (cumulative count 999756)
99.988% <= 13.583 milliseconds (cumulative count 999878)
99.994% <= 19.999 milliseconds (cumulative count 999939)
99.997% <= 20.591 milliseconds (cumulative count 999970)
99.998% <= 20.783 milliseconds (cumulative count 999986)
99.999% <= 20.863 milliseconds (cumulative count 999993)
100.000% <= 20.911 milliseconds (cumulative count 999997)
100.000% <= 20.943 milliseconds (cumulative count 999999)
100.000% <= 20.975 milliseconds (cumulative count 1000000)
100.000% <= 20.975 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.159% <= 0.207 milliseconds (cumulative count 1594)
19.023% <= 0.303 milliseconds (cumulative count 190232)
65.673% <= 0.407 milliseconds (cumulative count 656731)
87.845% <= 0.503 milliseconds (cumulative count 878454)
94.333% <= 0.607 milliseconds (cumulative count 943334)
96.398% <= 0.703 milliseconds (cumulative count 963978)
97.563% <= 0.807 milliseconds (cumulative count 975633)
98.174% <= 0.903 milliseconds (cumulative count 981743)
98.589% <= 1.007 milliseconds (cumulative count 985892)
98.848% <= 1.103 milliseconds (cumulative count 988478)
99.038% <= 1.207 milliseconds (cumulative count 990381)
99.185% <= 1.303 milliseconds (cumulative count 991851)
99.303% <= 1.407 milliseconds (cumulative count 993025)
99.383% <= 1.503 milliseconds (cumulative count 993834)
99.460% <= 1.607 milliseconds (cumulative count 994603)
99.522% <= 1.703 milliseconds (cumulative count 995224)
99.582% <= 1.807 milliseconds (cumulative count 995820)
99.626% <= 1.903 milliseconds (cumulative count 996260)
99.667% <= 2.007 milliseconds (cumulative count 996670)
99.695% <= 2.103 milliseconds (cumulative count 996949)
99.854% <= 3.103 milliseconds (cumulative count 998538)
99.913% <= 4.103 milliseconds (cumulative count 999134)
99.945% <= 5.103 milliseconds (cumulative count 999448)
99.960% <= 6.103 milliseconds (cumulative count 999605)
99.974% <= 7.103 milliseconds (cumulative count 999739)
99.976% <= 8.103 milliseconds (cumulative count 999757)
99.978% <= 9.103 milliseconds (cumulative count 999778)
99.985% <= 10.103 milliseconds (cumulative count 999846)
99.985% <= 13.103 milliseconds (cumulative count 999848)
99.990% <= 14.103 milliseconds (cumulative count 999900)
99.990% <= 16.103 milliseconds (cumulative count 999905)
99.992% <= 17.103 milliseconds (cumulative count 999925)
99.994% <= 20.111 milliseconds (cumulative count 999943)
100.000% <= 21.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 68059.62 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.411     0.128     0.375     0.631     1.191    20.975
====== GET ======
  1000000 requests completed in 14.27 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.359 milliseconds (cumulative count 521654)
75.000% <= 0.423 milliseconds (cumulative count 772140)
87.500% <= 0.479 milliseconds (cumulative count 883362)
93.750% <= 0.559 milliseconds (cumulative count 940087)
96.875% <= 0.703 milliseconds (cumulative count 969164)
98.438% <= 0.927 milliseconds (cumulative count 984470)
99.219% <= 1.271 milliseconds (cumulative count 992277)
99.609% <= 1.807 milliseconds (cumulative count 996113)
99.805% <= 2.535 milliseconds (cumulative count 998048)
99.902% <= 3.671 milliseconds (cumulative count 999024)
99.951% <= 5.471 milliseconds (cumulative count 999512)
99.976% <= 7.695 milliseconds (cumulative count 999757)
99.988% <= 12.031 milliseconds (cumulative count 999878)
99.994% <= 28.463 milliseconds (cumulative count 999940)
99.997% <= 29.263 milliseconds (cumulative count 999971)
99.998% <= 32.319 milliseconds (cumulative count 999988)
99.999% <= 32.335 milliseconds (cumulative count 999993)
100.000% <= 32.447 milliseconds (cumulative count 999998)
100.000% <= 32.463 milliseconds (cumulative count 1000000)
100.000% <= 32.463 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.233% <= 0.207 milliseconds (cumulative count 2334)
22.891% <= 0.303 milliseconds (cumulative count 228906)
72.170% <= 0.407 milliseconds (cumulative count 721696)
90.843% <= 0.503 milliseconds (cumulative count 908432)
95.367% <= 0.607 milliseconds (cumulative count 953667)
96.916% <= 0.703 milliseconds (cumulative count 969164)
97.841% <= 0.807 milliseconds (cumulative count 978411)
98.347% <= 0.903 milliseconds (cumulative count 983468)
98.719% <= 1.007 milliseconds (cumulative count 987189)
98.959% <= 1.103 milliseconds (cumulative count 989586)
99.142% <= 1.207 milliseconds (cumulative count 991425)
99.267% <= 1.303 milliseconds (cumulative count 992667)
99.372% <= 1.407 milliseconds (cumulative count 993723)
99.449% <= 1.503 milliseconds (cumulative count 994488)
99.518% <= 1.607 milliseconds (cumulative count 995178)
99.567% <= 1.703 milliseconds (cumulative count 995665)
99.611% <= 1.807 milliseconds (cumulative count 996113)
99.647% <= 1.903 milliseconds (cumulative count 996470)
99.687% <= 2.007 milliseconds (cumulative count 996869)
99.715% <= 2.103 milliseconds (cumulative count 997147)
99.864% <= 3.103 milliseconds (cumulative count 998637)
99.923% <= 4.103 milliseconds (cumulative count 999231)
99.945% <= 5.103 milliseconds (cumulative count 999454)
99.958% <= 6.103 milliseconds (cumulative count 999583)
99.970% <= 7.103 milliseconds (cumulative count 999700)
99.976% <= 8.103 milliseconds (cumulative count 999761)
99.981% <= 9.103 milliseconds (cumulative count 999808)
99.983% <= 10.103 milliseconds (cumulative count 999833)
99.984% <= 11.103 milliseconds (cumulative count 999845)
99.988% <= 12.103 milliseconds (cumulative count 999884)
99.991% <= 13.103 milliseconds (cumulative count 999906)
99.992% <= 14.103 milliseconds (cumulative count 999923)
99.992% <= 22.111 milliseconds (cumulative count 999924)
99.993% <= 23.103 milliseconds (cumulative count 999927)
99.996% <= 29.103 milliseconds (cumulative count 999958)
99.997% <= 30.111 milliseconds (cumulative count 999972)
99.998% <= 32.111 milliseconds (cumulative count 999977)
100.000% <= 33.119 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 70072.17 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.394     0.112     0.359     0.599     1.127    32.463
````

## 分析
数据量：100000
单个 value 大小: 20 字节
#### SET
吞吐量：68059.62/s
总耗时：14.69s
p99: 1.191ms
````text
  avg       min       p50       p95       p99       max
0.411     0.128     0.375     0.631     1.191    20.975
````
对于字节大小为 20B 的 value ，吞吐量达到 68k, 相对于字节为 10 的情况下的吞吐量有了 5% 的下降

#### GET
吞吐量：70072.17/s<br>
总耗时：14.27s<br>
p99: 1.127
````text
 avg       min       p50       p95       p99       max
0.394     0.112     0.359     0.599     1.127    32.463
````
