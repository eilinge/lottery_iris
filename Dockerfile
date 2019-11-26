FROM golang as build

#COPY web/public  /usr/local/go/src/lottery/web/public
COPY out/api_service /usr/local/go/src/lottery/out/api_service

WORKDIR /usr/local/go/src/lottery
ADD  out/api_service api_service
# 使用C语言版本的GO编译器，参数配置为0的时候就关闭C语言版本的编译器了 生成64位linux的可执行程序
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_DB=""

ENV MYSQL_DSN=""
ENV PORT=8080

# 指定aliyun镜像库, 并进行更新, 安装 ca-certificates 支持HTTPS访问,
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && apk add --no-cache tzdata && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

# 二段构建, --from 指定依赖镜像
COPY --from=build /usr/local/go/src/lottery/api_service /usr/bin/api_service

COPY web/public /www/public

RUN chmod +x /usr/bin/api_service

ENTRYPOINT ["api_service"]
