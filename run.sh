#!/bin/bash
if [ "${1}" != "" ]; then
    padded=$(printf "%02g" ${1})
    cd ${padded}
    go run *.go
    cd - > /dev/null
fi