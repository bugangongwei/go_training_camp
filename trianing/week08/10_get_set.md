
## 命令
````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 10 -t get,set -n 1000000
````

## 输出
````text
====== SET ======
  1000000 requests completed in 13.80 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 3)
50.000% <= 0.367 milliseconds (cumulative count 537998)
75.000% <= 0.423 milliseconds (cumulative count 764297)
87.500% <= 0.471 milliseconds (cumulative count 878314)
93.750% <= 0.527 milliseconds (cumulative count 941153)
96.875% <= 0.599 milliseconds (cumulative count 969299)
98.438% <= 0.751 milliseconds (cumulative count 984620)
99.219% <= 0.983 milliseconds (cumulative count 992199)
99.609% <= 1.279 milliseconds (cumulative count 996114)
99.805% <= 1.631 milliseconds (cumulative count 998066)
99.902% <= 2.399 milliseconds (cumulative count 999028)
99.951% <= 3.431 milliseconds (cumulative count 999513)
99.976% <= 4.591 milliseconds (cumulative count 999756)
99.988% <= 10.503 milliseconds (cumulative count 999878)
99.994% <= 13.535 milliseconds (cumulative count 999939)
99.997% <= 14.367 milliseconds (cumulative count 999970)
99.998% <= 15.103 milliseconds (cumulative count 999985)
99.999% <= 15.647 milliseconds (cumulative count 999993)
100.000% <= 17.855 milliseconds (cumulative count 999997)
100.000% <= 17.951 milliseconds (cumulative count 999999)
100.000% <= 18.063 milliseconds (cumulative count 1000000)
100.000% <= 18.063 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.226% <= 0.207 milliseconds (cumulative count 2262)
21.415% <= 0.303 milliseconds (cumulative count 214152)
70.977% <= 0.407 milliseconds (cumulative count 709772)
92.112% <= 0.503 milliseconds (cumulative count 921124)
97.084% <= 0.607 milliseconds (cumulative count 970836)
98.158% <= 0.703 milliseconds (cumulative count 981582)
98.721% <= 0.807 milliseconds (cumulative count 987206)
99.029% <= 0.903 milliseconds (cumulative count 990294)
99.265% <= 1.007 milliseconds (cumulative count 992654)
99.416% <= 1.103 milliseconds (cumulative count 994162)
99.541% <= 1.207 milliseconds (cumulative count 995409)
99.634% <= 1.303 milliseconds (cumulative count 996342)
99.707% <= 1.407 milliseconds (cumulative count 997072)
99.757% <= 1.503 milliseconds (cumulative count 997571)
99.799% <= 1.607 milliseconds (cumulative count 997986)
99.823% <= 1.703 milliseconds (cumulative count 998232)
99.844% <= 1.807 milliseconds (cumulative count 998438)
99.856% <= 1.903 milliseconds (cumulative count 998565)
99.868% <= 2.007 milliseconds (cumulative count 998678)
99.878% <= 2.103 milliseconds (cumulative count 998780)
99.941% <= 3.103 milliseconds (cumulative count 999407)
99.969% <= 4.103 milliseconds (cumulative count 999695)
99.979% <= 5.103 milliseconds (cumulative count 999794)
99.984% <= 6.103 milliseconds (cumulative count 999840)
99.986% <= 7.103 milliseconds (cumulative count 999860)
99.987% <= 9.103 milliseconds (cumulative count 999867)
99.988% <= 11.103 milliseconds (cumulative count 999881)
99.990% <= 12.103 milliseconds (cumulative count 999904)
99.992% <= 13.103 milliseconds (cumulative count 999923)
99.996% <= 14.103 milliseconds (cumulative count 999957)
99.999% <= 15.103 milliseconds (cumulative count 999985)
100.000% <= 16.103 milliseconds (cumulative count 999996)
100.000% <= 18.111 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 72474.27 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.384     0.152     0.367     0.543     0.895    18.063
====== GET ======
  1000000 requests completed in 13.85 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.359 milliseconds (cumulative count 526895)
