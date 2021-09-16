FROM golang:1.16-alpine AS build
#CGO_ENABLED 指明cgo工具是否可用的标识， 1位启用cgo工具，0是关闭，交叉编译不支持cgo工具
ENV    CGO_ENABLED=0 
ENV    GOOS=linux 
ENV    GOARCH=amd64 
ENV    GOPROXY=https://goproxy.cn,direct 


WORKDIR /workspace
COPY go.sum .
COPY go.mod .
RUN go mod tidy 

COPY . .
RUN go build -ldflags="-s -w " -o /app/service .

FROM alpine
    #添加国内源防止超时
    RUN echo -e 'https://mirrors.aliyun.com/alpine/v3.6/main/\nhttps://mirrors.aliyun.com/alpine/v3.6/community/' > /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache ca-certificates tzdata 
    ENV TZ Asia/Shanghai 
    ENV GIN_MODE=release
    
    # 设置工作目录
    WORKDIR /data/app
    # 复制生成的可执行命令
    COPY --from=build /app/service .
    COPY --from=build /workspace/docs ./docs
    COPY --from=build /workspace/configs ./configs
    COPY --from=build /workspace/src/public ./src/public

EXPOSE 7001
CMD [ "./service" ]