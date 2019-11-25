FROM golang as build

ADD . /usr/local/go/src/lottery

WORKDIR /usr/local/go/src/lottery

# 使用C语言版本的GO编译器，参数配置为0的时候就关闭C语言版本的编译器了 生成64位linux的可执行程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server


FROM alpine:3.7

ENV REDIS=ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""

ENV MYsqlDSN=""
# ENV GIN_MODE="release"
ENV PORT=8080

# 指定aliyun镜像库, 并进行更新, 安装 ca-certificates 支持HTTPS访问,
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

# 二段构建, --from 指定依赖镜像
COPY --from=build /usr/local/go/src/lottery/api_server /usr/bin/api_server

ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]