#! /bin/sh

if test $# = 1
then
single=-$1
else
single=-2
fi
echo "kill $single ddns"
for pid in `cat pids/ddns.pid`; do kill ${single} ${pid}; done
sleep 1s
ps aux|grep ./builder/ddns