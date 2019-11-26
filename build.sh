#!/bin/sh

make build

docker build . -t api_service:latest
docker login --username=eilinge registry.cn-shanghai.aliyuncs.com
docker tag api_service registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/eilinge:latest
docker push registry.cn-shanghai.aliyuncs.com/eilingeloveduzi/eilinge:latest

docker-compose up -d
