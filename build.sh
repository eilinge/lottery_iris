#!/bin/sh

make run

docker build . -t lottery:latest
docker login --username=eilinge registry.cn-shanghai.aliyuncs.com
docker push registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/eilinge:latest