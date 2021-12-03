#!/bin/bash
__year=2021

if [ "${1}" != "" ]; then
    padded=$(printf "%02g" ${1})
    mkdir -p ${padded}/inputs
    touch ${padded}/inputs/{input,testinput}
    curl -s -H "Cookie: session=`cat .cookie`" https://adventofcode.com/${__year}/day/${1##0}/input > "${padded}/inputs/input"
    cp -n main.go.tmpl ${padded}/main.go
    cp -n main.kt.tmpl ${padded}/main.kt
fi
