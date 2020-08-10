#! /bin/bash

set -x

check_arg () {
  t=$1
  arg_list="file tcp"

  for arg in $arg_list; do
    if [ $t == $arg ]; then
      return
    fi
  done
  echo 'set type!'
  exit 1
}

check_file () {
    if [ -f results/$1/$1.txt ]; then
        rm results/$1/$1.txt
    fi

    if [ $1 == "file" ]; then
      if [ -f results/$1/test.txt ]; then
        rm results/$1/test.txt
      fi
    fi
}

exec_trace () {
    t=$1
    ./syscaller -s=$t >/dev/null &

    PID=`ps -aux | grep syscaller | awk '{print $2}' | head -n 1`
    echo "target pid: $PID"

    TARGET_SYSCALLS="openat,read,write,close,fstat,socket,setsockopt,connect,accept4,getsockopt,getsockname,setsockopt"
    #TARGET_SYSCALLS="all" # default

    case $t in
      "file")
        strace -e trace=$TARGET_SYSCALLS -f -p $PID -o results/file/file.txt;;
      "tcp")
        strace -e trace=$TARGET_SYSCALLS -f -p $PID -o results/tcp/tcp.txt;;
      *)
        echo 'set type!'
        exit 1
    esac
}

type=$1

check_arg $type
check_file $type
exec_trace $type
