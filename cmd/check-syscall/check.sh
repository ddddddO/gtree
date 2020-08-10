#! /bin/bash

set -x

check_file () {
    if [ -f result.txt ]; then
        rm result.txt
    fi

    if [ -f test.txt ]; then
        rm test.txt
    fi
}

exec_trace () {
    ./syscaller -s=tcp >/dev/null &

    PID=`ps -aux | grep syscaller | awk '{print $2}' | head -n 1`
    echo "target pid: $PID"

    TARGET_SYSCALLS="openat,read,write,close,fstat"
    #TARGET_SYSCALLS="all" # default

    strace -e trace=$TARGET_SYSCALLS -f -p $PID -o result.txt
}

check_file
exec_trace