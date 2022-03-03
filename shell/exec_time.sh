#!/bin/bash

exec_time(){
    s_time=$(date +%s)
    echo $s_time
    #exec shell code
    sleep 5
    e_time=$(date +%s)
    echo $e_time
    echo "It takes $[$e_time-$s_time] sec."
}

exec_time
