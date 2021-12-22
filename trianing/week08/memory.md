### before
````text
~ redis-cli info memory
# Memory
used_memory:1087616
used_memory_human:1.04M
used_memory_rss:52064256
used_memory_rss_human:49.65M
used_memory_peak:501088128
used_memory_peak_human:477.87M
used_memory_peak_perc:0.22%
used_memory_overhead:1027608
used_memory_startup:1009120
used_memory_dataset:60008
used_memory_dataset_perc:76.45%
allocator_allocated:1032720
allocator_active:52017152
allocator_resident:52017152
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:47104
used_memory_lua_human:46.00K
used_memory_scripts:1048
used_memory_scripts_human:1.02K
number_of_cached_scripts:3
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:50.37
allocator_frag_bytes:50984432
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:47104
mem_fragmentation_ratio:50.41
mem_fragmentation_bytes:51031536
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
````

## 写入数据

````text
#!/bin/bash
v=""
t="a"

for ((i=0;i<100;i++))
do
v="${v}${t}"
done

echo $v

for ((i=0;i<100000;i++))
do
echo -en $v | redis-cli -x set name$i >>redis.log
done
echo "done"
````
```text
~ bash data.sh
done
```

## after
```text
~ redis-cli info memory
# Memory
used_memory:19736192
used_memory_human:18.82M
used_memory_rss:52953088
used_memory_rss_human:50.50M
used_memory_peak:501088128
used_memory_peak_human:477.87M
used_memory_peak_perc:3.94%
used_memory_overhead:6076184
used_memory_startup:1009120
used_memory_dataset:13660008
used_memory_dataset_perc:72.94%
allocator_allocated:19681296
allocator_active:52905984
allocator_resident:52905984
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:47104
used_memory_lua_human:46.00K
used_memory_scripts:1048
used_memory_scripts_human:1.02K
number_of_cached_scripts:3
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.69
allocator_frag_bytes:33224688
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:47104
mem_fragmentation_ratio:2.69
mem_fragmentation_bytes:33271792
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

19736192 - 1087616 = 2.22MB