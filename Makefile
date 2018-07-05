CURRENT_GIT_PATH := github.com/CX1ng
CURRENT_GIT_REPO := proxypool

#交叉编译为Linux
COMMONENVVAR      ?= GOOS=linux GOARCH=amd64
#禁用CGO用于静态编译
BUILDENVVAR       ?= CGO_ENABLED=0

#构建
all: deps linux_build

#下载依赖
deps:
	glide install

#更新并下载依赖
update_deps:
	glide update

#编译
build:
	go build -o bin/$(CURRENT_GIT_REPO) ./cmd/proxy_pool
	go build -o bin/web ./cmd/test_web
	
#交叉编译出linux下的静态可执行文件
linux_build: deps
	$(COMMONENVVAR) $(BUILDENVVAR) make build

test:
	go test -v $(CURRENT_GIT_PATH)/$(CURRENT_GIT_REPO)/dao
	go test -v $(CURRENT_GIT_PATH)/$(CURRENT_GIT_REPO)/models
	#FIX ME go test -v $(CURRENT_GIT_PATH)/$(CURRENT_GIT_REPO)/processord/parser

clean:
	@rm -rf bin _project vendor

.PHONY: all deps build linux_build