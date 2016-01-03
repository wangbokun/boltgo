#!/bin/bash

git add .
git commit -m  $1
git push


rm -rf ../github.com/wangbokun/boltgo/
go get github.com/wangbokun/boltgo
