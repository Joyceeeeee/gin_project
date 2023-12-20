# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest

# 设置工作目录
RUN mkdir -p /app
WORKDIR /app

# 将本地的代码复制到容器内的工作目录
COPY . /app

# 构建 Go 项目
RUN go build -o cmd/youdangzhe main.go

EXPOSE 9001

# 设置容器启动命令
CMD ["cmd/youdangzhe", "server"]
