# Cache Service Infura

## Intro

This package has been written as part of a showcase of my knowledge of golang.


My idea is to used this package inside a docker container and substitute the direct calls
to the third party with a large instance of radis that contact the third party api
or that have access to an eth node directly.

As well to be able to be geographically more close to the consumers in order to be able to aggregate requests
that otherwise would be directed to the infura api or to a eth full node.

I have as well decided to take the freedom to add a limiter for the usage of the api
for each source, I will show the api in action with and without the usage of the limiter.

I have used [goland](https://www.jetbrains.com/go/) as IDE.

I have used [exalys](https://github.com/exercism/exalysis) to run the suite of automated go tools.

## Requirements

 - docker

## Dependencies

   - [deep](github.com/go-test/deep) it has been useful in testing to compare the returned nested structure with the one expected
   - [gorilla/mux](github.com/gorilla/mux) it provides an easy to configure routing system compatible whit net/http
   - [http-cache](github.com/victorspringer/http-cache)  provides a wrapper that can be used around the routing handler to cache the responses, useful because it allows different eviction policies, for the purpose of this package i have chosen Lsu.
   
   thanks to go module there is no need to go get -t each package
   
   ``` go mod download```
    
it will download all the modules if not already available.

## Usage

clone this repo and enter the repo and run the following commands,
depending on your system privileges on the host machine it may require 
to insert your credentials if you have access to sudo privileges or
be restricted.  

The second command will assign 4 cpu to the docker container and 50MB of ram

```
   sudo docker build -t appinfura:1.0 .
   sudo docker run -d -p 8001:8123 -it --cpus="4" --memory=500m appinfura:1.0
```
 

## Test

### Unit Test Coverage

Looking at the results of the test coverage its a 100% in every folder contained in the
package, the file main and the function main is excluded from the unit test.

I would like to underline that this load testing is done on the same machine environment
and do not takes into account the latency due to connectivity in a real case scenario.

```
go test -cover  ./...
```

```
?       github.com/LucaPaterlini/infura [no test files]
ok      github.com/LucaPaterlini/infura/API     0.835s  coverage: 100.0% of statements
?       github.com/LucaPaterlini/infura/config  [no test files]
ok      github.com/LucaPaterlini/infura/dataCollection  1.048s  coverage: 100.0% of statements
ok      github.com/LucaPaterlini/infura/middlewares/limit       3.004s  coverage: 100.0% of statements
ok      github.com/LucaPaterlini/infura/middlewares/logger      0.004s  coverage: 100.0% of statements

```

### Load Test without limiter

This test has been conducted on a i7 8 core computer 3.5gh each core assigning 1 core in the first test
and 4 cores in the second test with 500Mb of ram allocated to the docker container
using [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html)

The ram choice its arbitrary but more than what http-cache and the scratch os requires.

I will executed the test again on 4 cpu later to show the effect of the function limiter,
with the change to apply to docker to be able to run it.

``` bash loadTest/loadtestblock.sh```

### cpu 1

First Request to cache, this first test is useful to understand how much time does it takes
in case the call is not already cached, as you can see looking at the first test its clear that its completely depended
from the latency on the third party api.

#### not cached 1 cpu block endpoint 
```
Concurrency Level:      1
Time taken for tests:   0.123 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      1592 bytes
HTML transferred:       1473 bytes
Requests per second:    8.11 [#/sec] (mean)
Time per request:       123.252 [ms] (mean)
Time per request:       123.252 [ms] (mean, across all concurrent requests)
Transfer rate:          12.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   123  123   0.0    123     123
Waiting:      123  123   0.0    123     123
Total:        123  123   0.0    123     123
```
#### cached 1 cpu block endpoint

```
Concurrency Level:      64
Time taken for tests:   31.034 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      159200000 bytes
HTML transferred:       147300000 bytes
Requests per second:    3222.30 [#/sec] (mean)
Time per request:       19.862 [ms] (mean)
Time per request:       0.310 [ms] (mean, across all concurrent requests)
Transfer rate:          5009.66 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       4
Processing:     0   19  32.4      6     511
Waiting:        0   19  32.3      6     511
Total:          0   20  32.3      6     511

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      7
  75%      8
  80%     12
  90%     70
  95%     89
  98%    100
  99%    108
 100%    511 (longest request)
```

#### not cached  1 cpu tx endpoint

```
Concurrency Level:      1
Time taken for tests:   0.352 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      159 bytes
HTML transferred:       42 bytes
Requests per second:    2.84 [#/sec] (mean)
Time per request:       351.576 [ms] (mean)
Time per request:       351.576 [ms] (mean, across all concurrent requests)
Transfer rate:          0.44 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   351  351   0.0    351     351
Waiting:      351  351   0.0    351     351
Total:        352  352   0.0    352     352
```
#### cached 1 cpu tx endpoint

```
Concurrency Level:      64
Time taken for tests:   32.678 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      15900000 bytes
HTML transferred:       4200000 bytes
Requests per second:    3060.12 [#/sec] (mean)
Time per request:       20.914 [ms] (mean)
Time per request:       0.327 [ms] (mean, across all concurrent requests)
Transfer rate:          475.16 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       4
Processing:     0   20  31.1      7     292
Waiting:        0   20  31.0      6     291
Total:          0   21  31.0      7     292

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%     10
  80%     59
  90%     71
  95%     91
  98%     97
  99%    101
 100%    292 (longest request)

```

#### not cached 4 cpu block endpoint 

```
Concurrency Level:      1
Time taken for tests:   0.127 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      1592 bytes
HTML transferred:       1473 bytes
Requests per second:    7.85 [#/sec] (mean)
Time per request:       127.389 [ms] (mean)
Time per request:       127.389 [ms] (mean, across all concurrent requests)
Transfer rate:          12.20 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   127  127   0.0    127     127
Waiting:      127  127   0.0    127     127
Total:        127  127   0.0    127     127
```

#### cached 4 cpu block endpoint

```
Concurrency Level:      64
Time taken for tests:   11.643 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      159200000 bytes
HTML transferred:       147300000 bytes
Requests per second:    8589.00 [#/sec] (mean)
Time per request:       7.451 [ms] (mean)
Time per request:       0.116 [ms] (mean, across all concurrent requests)
Transfer rate:          13353.21 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.6      0      15
Processing:     0    7   6.9      5     240
Waiting:        0    6   6.8      5     239
Total:          0    7   6.8      6     240
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      7
  75%      7
  80%      8
  90%     11
  95%     21
  98%     32
  99%     37
 100%    240 (longest request)
```

#### not cached  4 cpu tx endpoint

```
Concurrency Level:      1
Time taken for tests:   0.120 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      159 bytes
HTML transferred:       42 bytes
Requests per second:    8.30 [#/sec] (mean)
Time per request:       120.470 [ms] (mean)
Time per request:       120.470 [ms] (mean, across all concurrent requests)
Transfer rate:          1.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   120  120   0.0    120     120
Waiting:      120  120   0.0    120     120
Total:        120  120   0.0    120     120

```
#### cached  4 cpu tx endpoint

```
Concurrency Level:      64
Time taken for tests:   14.149 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      15900000 bytes
HTML transferred:       4200000 bytes
Requests per second:    7067.64 [#/sec] (mean)
Time per request:       9.055 [ms] (mean)
Time per request:       0.141 [ms] (mean, across all concurrent requests)
Transfer rate:          1097.42 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.8      1      18
Processing:     0    8   6.7      7     242
Waiting:        0    8   6.7      6     240
Total:          1    9   6.6      8     243

Percentage of the requests served within a certain time (ms)
  50%      8
  66%      9
  75%     10
  80%     10
  90%     13
  95%     19
  98%     31
  99%     39
 100%    243 (longest request)

```

## Results of the load test

As shown in those load tests where we have excluded the potential issue related to lack of ram
adding a new cpu at 3.5gh speed will allow to server 1500 more requests each second,
as expected as well there adding more cpu reduce the impact due to the required time for sync 
of the routines.

As well as expected the call of a not cached endpoint is fully dependent on the infura api response,
time.

## Showing the Limiter

In order to be able to activate the logger replace line 30 and re-initialize the docker with the following line.
```
CMD ["./main","-limiter=true"]
```

In this load test with the limiter activated its visible an improvement in performance
even though checking if a user have already accessed its still expensive ~40% time
of providing the actual cached copy of the endpoint but still  ~ 0.0025% of
providing a not cached endpoint from the infura api.

```
Concurrency Level:      64
Time taken for tests:   12.798 seconds
Complete requests:      100000
Failed requests:        99858
   (Connect: 0, Receive: 0, Length: 99858, Exceptions: 0)
Non-2xx responses:      99858
Total transferred:      18500078 bytes
HTML transferred:       2006610 bytes
Requests per second:    7813.49 [#/sec] (mean)
Time per request:       8.191 [ms] (mean)
Time per request:       0.128 [ms] (mean, across all concurrent requests)
Transfer rate:          1411.62 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.6      1      14
Processing:     0    7  14.5      3     213
Waiting:        0    7  14.4      2     213
Total:          0    8  14.4      4     213

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      5
  80%      5
  90%     10
  95%     47
  98%     64
  99%     78
 100%    213 (longest request)

```

## Particular behaviour

I have noticed a not expected behaviour in the infura api response.
As I have written in dataCollection/testCases.go I was expecting 
the parameter containing the length of the http response inside the header of the response 
itself in any case.


## Last Notes and Comments

I have found this test interesting, not having a clear limit of the scope
of the excise I have tried to make it more real world ready as possible,
with clear limitations on hw trying to imagine the best solution for a retail level sever.

Last mention is on coding time, considering I was working full time while doing test I am proud
of myself of been able to squeeze the 24 hours that took to me to produce it  on as Saturday,
Sunday and a Monday after 8 hours at work. 
