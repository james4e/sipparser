#!/bin/bash

6g bench.go
6l -o bench bench.6
./bench

##rm bench
##rm bench.6
