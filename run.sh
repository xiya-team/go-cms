#!/bin/bash

PID=$(ps -ef | grep go-cms | grep -v grep | awk '{ print $2 }')

case $1 in 
	start)
	    go build
		nohup ./go-cms 2>&1 >> project.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
	    if [ -z "$PID" ]
        then
            echo Application is already stopped
        else
            echo kill $PID
            kill $PID
        fi
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		if [ -z "$PID" ]
        then
            echo Application is already stopped
        else
            echo kill $PID
            kill $PID
        fi
        go build
		sleep 1
		nohup ./go-cms 2>&1 >> project.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac