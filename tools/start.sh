#! /bin/sh
mkdir -p logs
touch logs/server.log
mkdir -p pids
#nohup ./build/center --id 1 2 >> logs/kingwar.slog & echo $! > pids/center1.pid
nohup ./builder/cmd >> logs/ddns.log & echo $! > pids/cmd.pid
ps aux|grep ./builder/cmd