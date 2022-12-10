# 使用FROM关键字指定基础镜像为golang:1.17，ginrestful设置为项目的名字的全部小写
FROM golang:1.17 AS ginrestful

# 为镜像设置必要的环境变量，具体的参数作用自行去了解，GOPROXY是开启全局代理，下载依赖会更快
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 移动到工作目录，这个目录是项目代码在机器本地的目录，在这个目录里可以看到main.go/go.mod/Dockerfile文件
WORKDIR /Users/chenhuajun/golangprojects/GinRESTful/restapi

# 复制项目中的go.mod和go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

###################
# 创建一个ubuntu的小镜像，用于安装Bash shell
# 这样才能通过docker exec -it ginrestfulapp /bin/bash进入到容器
###################
FROM ubuntu:latest
RUN apt-get update && apt-get install -y bash

# 复制静态文件
COPY ./static /static
COPY ./setting-dev.yaml /setting-dev.yaml

# 从ginrestful镜像中把/app拷贝到当前目录
COPY --from=ginrestful /Users/chenhuajun/golangprojects/GinRESTful/restapi/app /

# 需要运行的命令
ENTRYPOINT ["/app"]