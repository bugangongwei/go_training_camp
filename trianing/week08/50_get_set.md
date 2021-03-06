````text
redis-benchmark -h 127.0.0.1 -p 6379 -d 50 -t get,set -n 1000000
````

````text
====== SET ======
  1000000 requests completed in 17.13 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.095 milliseconds (cumulative count 1)
50.000% <= 0.383 milliseconds (cumulative count 516574)
75.000% <= 0.471 milliseconds (cumulative count 756832)
87.500% <= 0.583 milliseconds (cumulative count 876443)
93.750% <= 0.767 milliseconds (cumulative count 939130)
96.875% <= 1.031 milliseconds (cumulative count 969061)
98.438% <= 1.431 milliseconds (cumulative count 984542)
99.219% <= 2.111 milliseconds (cumulative count 992192)
99.609% <= 3.599 milliseconds (cumulative count 996104)
99.805% <= 8.487 milliseconds (cumulative count 998047)
99.902% <= 12.207 milliseconds (cumulative count 999024)
99.951% <= 23.391 milliseconds (cumulative count 999512)
99.976% <= 34.751 milliseconds (cumulative count 999756)
99.988% <= 59.295 milliseconds (cumulative count 999878)
99.994% <= 114.623 milliseconds (cumulative count 999939)
99.997% <= 116.095 milliseconds (cumulative count 999970)
99.998% <= 124.735 milliseconds (cumulative count 999985)
99.999% <= 127.423 milliseconds (cumulative count 999994)
100.000% <= 127.615 milliseconds (cumulative count 999997)
100.000% <= 127.679 milliseconds (cumulative count 1000000)
100.000% <= 127.679 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 1)
0.096% <= 0.207 milliseconds (cumulative count 960)
15.493% <= 0.303 milliseconds (cumulative count 154933)
59.838% <= 0.407 milliseconds (cumulative count 598382)
80.606% <= 0.503 milliseconds (cumulative count 806058)
88.908% <= 0.607 milliseconds (cumulative count 889079)
92.400% <= 0.703 milliseconds (cumulative count 923997)
94.625% <= 0.807 milliseconds (cumulative count 946251)
95.856% <= 0.903 milliseconds (cumulative count 958564)
96.739% <= 1.007 milliseconds (cumulative count 967390)
97.321% <= 1.103 milliseconds (cumulative count 973212)
97.809% <= 1.207 milliseconds (cumulative count 978095)
98.139% <= 1.303 milliseconds (cumulative count 981388)
98.406% <= 1.407 milliseconds (cumulative count 984060)
98.594% <= 1.503 milliseconds (cumulative count 985941)
98.755% <= 1.607 milliseconds (cumulative count 987545)
98.879% <= 1.703 milliseconds (cumulative count 988793)
98.990% <= 1.807 milliseconds (cumulative count 989900)
99.076% <= 1.903 milliseconds (cumulative count 990760)
99.147% <= 2.007 milliseconds (cumulative count 991471)
99.215% <= 2.103 milliseconds (cumulative count 992145)
99.541% <= 3.103 milliseconds (cumulative count 995407)
99.660% <= 4.103 milliseconds (cumulative count 996604)
99.706% <= 5.103 milliseconds (cumulative count 997055)
99.754% <= 6.103 milliseconds (cumulative count 997540)
99.773% <= 7.103 milliseconds (cumulative count 997735)
99.792% <= 8.103 milliseconds (cumulative count 997920)
99.817% <= 9.103 milliseconds (cumulative count 998173)
99.833% <= 10.103 milliseconds (cumulative count 998331)
99.869% <= 11.103 milliseconds (cumulative count 998694)
99.901% <= 12.103 milliseconds (cumulative count 999009)
99.925% <= 13.103 milliseconds (cumulative count 999254)
99.931% <= 14.103 milliseconds (cumulative count 999315)
99.937% <= 16.103 milliseconds (cumulative count 999367)
99.942% <= 17.103 milliseconds (cumulative count 999423)
99.942% <= 18.111 milliseconds (cumulative count 999424)
99.943% <= 19.103 milliseconds (cumulative count 999431)
99.947% <= 20.111 milliseconds (cumulative count 999474)
99.951% <= 23.103 milliseconds (cumulative count 999506)
99.955% <= 24.111 milliseconds (cumulative count 999552)
99.957% <= 26.111 milliseconds (cumulative count 999569)
99.961% <= 27.103 milliseconds (cumulative count 999610)
99.965% <= 28.111 milliseconds (cumulative count 999647)
99.965% <= 29.103 milliseconds (cumulative count 999651)
99.967% <= 30.111 milliseconds (cumulative count 999670)
99.974% <= 31.103 milliseconds (cumulative count 999743)
99.975% <= 32.111 milliseconds (cumulative count 999750)
99.976% <= 35.103 milliseconds (cumulative count 999764)
99.978% <= 36.127 milliseconds (cumulative count 999783)
99.979% <= 37.119 milliseconds (cumulative count 999792)
99.980% <= 38.111 milliseconds (cumulative count 999800)
99.980% <= 46.111 milliseconds (cumulative count 999801)
99.981% <= 52.127 milliseconds (cumulative count 999813)
99.987% <= 54.111 milliseconds (cumulative count 999871)
99.987% <= 55.103 milliseconds (cumulative count 999872)
99.990% <= 60.127 milliseconds (cumulative count 999899)
99.990% <= 65.119 milliseconds (cumulative count 999903)
99.993% <= 72.127 milliseconds (cumulative count 999934)
99.995% <= 115.135 milliseconds (cumulative count 999950)
99.997% <= 116.159 milliseconds (cumulative count 999971)
99.998% <= 117.119 milliseconds (cumulative count 999984)
99.999% <= 125.119 milliseconds (cumulative count 999985)
100.000% <= 128.127 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 58390.75 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.491     0.088     0.383     0.839     1.823   127.679
====== GET ======
  1000000 requests completed in 17.03 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.111 milliseconds (cumulative count 3)