75.000% <= 0.415 milliseconds (cumulative count 756541)
87.500% <= 0.471 milliseconds (cumulative count 880406)
93.750% <= 0.543 milliseconds (cumulative count 941299)
96.875% <= 0.671 milliseconds (cumulative count 969098)
98.438% <= 0.895 milliseconds (cumulative count 984540)
99.219% <= 1.191 milliseconds (cumulative count 992325)
99.609% <= 1.511 milliseconds (cumulative count 996148)
99.805% <= 1.887 milliseconds (cumulative count 998047)
99.902% <= 2.471 milliseconds (cumulative count 999034)
99.951% <= 3.039 milliseconds (cumulative count 999514)
99.976% <= 3.823 milliseconds (cumulative count 999756)
99.988% <= 4.567 milliseconds (cumulative count 999879)
99.994% <= 6.063 milliseconds (cumulative count 999941)
99.997% <= 6.159 milliseconds (cumulative count 999970)
99.998% <= 6.919 milliseconds (cumulative count 999985)
99.999% <= 8.359 milliseconds (cumulative count 999993)
100.000% <= 8.767 milliseconds (cumulative count 999997)
100.000% <= 8.967 milliseconds (cumulative count 999999)
100.000% <= 9.903 milliseconds (cumulative count 1000000)
100.000% <= 9.903 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.247% <= 0.207 milliseconds (cumulative count 2466)
23.705% <= 0.303 milliseconds (cumulative count 237046)
73.022% <= 0.407 milliseconds (cumulative count 730223)
91.610% <= 0.503 milliseconds (cumulative count 916102)
95.997% <= 0.607 milliseconds (cumulative count 959966)
97.240% <= 0.703 milliseconds (cumulative count 972398)
98.017% <= 0.807 milliseconds (cumulative count 980166)
98.484% <= 0.903 milliseconds (cumulative count 984843)
98.838% <= 1.007 milliseconds (cumulative count 988383)
99.065% <= 1.103 milliseconds (cumulative count 990650)
99.261% <= 1.207 milliseconds (cumulative count 992607)
99.401% <= 1.303 milliseconds (cumulative count 994008)
99.519% <= 1.407 milliseconds (cumulative count 995191)
99.608% <= 1.503 milliseconds (cumulative count 996083)
99.681% <= 1.607 milliseconds (cumulative count 996812)
99.732% <= 1.703 milliseconds (cumulative count 997324)
99.777% <= 1.807 milliseconds (cumulative count 997768)
99.810% <= 1.903 milliseconds (cumulative count 998099)
99.836% <= 2.007 milliseconds (cumulative count 998359)
99.853% <= 2.103 milliseconds (cumulative count 998529)
99.954% <= 3.103 milliseconds (cumulative count 999536)
99.980% <= 4.103 milliseconds (cumulative count 999804)
99.992% <= 5.103 milliseconds (cumulative count 999917)
99.997% <= 6.103 milliseconds (cumulative count 999966)
99.999% <= 7.103 milliseconds (cumulative count 999986)
99.999% <= 8.103 milliseconds (cumulative count 999990)
100.000% <= 9.103 milliseconds (cumulative count 999999)
100.000% <= 10.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 72196.95 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.384     0.136     0.359     0.567     1.079     9.903
````
## 分析
数据量：100000
单个 value 大小: 10 字节
#### SET
吞吐量：72474.27/s
总耗时：13.80s
p99: 0.895ms
````text
  avg       min       p50       p95       p99       max
0.384     0.152     0.367     0.543     0.895    18.063
````
对于字节大小为 10B 的 value ，吞吐量达到 72k, 请求延迟大部分都能达到 0.895ms

#### GET
吞吐量：72196.957/s<br>
总耗时：13.85s<br>
p99: 1.079s
````text
  avg       min       p50       p95       p99       max
0.384     0.136     0.359     0.567     1.079     9.903
````
GET 和 SET 命令在吞吐量上没有差别；在延迟上也差别不大；