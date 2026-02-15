# 使用轻量级 Alpine 作为运行镜像
FROM ubuntu:24.04

WORKDIR /app

# 直接复制本地编译好的 server 二进制文件到镜像
COPY server ./

# 暴露 gRPC 端口
EXPOSE 50051

# 启动 server
CMD ["./server"]
