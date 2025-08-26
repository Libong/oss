FROM alpine AS runner
#作者
MAINTAINER libong
#修改Alpine Linux的APK源，将其从默认的dl-cdn.alpinelinux.org更换为阿里云的镜像源mirrors.aliyun.com 加速后续下载其他软件包的速度
RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#将编译好的执行文件拷贝到docker容器指定位置
COPY ./main /app/
#设置工作目录
WORKDIR /app
#设置环境变量
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
# 暴露服务端口
#EXPOSE 8080
# 定义容器运行时的命令
ENTRYPOINT ["./main"]
CMD ["-configPath","/etc/myconfig","-env","prod","-log","./log"]