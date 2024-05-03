#源镜像
FROM golang:1.20-alpine3.16 as builder
#作者
MAINTAINER libong

#RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#    && apk --update add tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && apk --no-cache add ca-certificates

# 创建了一个app-runner的用户, -D表示无密码
#RUN adduser -u 10001 -D app-runner
# 设置环境变量
#ENV GOPROXY https://goproxy.cn
#ENV GOPRIVATE github.com/Libong
#ENV CGO_ENABLED 1
#ENV GO111MODULE on
#RUN git config --global url."https://libong:${{secrets.GO_MOD}}@github.com".insteadOf "https://github.com"
RUN apk add --no-cache libc6-compat
# 将当前目录的代码推送到docker容器里的目录下
#COPY go.mod /app/
#COPY go.sum /app/
# 把当前目录的文件拷过去，编译代码
COPY ./main /app/
WORKDIR /app
RUN ls -l
#RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main ./app/interface/oss/cmd
#RUN ls -l
#RUN go mod tidy
#RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a  -ldflags '-w -s' -o main .

# 暴露服务端口
#EXPOSE 8088
#FROM alpine:3.16 AS final
# 把构建结果、配置文件（有的话）和用户的相关文件拷过去
#WORKDIR /app
#COPY --from=builder /build/app/interface/oss/cmd/main /app
#COPY --from=builder /build/app/interface/oss/cmd/config.yaml /app/conf
#COPY --from=builder /build/config.toml /app
# 下载时区包
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# 设置当前时区
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# https ssl证书
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 使用app-runner启动
#USER app-runner
#ENTRYPOINT ["/app/blockchain-middleware"]

# 定义容器运行时的命令

#CMD ["/app/main"]

#RUN rm -rf /var/cache/apk/* 无用

