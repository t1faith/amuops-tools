#!/bin/bash

exec_time(){
    s_time=$(date +%s)
    echo $srt
    #exec shell code
    sleep 5
    e_time=$(date +%s)
    echo $end
    echo "It takes $[$e_time-$s_time] sec."
}

exec_time
