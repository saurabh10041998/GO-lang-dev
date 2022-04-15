#! /bin/bash

go build $1.go && \
    ./$1 && \
    rm -rf $1
