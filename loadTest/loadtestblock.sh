#!/bin/bash
## make the server store the response and get result for a fresh request
ab -n 50000 -c 64 http://127.0.0.1:8123/v1/block/10001
## make the server store the response and get result for a stored request
ab -n 50000 -c 64 http://127.0.0.1:8123/v1/block/10001


## execute the same test for the transaction endpoint
ab -n 50000 -c 64 http://127.0.0.1:8123/v1/tx/10000/1
ab -n 50000 -c 64 http://127.0.0.1:8123/v1/tx/10000/1

