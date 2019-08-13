#!/usr/bin/env bash

EXE_FILENAME='haveyouordered'

echo 'Starting have-you-order server.'

cd scripts
cd ..

pwd

PID=$(pgrep -f ${EXE_FILENAME})
echo "Previous executing: ${PID}"
[[ ${PID} == '' ]] || kill "${PID}"

echo 'Building...'

pwd
go build -o ${EXE_FILENAME} cmd/orderserver/orderserver.go
nohup ./${EXE_FILENAME} &

echo 'have-you-ordered Started.'
