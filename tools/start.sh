#! /bin/sh
mkdir -p logs
touch logs/ddns.log
mkdir -p pids
#nohup ./build/center --id 1 2 >> logs/kingwar.slog & echo $! > pids/center1.pid
nohup ./build/cmd >> logs/ddns.log & echo $! > pids/cmd.pid
ps aux|grep ./build/cmd