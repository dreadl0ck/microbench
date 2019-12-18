# Throughput benchmarks


## Qemu
```bash
Benchmarking 10.10.10.2 (be patient)
Completed 5000 requests
Completed 10000 requests
Completed 15000 requests
Completed 20000 requests
Finished 20053 requests


Server Software:        
Server Hostname:        10.10.10.2
Server Port:            80

Document Path:          /ping
Document Length:        81920 bytes

Concurrency Level:      5
Time taken for tests:   30.003 seconds
Complete requests:      20053
Failed requests:        0
Keep-Alive requests:    0
Total transferred:      1644884285 bytes
HTML transferred:       1642938756 bytes
Requests per second:    668.36 [#/sec] (mean)
Time per request:       7.481 [ms] (mean)
Time per request:       1.496 [ms] (mean, across all concurrent requests)
Transfer rate:          53538.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     1    7   2.4      7      27
Waiting:        1    4   1.8      4      20
Total:          2    7   2.4      7      27

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%      9
  80%      9
  90%     11
  95%     12
  98%     13
  99%     14
 100%     27 (longest request)
```

## Firecracker

```bash
Benchmarking 10.10.10.2 (be patient)                                           
Completed 5000 requests                                                        
Completed 10000 requests                                                       
Completed 15000 requests                                                       
Completed 20000 requests                                                       
Finished 20072 requests                                                        
                                                                               
                                                                               
Server Software:                                                               
Server Hostname:        10.10.10.2                                             
Server Port:            80                                                     
                                                                               
Document Path:          /ping                                                  
Document Length:        81920 bytes                                            
                                                                               
Concurrency Level:      5                                                      
Time taken for tests:   30.000 seconds                                         
Complete requests:      20072                                                  
Failed requests:        0                                                      
Keep-Alive requests:    0                                                      
Total transferred:      1646288416 bytes                                       
HTML transferred:       1644341335 bytes                                       
Requests per second:    669.07 [#/sec] (mean)                                  
Time per request:       7.473 [ms] (mean)                                      
Time per request:       1.495 [ms] (mean, across all concurrent requests)      
Transfer rate:          53590.02 [Kbytes/sec] received                         
                                                                               
Connection Times (ms)                                                          
              min  mean[+/-sd] median   max                                                                                                                    
Connect:        0    0   0.1      0       1                                                                                                                    
Processing:     2    7   2.5      7      22                                                                                                                    
Waiting:        1    4   1.8      4      16                                                                                                                    
Total:          2    7   2.5      7      23                                                                                                                    
                                                                                                                                                               
Percentage of the requests served within a certain time (ms)                                                                                                   
  50%      7                                                                                                                                                   
  66%      8                                                                                                                                                   
  75%      9                                                                                                                                                   
  80%      9                                                                                                                                                   
  90%     11                                                                                                                                                   
  95%     12                                                                                                                                                   
  98%     13                   
  99%     14                                                                                                                                                   
 100%     23 (longest request)
```
