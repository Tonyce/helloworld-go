#!/bin/bash

NAME="graceful"
EXEC_FILE="/Users/ttang/Project/helloworld-go/gracefully/${NAME}"
# PID_FILE="/var/run/${NAME}.go.http.pid"
PID_FILE="/Users/ttang/Project/helloworld-go/gracefully/${NAME}.pid"

case "$1" in 
start)
#    echo ${EXEC_FILE} &
   ${EXEC_FILE} &
   echo $!>${PID_FILE}
   ;;
stop)
   kill -INT `cat ${PID_FILE}`
   echo "------------------------------------"
   rm ${PID_FILE}
   ;;
restart)
   $0 stop
   $0 start
   ;;
status)
   if [ -e ${PID_FILE} ]; then
      echo ${NAME} is running, pid=`cat ${PID_FILE}`
   else
      echo ${NAME} is NOT running
      exit 1
   fi
   ;;
*)
   echo "Usage: $0 {start|stop|status|restart}"
esac

exit 0 
