# 使用官方的Golang基础镜像
FROM golang:1.22rc2-bullseye

# 设置工作目录
WORKDIR /app

# 将本地的代码复制到容器中
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 暴露应用程序监听的端口
EXPOSE 8089

# 启动应用程序
CMD ["./main"]
