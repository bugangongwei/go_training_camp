
## 命令
````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 5000 -t get,set -n 1000000
````

## 输出
````text
====== SET ======
  1000000 requests completed in 14.48 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.391 milliseconds (cumulative count 522736)
75.000% <= 0.455 milliseconds (cumulative count 761857)
87.500% <= 0.511 milliseconds (cumulative count 884314)
93.750% <= 0.567 milliseconds (cumulative count 942073)
96.875% <= 0.639 milliseconds (cumulative count 970060)
98.438% <= 0.767 milliseconds (cumulative count 984402)
99.219% <= 0.975 milliseconds (cumulative count 992353)
99.609% <= 1.239 milliseconds (cumulative count 996153)
99.805% <= 1.583 milliseconds (cumulative count 998051)
99.902% <= 2.151 milliseconds (cumulative count 999026)
99.951% <= 3.111 milliseconds (cumulative count 999512)
99.976% <= 4.199 milliseconds (cumulative count 999760)
99.988% <= 5.103 milliseconds (cumulative count 999878)
99.994% <= 7.367 milliseconds (cumulative count 999939)
99.997% <= 11.607 milliseconds (cumulative count 999970)
99.998% <= 11.927 milliseconds (cumulative count 999985)
99.999% <= 12.191 milliseconds (cumulative count 999993)
100.000% <= 12.679 milliseconds (cumulative count 999997)
100.000% <= 12.879 milliseconds (cumulative count 999999)
100.000% <= 16.863 milliseconds (cumulative count 1000000)
100.000% <= 16.863 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.077% <= 0.207 milliseconds (cumulative count 775)
14.684% <= 0.303 milliseconds (cumulative count 146837)
59.023% <= 0.407 milliseconds (cumulative count 590235)
87.169% <= 0.503 milliseconds (cumulative count 871691)
96.142% <= 0.607 milliseconds (cumulative count 961416)
97.943% <= 0.703 milliseconds (cumulative count 979431)
98.658% <= 0.807 milliseconds (cumulative count 986578)
99.047% <= 0.903 milliseconds (cumulative count 990474)
99.302% <= 1.007 milliseconds (cumulative count 993017)
99.460% <= 1.103 milliseconds (cumulative count 994596)
99.584% <= 1.207 milliseconds (cumulative count 995838)
99.666% <= 1.303 milliseconds (cumulative count 996656)
99.731% <= 1.407 milliseconds (cumulative count 997310)
99.777% <= 1.503 milliseconds (cumulative count 997771)
99.812% <= 1.607 milliseconds (cumulative count 998122)
99.836% <= 1.703 milliseconds (cumulative count 998364)
99.856% <= 1.807 milliseconds (cumulative count 998557)
99.870% <= 1.903 milliseconds (cumulative count 998703)
99.885% <= 2.007 milliseconds (cumulative count 998854)
99.898% <= 2.103 milliseconds (cumulative count 998980)
99.951% <= 3.103 milliseconds (cumulative count 999510)
99.973% <= 4.103 milliseconds (cumulative count 999733)
99.988% <= 5.103 milliseconds (cumulative count 999878)
99.990% <= 6.103 milliseconds (cumulative count 999901)
99.993% <= 7.103 milliseconds (cumulative count 999927)
99.994% <= 8.103 milliseconds (cumulative count 999942)
99.995% <= 9.103 milliseconds (cumulative count 999950)
99.995% <= 10.103 milliseconds (cumulative count 999953)
99.999% <= 12.103 milliseconds (cumulative count 999989)
100.000% <= 13.103 milliseconds (cumulative count 999999)
100.000% <= 17.103 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 69079.86 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.407     0.112     0.391     0.583     0.895    16.863
====== GET ======
  1000000 requests completed in 18.16 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.407 milliseconds (cumulative count 509439)
