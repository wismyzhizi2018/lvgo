# 以下内容生成docker镜像
 
# 构建阶段 可指定阶段别名 FROM amd64/golang:latest as build_stage
# 基础镜像
FROM golang:1.17-alpine AS builder

# 容器环境变量添加，会覆盖默认的变量值
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"
ENV GIN_MODE=release


# 工作区
WORKDIR /go/src/app

# 复制仓库源文件到容器里
COPY . .

# 编译可执行二进制文件(一定要写这些编译参数，指定了可执行程序的运行平台,参考：https://www.jianshu.com/p/4b345a9e768e)
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver
RUN go install ./../...


# 构建生产镜像，使用最小的linux镜像，只有5M
# 同一个文件里允许多个FROM出现的，每个FROM被称为一个阶段，多个FROM就是多个阶段，最终以最后一个FROM有效，以前的FROM被抛弃
# 多个阶段的使用场景就是将编译环境和生产环境分开
# 参考：https://docs.docker.com/engine/reference/builder/#from
FROM alpine:latest

WORKDIR /root/

# 从编译阶段复制文件
# 这里使用了阶段索引值，第一个阶段从0开始，如果使用阶段别名则需要写成 COPY --from=build_stage /go/src/app/webserver /
#COPY --from=builder /go/src/app/webserver .
COPY --from=builder /go/bin/order /bin/order

# 容器向外提供服务的暴露端口
EXPOSE 8090
 
# 启动服务
#ENTRYPOINT ["./webserver"]
ENTRYPOINT ["/bin/order"]