50.000% <= 0.359 milliseconds (cumulative count 511013)
75.000% <= 0.455 milliseconds (cumulative count 763077)
87.500% <= 0.575 milliseconds (cumulative count 876179)
93.750% <= 0.767 milliseconds (cumulative count 937826)
96.875% <= 1.039 milliseconds (cumulative count 968835)
98.438% <= 1.455 milliseconds (cumulative count 984457)
99.219% <= 2.247 milliseconds (cumulative count 992203)
99.609% <= 5.199 milliseconds (cumulative count 996098)
99.805% <= 10.703 milliseconds (cumulative count 998047)
99.902% <= 13.143 milliseconds (cumulative count 999025)
99.951% <= 23.167 milliseconds (cumulative count 999512)
99.976% <= 40.415 milliseconds (cumulative count 999756)
99.988% <= 72.127 milliseconds (cumulative count 999885)
99.994% <= 109.247 milliseconds (cumulative count 999939)
99.997% <= 114.111 milliseconds (cumulative count 999971)
99.998% <= 124.031 milliseconds (cumulative count 1000000)
100.000% <= 124.031 milliseconds (cumulative count 1000000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.212% <= 0.207 milliseconds (cumulative count 2124)
23.928% <= 0.303 milliseconds (cumulative count 239284)
66.191% <= 0.407 milliseconds (cumulative count 661907)
82.522% <= 0.503 milliseconds (cumulative count 825217)
89.117% <= 0.607 milliseconds (cumulative count 891171)
92.321% <= 0.703 milliseconds (cumulative count 923211)
94.492% <= 0.807 milliseconds (cumulative count 944920)
95.781% <= 0.903 milliseconds (cumulative count 957806)
96.678% <= 1.007 milliseconds (cumulative count 966777)
97.242% <= 1.103 milliseconds (cumulative count 972421)
97.717% <= 1.207 milliseconds (cumulative count 977174)
98.055% <= 1.303 milliseconds (cumulative count 980553)
98.333% <= 1.407 milliseconds (cumulative count 983331)
98.545% <= 1.503 milliseconds (cumulative count 985451)
98.728% <= 1.607 milliseconds (cumulative count 987281)
98.859% <= 1.703 milliseconds (cumulative count 988585)
98.965% <= 1.807 milliseconds (cumulative count 989654)
99.046% <= 1.903 milliseconds (cumulative count 990462)
99.117% <= 2.007 milliseconds (cumulative count 991168)
99.167% <= 2.103 milliseconds (cumulative count 991672)
99.416% <= 3.103 milliseconds (cumulative count 994160)
99.531% <= 4.103 milliseconds (cumulative count 995312)
99.603% <= 5.103 milliseconds (cumulative count 996028)
99.663% <= 6.103 milliseconds (cumulative count 996629)
99.713% <= 7.103 milliseconds (cumulative count 997125)
99.740% <= 8.103 milliseconds (cumulative count 997403)
99.760% <= 9.103 milliseconds (cumulative count 997597)
99.778% <= 10.103 milliseconds (cumulative count 997784)
99.827% <= 11.103 milliseconds (cumulative count 998270)
99.858% <= 12.103 milliseconds (cumulative count 998585)
99.902% <= 13.103 milliseconds (cumulative count 999022)
99.911% <= 14.103 milliseconds (cumulative count 999106)
99.926% <= 15.103 milliseconds (cumulative count 999263)
99.929% <= 16.103 milliseconds (cumulative count 999290)
99.932% <= 17.103 milliseconds (cumulative count 999320)
99.935% <= 18.111 milliseconds (cumulative count 999348)
99.940% <= 20.111 milliseconds (cumulative count 999398)
99.947% <= 21.103 milliseconds (cumulative count 999474)
99.949% <= 22.111 milliseconds (cumulative count 999491)
99.951% <= 23.103 milliseconds (cumulative count 999508)
99.952% <= 24.111 milliseconds (cumulative count 999519)
99.953% <= 25.103 milliseconds (cumulative count 999528)
99.957% <= 26.111 milliseconds (cumulative count 999569)
99.959% <= 28.111 milliseconds (cumulative count 999594)
99.961% <= 31.103 milliseconds (cumulative count 999606)
99.963% <= 32.111 milliseconds (cumulative count 999634)
99.965% <= 33.119 milliseconds (cumulative count 999648)
99.967% <= 35.103 milliseconds (cumulative count 999670)
99.972% <= 36.127 milliseconds (cumulative count 999722)
99.974% <= 38.111 milliseconds (cumulative count 999737)
99.974% <= 39.103 milliseconds (cumulative count 999740)
99.974% <= 40.127 milliseconds (cumulative count 999744)
99.976% <= 41.119 milliseconds (cumulative count 999761)
99.977% <= 44.127 milliseconds (cumulative count 999773)
99.977% <= 45.119 milliseconds (cumulative count 999774)
99.978% <= 52.127 milliseconds (cumulative count 999777)
99.980% <= 53.119 milliseconds (cumulative count 999804)
99.983% <= 56.127 milliseconds (cumulative count 999829)
99.984% <= 57.119 milliseconds (cumulative count 999835)
99.985% <= 61.119 milliseconds (cumulative count 999849)
99.985% <= 64.127 milliseconds (cumulative count 999850)
99.986% <= 66.111 milliseconds (cumulative count 999864)
99.989% <= 72.127 milliseconds (cumulative count 999885)
99.990% <= 73.151 milliseconds (cumulative count 999905)
99.992% <= 109.119 milliseconds (cumulative count 999919)
99.994% <= 110.143 milliseconds (cumulative count 999939)
99.997% <= 114.111 milliseconds (cumulative count 999971)
99.997% <= 115.135 milliseconds (cumulative count 999973)
99.997% <= 123.135 milliseconds (cumulative count 999974)
100.000% <= 124.159 milliseconds (cumulative count 1000000)

Summary:
  throughput summary: 58706.12 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.489     0.104     0.359     0.847     1.847   124.031
````
