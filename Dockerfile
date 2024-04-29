#源镜像
FROM golang:1.20-alpine3.16 as builder
#作者
MAINTAINER libong

#RUN set -ex \
#    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
#    && apk --update add tzdata \
#    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#    && apk --no-cache add ca-certificates

#WORKDIR /build
RUN pwd
#RUN ls -l
# 创建了一个app-runner的用户, -D表示无密码
#RUN adduser -u 10001 -D app-runner
# 安装依赖包
ENV GOPROXY https://goproxy.cn
ENV GOPRIVATE github.com/Libong
ENV GO111MODULE on
#RUN git config --global url."https://libong:${{secrets.GO_MOD}}@github.com".insteadOf "https://github.com"
COPY go.mod /app/
COPY go.sum /app/
RUN go mod tidy
# 把当前目录的文件拷过去，编译代码
COPY . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a  -ldflags '-w -s' -o main .

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

