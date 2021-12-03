#!/bin/bash
if [ "${1}" != "" ]; then
    padded=$(printf "%02g" ${1})
    cd ${padded}
    go run *.go
    kotlinc main.kt -include-runtime -d ${padded}.jar && java -jar ${padded}.jar && rm -f ${padded}.jar
    cd - > /dev/null
fi
