#!/usr/bin/env bash

EXE_FILENAME='haveyouordered'

echo 'Starting have-you-order server.'

cd ..
go build cmd/orderserver/orderserver.go

PID=$(pgrep -f ${EXE_FILENAME})
echo "Previous executing: ${PID}"
[[ ${PID} == '' ]] || kill "${PID}"

go build -o ${EXE_FILENAME}
nohup ./${EXE_FILENAME} &

echo 'have-you-ordered Started.'
