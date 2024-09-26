# 使用官方 Go 镜像作为构建环境
FROM golang:1.22.2 As builder

# 设置工作目录
WORKDIR /app

# 拷贝 go.mod 和 go.sum 文件
COPY go.mod go.sum./
# 下载依赖
RUN go mod download

# 拷贝源代码文件
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main.

# 设置工作目录
WORKDIR /app

# 从构建阶段拷贝编译好的可执行文件
COPY --from=builder /app/main /app/main

# 暴露应用监听的端口号
EXPOSE 8080

# 设置环境变量，指定 Gin 运行模式为 release
ENV GIN_MODE=release

# 运行程序
CMD ["./main"]