75.000% <= 0.495 milliseconds (cumulative count 756314)
87.500% <= 0.623 milliseconds (cumulative count 876670)
93.750% <= 0.863 milliseconds (cumulative count 938086)
96.875% <= 1.191 milliseconds (cumulative count 969070)
98.438% <= 1.679 milliseconds (cumulative count 984477)
99.219% <= 2.471 milliseconds (cumulative count 992217)
99.609% <= 3.999 milliseconds (cumulative count 996098)
99.805% <= 5.831 milliseconds (cumulative count 998047)
99.902% <= 9.143 milliseconds (cumulative count 999025)
99.951% <= 17.343 milliseconds (cumulative count 999512)
99.976% <= 27.903 milliseconds (cumulative count 999756)
99.988% <= 61.247 milliseconds (cumulative count 999879)
99.994% <= 78.015 milliseconds (cumulative count 999939)
99.997% <= 107.519 milliseconds (cumulative count 999970)
99.998% <= 107.775 milliseconds (cumulative count 999994)
100.000% <= 107.903 milliseconds (cumulative count 999997)
100.000% <= 108.287 milliseconds (cumulative count 1000000)
100.000% <= 108.287 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.030% <= 0.207 milliseconds (cumulative count 301)
7.231% <= 0.303 milliseconds (cumulative count 72313)
50.944% <= 0.407 milliseconds (cumulative count 509439)
76.993% <= 0.503 milliseconds (cumulative count 769934)
86.896% <= 0.607 milliseconds (cumulative count 868964)
90.532% <= 0.703 milliseconds (cumulative count 905321)
92.888% <= 0.807 milliseconds (cumulative count 928885)
94.368% <= 0.903 milliseconds (cumulative count 943683)
95.558% <= 1.007 milliseconds (cumulative count 955581)
96.352% <= 1.103 milliseconds (cumulative count 963521)
96.994% <= 1.207 milliseconds (cumulative count 969942)
97.459% <= 1.303 milliseconds (cumulative count 974590)
97.822% <= 1.407 milliseconds (cumulative count 978224)
98.085% <= 1.503 milliseconds (cumulative count 980853)
98.308% <= 1.607 milliseconds (cumulative count 983078)
98.489% <= 1.703 milliseconds (cumulative count 984886)
98.659% <= 1.807 milliseconds (cumulative count 986589)
98.790% <= 1.903 milliseconds (cumulative count 987902)
98.903% <= 2.007 milliseconds (cumulative count 989035)
98.991% <= 2.103 milliseconds (cumulative count 989907)
99.418% <= 3.103 milliseconds (cumulative count 994185)
99.622% <= 4.103 milliseconds (cumulative count 996219)
99.726% <= 5.103 milliseconds (cumulative count 997257)
99.821% <= 6.103 milliseconds (cumulative count 998213)
99.852% <= 7.103 milliseconds (cumulative count 998524)
99.883% <= 8.103 milliseconds (cumulative count 998826)
99.901% <= 9.103 milliseconds (cumulative count 999015)
99.913% <= 10.103 milliseconds (cumulative count 999125)
99.928% <= 11.103 milliseconds (cumulative count 999285)
99.938% <= 12.103 milliseconds (cumulative count 999376)
99.943% <= 13.103 milliseconds (cumulative count 999431)
99.946% <= 14.103 milliseconds (cumulative count 999459)
99.949% <= 15.103 milliseconds (cumulative count 999490)
99.950% <= 16.103 milliseconds (cumulative count 999503)
99.951% <= 17.103 milliseconds (cumulative count 999507)
99.952% <= 18.111 milliseconds (cumulative count 999523)
99.960% <= 19.103 milliseconds (cumulative count 999605)
99.966% <= 20.111 milliseconds (cumulative count 999657)
99.966% <= 21.103 milliseconds (cumulative count 999661)
99.967% <= 22.111 milliseconds (cumulative count 999672)
99.971% <= 23.103 milliseconds (cumulative count 999705)
99.971% <= 24.111 milliseconds (cumulative count 999713)
99.972% <= 25.103 milliseconds (cumulative count 999717)
99.973% <= 26.111 milliseconds (cumulative count 999733)
99.974% <= 27.103 milliseconds (cumulative count 999743)
99.976% <= 28.111 milliseconds (cumulative count 999764)
99.977% <= 30.111 milliseconds (cumulative count 999766)
99.980% <= 31.103 milliseconds (cumulative count 999799)
99.980% <= 34.111 milliseconds (cumulative count 999804)
99.981% <= 37.119 milliseconds (cumulative count 999807)
99.983% <= 38.111 milliseconds (cumulative count 999828)
99.986% <= 50.111 milliseconds (cumulative count 999858)
99.986% <= 59.103 milliseconds (cumulative count 999861)
99.986% <= 60.127 milliseconds (cumulative count 999862)
99.987% <= 61.119 milliseconds (cumulative count 999867)
99.990% <= 62.111 milliseconds (cumulative count 999900)
99.990% <= 63.103 milliseconds (cumulative count 999901)
99.991% <= 69.119 milliseconds (cumulative count 999911)
99.991% <= 77.119 milliseconds (cumulative count 999912)
99.994% <= 78.143 milliseconds (cumulative count 999939)
99.997% <= 105.151 milliseconds (cumulative count 999966)
100.000% <= 108.159 milliseconds (cumulative count 999997)
100.000% <= 109.119 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 55069.11 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.519     0.136     0.407     0.959     2.119   108.287
````
## 分析
数据量：100000
单个 value 大小: 5000 字节
#### SET
吞吐量：69079.86/s
总耗时：14.48s
p99: 0.895ms
````text
  avg       min       p50       p95       p99       max
0.407     0.112     0.391     0.583     0.895    16.863
````

#### GET
吞吐量：55069.11/s<br>
总耗时：18.16s<br>
p99: 2.119s
````text
  avg       min       p50       p95       p99       max
0.519     0.136     0.407     0.959     2.119   108.287
````

10B -> 5000B 的变化，延迟增加一倍， 吞吐量降低了一丢丢； 总体来说，差别不大；