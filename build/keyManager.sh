#!/bin/bash

# Describe：
# Author：yubo.shen
# Create Date：2022.06.01
# Current version：v1.0
# history version：v1.0
# his version Desc：


CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o ../bin/ethKeyManage_mac  ../main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/ethKeyManage_linux  ../main.go

