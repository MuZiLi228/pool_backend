CTL = goctl
LINUXENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GOCMD = go
GOBUILD = $(GOCMD) build
PPROFTOOL = $(GOCMD) tool pprof
# pprof 数据来源
PPROFADDR = http://127.0.0.1:7001/sys/pprof
# pprof 结果展示端口
PPROFSHOWADDR = :8086

# 根据proto文件生成对应的数据
#genrpc:
#	$(CTL) rpc proto -src vmallpoint.proto -dir .
# 以linux环境打包程序
# build:
# 	$(LINUXENV) $(GOBUILD) -o program vmallpoint.go
# # 根据Dockerfile构建镜像
# dockerbd:
# 	docker build -t merchant .
# 代码检阅
# codeanl:
# 	golangci-lint run


# pprof  监听
# 当前所有协程的堆栈信息
pprof-goroutine:
	$(PPROFTOOL) -http $(PPROFSHOWADDR) $(PPROFADDR)/goroutine
# 堆上内存使用情况的采样信息
pprof-heap:
	$(PPROFTOOL) -http $(PPROFSHOWADDR) $(PPROFADDR)/heap
# 锁争用情况的采样信息
pprof-mutex:
	$(PPROFTOOL) -http $(PPROFSHOWADDR) $(PPROFADDR)/mutex
# 内存分配情况的采样信息
pprof-allocs:
	$(PPROFTOOL) -http $(PPROFSHOWADDR) $(PPROFADDR)/allocs
# CPU 占用情况的采样信息
pprof-profile:
	$(PPROFTOOL) -http $(PPROFADDR)/profile
# 阻塞操作情况的采样信息
pprof-block:
	$(PPROFTOOL) -http $(PPROFSHOWADDR) $(PPROFADDR)/block